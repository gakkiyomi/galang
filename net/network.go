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
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/dselans/dmidecode"
	str "github.com/gakkiyomi/galang/string"
	"github.com/gakkiyomi/galang/utils"
	"github.com/songtianyi/rrframework/logs"
)

//SubnetInfo  apache subnetutils for go version
type SubnetInfo struct {
	netmask   uint32 //子网掩码
	network   uint32 //网络位
	address   uint32 //IP地址
	broadcast uint32 //广播位
	cidr      string
	fullCidr  string
}

//NewSubnetInfo 创建一个subnetInfo对象
func NewSubnetInfo(cidr string) (*SubnetInfo, error) {

	b, _ := regexp.MatchString(CIDR_REG, cidr)
	c, _ := regexp.MatchString(FULL_CIDR_REG, cidr)

	var fullCidr string

	if b == false {

		if c == false {
			return nil, fmt.Errorf("cidr:%v is not valid, pattern should like: 192.168.1.0/24 or 192.168.1.0/255.255.255.0", cidr)
		}

		fullCidr = cidr

		str := strings.Split(cidr, "/")
		addr := str[0]
		mask := str[1]
		maskLength, _ := Network.NetmaskToMaskLength(mask)
		cidr = addr + "/" + utils.Transform.IntToString(maskLength)
	}

	_, sub, _ := net.ParseCIDR(cidr)

	maskLength, _ := sub.Mask.Size()
	suffix, _ := Network.MaskLengthToNetmask(maskLength)

	str := strings.Split(cidr, "/")
	addr := str[0]
	fullCidr = addr + "/" + suffix

	longIP, _ := iP2long(sub.IP.String())
	longMask, _ := iP2long(suffix)
	netwrokAddr := (longIP & longMask)
	broadcastAddr := netwrokAddr | (^longMask)

	return &SubnetInfo{
		address:   longIP,
		netmask:   longMask,
		network:   netwrokAddr,
		broadcast: broadcastAddr,
		cidr:      cidr,
		fullCidr:  fullCidr,
	}, nil
}

//ToString 返回SubnetInfo的数据
func (s *SubnetInfo) ToString() string {
	builder := str.String.NewStringBuilder("")
	builder.Append("CIDR Signature:\t[").Append(s.GetCidrSignature()).Append("]")
	builder.Append(" Netmask: [").Append(s.NetmaskString()).Append("]\n")
	builder.Append("Network:\t[").Append(s.NetworkString()).Append("]\n")
	builder.Append("Broadcast:\t[").Append(s.BradcastString()).Append("]\n")
	builder.Append("First Address:\t[").Append(s.LowAddress()).Append("]\n")
	builder.Append("Last Address:\t[").Append(s.HighAddress()).Append("]\n")
	builder.Append("# Addresses:\t[").Append(strconv.Itoa(int(s.Size()))).Append("]\n")
	return builder.ToString()
}

//AddressString 获取ip地址
func (sub *SubnetInfo) AddressString() string {
	return Network.Long2ip(sub.address)
}

//NetmaskString 获取子网掩码
func (sub *SubnetInfo) NetmaskString() string {
	return Network.Long2ip(sub.netmask)
}

//NetworkString 获取网络地址
func (sub *SubnetInfo) NetworkString() string {
	return Network.Long2ip(sub.network)
}

//BradcastString 获取广播地址
func (sub *SubnetInfo) BradcastString() string {
	return Network.Long2ip(sub.broadcast)
}

//IsRangeOf 判断输入的ip是否在当前网段内
func (sub *SubnetInfo) IsRangeOf(addr string) (bool, error) {
	return isRangeOf(addr, sub.cidr)
}

func (sub *SubnetInfo) low() uint32 {
	if sub.broadcast-sub.network > 1 {
		return sub.network + 1
	}
	return 0
}

//LowAddress return the first available address in current subnet
func (sub *SubnetInfo) LowAddress() string {
	return Network.Long2ip(sub.low())
}

func (sub *SubnetInfo) high() uint32 {
	if sub.broadcast-sub.network > 1 {
		return sub.broadcast - 1
	}
	return 0
}

//HighAddress return the last available address in current subnet
func (sub *SubnetInfo) HighAddress() string {
	return Network.Long2ip(sub.high())
}

//Size returns available ip address size
func (sub *SubnetInfo) Size() uint32 {
	if sub.high() == sub.low() {
		return 0
	}
	return sub.high() - sub.low() + 1
}

//GetCidrSignature 获取无分类地址
func (sub *SubnetInfo) GetCidrSignature() string {
	return sub.cidr
}

//NetmaskToMaskLength 255.255.255.0 >>> 24
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

//MaskLengthToNetmask 24 >>> 255.255.255.0
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

//Long2ip .
func (*GalangNet) Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

//IP2long .
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

//IsRangeOf if addr is range of cidr returns true
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
	_, subNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("prase cidr %v failed", cidr)
	}
	if subNet.Contains(ip) {
		return true, nil
	}
	return false, nil
}

//GetSystemUUIDForLinux get linux dmidecode -s system-uuid
func (*GalangNet) GetSystemUUIDForLinux() (string, error) {

	if !isLinux() {
		return "", fmt.Errorf("only support linux system")
	}

	dmi := dmidecode.New()

	if err := dmi.Run(); err != nil {
		return "", fmt.Errorf("Unable to get dmidecode information. Error: %v", err)
	}
	systemInfo, _ := dmi.SearchByName("System Information")

	for _, v := range systemInfo {
		if _, ok := v["UUID"]; ok {
			return v["UUID"], nil
		}

	}
	return "", fmt.Errorf("can't get system uuid")
}

func isWindows() bool {
	return "windows" == runtime.GOOS
}

func isLinux() bool {
	return "linux" == runtime.GOOS
}

//GetCIDRAvailableAddrList retrun available ip address list
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

//GetRangeAddrList 获取一个范围里的ip地址   192.168.1.222-192.168.1.228
func (*GalangNet) GetRangeAddrList(_range string) ([]string, error) {

	isRange, _ := regexp.MatchString(RANGE_REG, _range)
	if isRange == false {
		return nil, fmt.Errorf("prase range addr %v failed", _range)
	}
	var lowAndhigh []string
	if b := strings.Contains(_range, "~"); b {
		lowAndhigh = strings.Split(_range, "~")
	} else {
		lowAndhigh = strings.Split(_range, "-")
	}

	low := net.ParseIP(lowAndhigh[0])
	high := net.ParseIP(lowAndhigh[1])

	var ips []string

	lowLong, _ := iP2long(low.String())
	highLong, _ := iP2long(high.String())
	for ; lowLong < highLong; inc(low) {
		lowLong, _ = iP2long(low.String())
		ips = append(ips, low.String())
	}

	return ips, nil
}

//LPM return Longest prefix match
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

		_, ipnetMax, _ := net.ParseCIDR(maxVal)
		_, ipnet, _ := net.ParseCIDR(filter[i])

		max, _ := ipnetMax.Mask.Size()
		vi, _ := ipnet.Mask.Size()

		if max < vi {
			maxVal = filter[i]
		}

	}
	return maxVal, nil
}

//GetLocalIPv4 获取本机ipv4地址
func (*GalangNet) GetLocalIPv4() (string, error) {
	addrs, err := net.InterfaceAddrs()
	var ip string
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip, nil
}

// GetExternalIP 获取公网IP地址
func (*GalangNet) GetExternalIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(b)), nil
}

//GetIPSFormat will format target string to a valid string array
//if target is a ip like 192.168.1.222  will format ['192.168.1.222']
//if target is a range of ip like 192.168.1.222[-||~]192.168.1.224  will format ['192.168.1.222','192.168.1.223','192.168.1.224']
func (*GalangNet) IPSFormat(target string) ([]string, error) {
	is_Range, _ := regexp.MatchString(RANGE_REG, target)
	is_Ip, _ := regexp.MatchString(IP_REG, target)

	if is_Range == true {
		list, err := Network.GetRangeAddrList(target)
		if err != nil {
			logs.Error(err.Error())
		}
		return list, nil
	}

	if is_Ip == true {
		return []string{target}, nil
	}

	return nil, fmt.Errorf("target: %s  is not vaild ip or range of ip", target)
}
