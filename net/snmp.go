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
	"github.com/alouca/gosnmp"
)

var System_obj = map[string]string{
	".1.3.6.1.2.1.1.1.0": "Description",
	".1.3.6.1.2.1.1.2.0": "Vendor",
	".1.3.6.1.2.1.1.3.0": "UpTime",
	".1.3.6.1.2.1.1.4.0": "Contract",
	".1.3.6.1.2.1.1.5.0": "Hostname",
	".1.3.6.1.2.1.1.6.0": "Location",
}
var System_oid = map[string]string{
	"Description": ".1.3.6.1.2.1.1.1.0",
	"Vendor":      ".1.3.6.1.2.1.1.2.0",
	"UpTime":      ".1.3.6.1.2.1.1.3.0",
	"Contract":    ".1.3.6.1.2.1.1.4.0",
	"Hostname":    ".1.3.6.1.2.1.1.5.0",
	"Location":    ".1.3.6.1.2.1.1.6.0",
}

var System_oids = []string{System_oid[`Description`], System_oid[`Vendor`], System_oid[`UpTime`], System_oid[`Contract`], System_oid[`Hostname`], System_oid[`Location`]}

type SNMPWrapper struct {
	Client *gosnmp.GoSNMP
}

type System struct {
	Hostname    string
	Location    string
	Contract    string
	Vendor      string
	UpTime      int
	Description string
}

func (*GalangSNMP) NewSnmpWrapper(target, community string, version gosnmp.SnmpVersion, timeout int64) (*SNMPWrapper, error) {
	client, err := gosnmp.NewGoSNMP(target, community, version, timeout)
	if err != nil {
		return nil, err
	}
	return &SNMPWrapper{
		Client: client,
	}, nil
}

func (wrap *SNMPWrapper) SystemInfo() (*System, error) {
	pack, err := wrap.Client.GetMulti(System_oids)
	if err != nil {
		return nil, err
	}
	vars := pack.Variables
	res := &System{}
	for _, v := range vars {

		switch v.Name {
		case System_oid[`Description`]:
			res.Description = v.Value.(string)
		case System_oid[`Location`]:
			res.Location = v.Value.(string)
		case System_oid[`Contract`]:
			res.Contract = v.Value.(string)
		case System_oid[`Vendor`]:
			//TODO
		case System_oid[`UpTime`]:
			res.UpTime = v.Value.(int)
		case System_oid[`Hostname`]:
			res.Hostname = v.Value.(string)
		}

	}
	return res, nil
}
