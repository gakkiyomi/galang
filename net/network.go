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

	"github.com/gakkiyomi/galang/utils"
)

// 255.255.255.0 >>> 24
func (*GalangNet) NetmaskToCIDR(netmask string) (int, error) {

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

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

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

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
