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

type SubnetInfo struct {
	netmask   uint32 //子网掩码
	network   uint32 //网络位
	address   uint32 //IP地址
	broadcast uint32 //广播位
	cidr      string
	full_cidr string
}

func NewSubnetInfo(cidr string) (*SubnetInfo, error) {

	b, _ := regexp.MatchString(CIDR_REG, cidr)
	c, _ := regexp.MatchString(FULL_CIDR_REG, cidr)

	var full_cidr string

	if b == false {

		if c == false {
			return nil, fmt.Errorf("cidr:%v is not valid, pattern should like: 192.168.1.0/24 or 192.168.1.0/255.255.255.0", cidr)
		}

		full_cidr = cidr

		str := strings.Split(cidr, "/")
		addr := str[0]
		mask := str[1]
		mask_length, _ := Network.NetmaskToMaskLength(mask)
		cidr = addr + "/" + utils.Transform.IntToString(mask_length)
	}

	str := strings.Split(cidr, "/")
	addr := str[0]
	mask_length, _ := utils.Transform.StringToInt(str[1])
	mask, _ := Network.MaskLengthToNetmask(mask_length)
	full_cidr = addr + "/" + mask

	_, sub, _ := net.ParseCIDR(cidr)

	cidr_sr, _ := sub.Mask.Size()
	suffix, _ := Network.MaskLengthToNetmask(cidr_sr)

	longIp, _ := iP2long(sub.IP.String())
	longMask, _ := iP2long(suffix)
	netwrok_addr := (longIp & longMask)
	broadcast_addr := netwrok_addr | (^longMask)

	return &SubnetInfo{
		address:   longIp,
		netmask:   longMask,
		network:   netwrok_addr,
		broadcast: broadcast_addr,
		cidr:      cidr,
		full_cidr: full_cidr,
	}, nil
}

func (sub *SubnetInfo) AddressString() string {
	return Network.Long2ip(sub.address)
}

func (sub *SubnetInfo) NetmaskString() string {
	return Network.Long2ip(sub.netmask)
}

func (sub *SubnetInfo) NetworkString() string {
	return Network.Long2ip(sub.network)
}

func (sub *SubnetInfo) BradcastString() string {
	return Network.Long2ip(sub.broadcast)
}

func (sub *SubnetInfo) IsRangeOf(addr string) (bool, error) {
	return isRangeOf(addr, sub.cidr)
}

func (sub *SubnetInfo) low() uint32 {
	if sub.broadcast-sub.network > 1 {
		return sub.network + 1
	}
	return 0
}

//return the first available address in current subnet
func (sub *SubnetInfo) LowAddress() string {
	return Network.Long2ip(sub.low())
}

func (sub *SubnetInfo) high() uint32 {
	if sub.broadcast-sub.network > 1 {
		return sub.broadcast - 1
	}
	return 0
}

//return the last available address in current subnet
func (sub *SubnetInfo) HighAddress() string {
	return Network.Long2ip(sub.high())
}

//returns available ip address size
func (sub *SubnetInfo) Size() uint32 {
	if sub.high() == sub.low() {
		return 0
	}
	return sub.high() - sub.low() + 1
}

func (sub *SubnetInfo) GetCidrSignature() string {
	return sub.cidr
}

// 255.255.255.0 >>> 24
func (*GalangNet) NetmaskToMaskLength(netmask string) (int, error) {

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
func (*GalangNet) MaskLengthToNetmask(cidr int) (string, error) {

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

	return iP2long(ipstr)
}

func iP2long(ipstr string) (uint32, error) {

	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0, nil
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

// if addr is range of cidr returns true
func (*GalangNet) IsRangeOf(addr, cidr string) (bool, error) {

	a, _ := regexp.MatchString(IP_REG, addr)
	b, _ := regexp.MatchString(CIDR_REG, cidr)

	if a == false || b == false {
		return false, fmt.Errorf("prase addr %v or %v failed", addr, cidr)
	}
	return isRangeOf(addr, cidr)
}

func isRangeOf(addr, cidr string) (bool, error) {
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

func (*GalangNet) GetRangeAddrList(_range string) ([]string, error) {

	is_Range, _ := regexp.MatchString(RANGE_REG, _range)
	if is_Range == false {
		return nil, fmt.Errorf("prase range addr %v failed", _range)
	}
	var low_high []string
	if b := strings.Contains(_range, "~"); b {
		low_high = strings.Split(_range, "~")
	} else {
		low_high = strings.Split(_range, "-")
	}

	low := net.ParseIP(low_high[0])
	high := net.ParseIP(low_high[1])

	var ips []string

	low_long, _ := iP2long(low.String())
	high_long, _ := iP2long(high.String())
	for ; low_long < high_long; inc(low) {
		low_long, _ = iP2long(low.String())
		ips = append(ips, low.String())
	}

	return ips, nil
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
