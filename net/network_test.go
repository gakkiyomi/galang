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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCIDRToNetmask(t *testing.T) {
	mask, err := Net.CIDRToNetmask(19)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "255.255.224.0", mask, "The two netmask should be the same.")
}

func TestNetmaskToCIDR(t *testing.T) {
	cidr_suffix, err := Net.NetmaskToCIDR("255.255.aa.0")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, 18, cidr_suffix, "The two cidr_suffix should be the same.")
}

func TestLong2ip(t *testing.T) {

	cidr_suffix := Net.Long2ip(3232235521)
	assert.Equal(t, "192.168.0.1", cidr_suffix, "The two cidr_suffix should be the same.")
}
