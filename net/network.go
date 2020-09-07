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
	"errors"

	"github.com/gakkiyomi/galang/utils"
)

func (*GalangNet) NetmaskToCIDR(mask string) (int, error) {
	//TODO
	return 0, nil
}

func (*GalangNet) CIDRToNetmask(cidr int) (string, error) {

	if cidr < 0 || cidr > 32 {
		return "", errors.New("cidr must be less than 32 and more than 0")
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
