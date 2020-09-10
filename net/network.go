// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/dselans/dmidecode"
	"github.com/gakkiyomi/galang/utils"
)

const (
	IP_REG   = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`
	CIDR_REG = `^((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.){3}(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\[[^\[\].;\s]{1,100}\]|)/(1[0-9]|2[0-9]|3[0-2]|[0-9])$`
)

type SubnetInfo struct {
	Netmask   uint32 //子网掩码
	Network   uint32 //网络位
	Address   uint32 //IP地址
	Broadcast uint32 //广播位
}

func NewSubnetInfo(cidr string) (*SubnetInfo, error) {

	b, _ := regexp.MatchString(CIDR_REG, cidr)
	if b == false {
		return nil, fmt.Errorf("cidr:%v is not valid, pattern should like: 192.168.1.0/24", cidr)
	}

	_, sub, _ := net.ParseCIDR(cidr)

	cidr_sr, _ := sub.Mask.Size()
	suffix, _ := Network.CIDRToNetmask(cidr_sr)

	longIp, _ := Network.IP2long(sub.IP.String())
	longMask, _ := Network.IP2long(suffix)
	netwrok_addr := (longIp & longMask)
	broadcast_addr := netwrok_addr | (^longMask)

	return &SubnetInfo{
		Address:   longIp,
		Netmask:   longMask,
		Network:   netwrok_addr,
		Broadcast: broadcast_addr,
	}, nil
}

func (sub *SubnetInfo) AddressString() string {
	return Network.Long2ip(sub.Address)
}

func (sub *SubnetInfo) NetmaskString() string {
	return Network.Long2ip(sub.Netmask)
}

func (sub *SubnetInfo) NetworkString() string {
	return Network.Long2ip(sub.Network)
}

func (sub *SubnetInfo) BradcastString() string {
	return Network.Long2ip(sub.Broadcast)
}

// 255.255.255.0 >>> 24
func (*GalangNet) NetmaskToCIDR(netmask string) (int, error) {

	re := regexp.MustCompile(IP_REG)

	if re.MatchString(netmask) == false {
		return 0, fmt.Errorf("netmask:%v is not valid, pattern should like: 255.255.255.0", netmask)
	}

	ipSplitArr := strings.Split(netmask, ".")

	ipv4MaskArr := make([]byte, 4)
	for i, value := range ipSplitArr {
		intValue, err := utils.Transform.StringToInt(value)
		if err != nil {
			return 0, fmt.Errorf("type transform error:[%v] string value is: [%s]", err, value)
		}
		if intValue > 255 {
			return 0, fmt.Errorf("netmask cannot greater than 255, current value is: [%s]", value)
		}
		ipv4MaskArr[i] = byte(intValue)
	}

	ones, _ := net.IPv4Mask(ipv4MaskArr[0], ipv4MaskArr[1], ipv4MaskArr[2], ipv4MaskArr[3]).Size()
	return ones, nil
}

// 24 >>> 255.255.255.0
func (*GalangNet) CIDRToNetmask(cidr int) (string, error) {

	if cidr < 0 || cidr > 32 {
		return "", fmt.Errorf("cidr must be less than 32 and greater than 0")
	}

	mask := (0xFFFFFFFF << (32 - cidr)) & 0xFFFFFFFF
	var dmask uint64
	dmask = 32
	var localmask string
	for i := 1; i <= 4; i++ {
		tmp := mask >> (dmask - 8) & 0xFF
		tmpStr := utils.Transform.IntToString(tmp)
		if i == 1 {
			localmask = tmpStr
		} else {
			localmask = localmask + "." + tmpStr
		}

		dmask -= 8
	}
	return localmask, nil
}

func (*GalangNet) Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

func (*GalangNet) IP2long(ipstr string) (uint32, error) {

	re := regexp.MustCompile(IP_REG)

	if re.MatchString(ipstr) == false {
		return 0, fmt.Errorf("ip:%v is not valid, pattern should like: 192.168.1.1", ipstr)
	}

	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0, nil
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

// if addr is range of cidr returns true
func (*GalangNet) IsRangeOf(addr, cidr string) (bool, error) {
	ip := net.ParseIP(addr)
	_, sub_net, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("prase cidr %v failed", cidr)
	}
	if sub_net.Contains(ip) {
		return true, nil
	}
	return false, nil
}

//get linux dmidecode -s system-uuid
func (*GalangNet) GetSystemUUID_Linux() (string, error) {
	dmi := dmidecode.New()

	if err := dmi.Run(); err != nil {
		return "", fmt.Errorf("Unable to get dmidecode information. Error: %v\n", err)
	}
	system_info, _ := dmi.SearchByName("System Information")

	for _, v := range system_info {
		if _, ok := v["UUID"]; ok {
			return v["UUID"], nil
		}

	}
	return "", fmt.Errorf("can't get system uuid")
}

//retrun available ip address list
func (*GalangNet) GetCIDRAvailableAddrList(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) > 1 {
		ips = ips[1 : len(ips)-1]
	}

	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// return Longest prefix match
func (n *GalangNet) LPM(ip string, subnets []string) (string, error) {
	var filter []string
	for _, cidr := range subnets {
		b, err := n.IsRangeOf(ip, cidr)
		if err != nil || b == false {
			continue
		}
		filter = append(filter, cidr)
	}
	if len(filter) == 0 {
		return "", fmt.Errorf("longest prefix match fail in %v", subnets)
	}

	maxVal := filter[0]

	for i := 1; i < len(filter); i++ {

		_, ipnet_max, _ := net.ParseCIDR(maxVal)
		_, ipnet_i, _ := net.ParseCIDR(filter[i])

		max, _ := ipnet_max.Mask.Size()
		vi, _ := ipnet_i.Mask.Size()

		if max < vi {
			maxVal = filter[i]
		}

	}
	return maxVal, nil
}
