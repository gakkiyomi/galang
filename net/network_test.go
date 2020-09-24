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
	mask, err := Network.MaskLengthToNetmask(19)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "255.255.224.0", mask, "The two netmask should be the same.")
}

func TestNetmaskToCIDR(t *testing.T) {
	cidr_suffix, err := Network.NetmaskToMaskLength("255.255.255.0")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, 24, cidr_suffix, "The two item should be the same.")
}

func TestLong2ip(t *testing.T) {

	cidr_suffix := Network.Long2ip(3232235521)
	assert.Equal(t, "192.168.0.1", cidr_suffix, "The two item should be the same.")
}

func TestIP2long(t *testing.T) {

	l, err := Network.IP2long("192.168.0.1")
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, uint32(3232235521), l, "The two item should be the same.")
}

//must be linux
/*func TestGetSystemUUID_Linux(t *testing.T) {

	l, err := Network.GetSystemUUID_Linux()
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, "5F190D42-AD0F-D15F-28C4-44C5A755338C", l, "The two item should be the same.")
}*/
func TestIsRangeOf(t *testing.T) {

	l, err := Network.IsRangeOf("192.167.21.1", "192.168.1.0/16")
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, false, l, "The two item should be the same.")
}

func TestGetCIDRAvailableAddrList(t *testing.T) {

	l, err := Network.GetCIDRAvailableAddrList("192.168.1.0/24")
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 254, len(l), "The two item should be the same.")
}

func TestLPM(t *testing.T) {

	sbnts := []string{`192.168.1.0/24`, `192.168.0.0/16`, `192.0.0.0/8`, `192.168.2.0/24`}

	cidr, err := Network.LPM("192.168.1.2", sbnts)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, `192.168.1.0/24`, cidr, "The two item should be the same.")
}

func TestSubnetInfo(t *testing.T) {
	info, err := NewSubnetInfo(`192.168.0.0/17`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, `192.168.0.0`, info.AddressString(), "The two item should be the same.")
	assert.Equal(t, `255.255.128.0`, info.NetmaskString(), "The two item should be the same.")
	assert.Equal(t, `192.168.0.0`, info.NetworkString(), "The two item should be the same.")
	assert.Equal(t, `192.168.127.255`, info.BradcastString(), "The two item should be the same.")

	info2, err := NewSubnetInfo(`192.168.0.0/255.255.128.0`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, `192.168.0.0`, info2.AddressString(), "The two item should be the same.")
	assert.Equal(t, `255.255.128.0`, info2.NetmaskString(), "The two item should be the same.")
	assert.Equal(t, `192.168.0.0`, info2.NetworkString(), "The two item should be the same.")
	assert.Equal(t, `192.168.127.255`, info2.BradcastString(), "The two item should be the same.")
}

func TestLowHighAddress(t *testing.T) {

	info, err := NewSubnetInfo(`192.168.1.2/21`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, `192.168.0.1`, info.LowAddress(), "The two item should be the same.")
	assert.Equal(t, `192.168.7.254`, info.HighAddress(), "The two item should be the same.")

}

func TestSubnetSize(t *testing.T) {

	info, err := NewSubnetInfo(`192.168.1.0/24`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, uint32(254), info.Size(), "The two item should be the same.")

}

func TestGetCidrSignature(t *testing.T) {

	info, err := NewSubnetInfo(`192.168.1.0/255.255.255.0`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, `192.168.1.0/24`, info.GetCidrSignature(), "The two item should be the same.")

}

func TestRange(t *testing.T) {

	info, err := Network.GetRangeAddrList(`192.168.1.1-192.168.1.254`)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 254, len(info), "The two item should be the same.")

}
