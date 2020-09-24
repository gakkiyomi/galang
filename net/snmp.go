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
	"net"

	"github.com/alouca/gosnmp"
	"github.com/gakkiyomi/galang/utils"
	"github.com/songtianyi/rrframework/logs"
)

var System_obj = map[string]string{
	".1.3.6.1.2.1.1.1.0": "sysDescr",
	".1.3.6.1.2.1.1.2.0": "sysObjectID",
	".1.3.6.1.2.1.1.3.0": "sysUpTime",
	".1.3.6.1.2.1.1.4.0": "sysContact",
	".1.3.6.1.2.1.1.5.0": "sysName",
	".1.3.6.1.2.1.1.6.0": "sysLocation",
}
var System_oid = map[string]string{
	"sysDescr":    ".1.3.6.1.2.1.1.1.0",
	"sysObjectID": ".1.3.6.1.2.1.1.2.0",
	"sysUpTime":   ".1.3.6.1.2.1.1.3.0",
	"sysContact":  ".1.3.6.1.2.1.1.4.0",
	"sysName":     ".1.3.6.1.2.1.1.5.0",
	"sysLocation": ".1.3.6.1.2.1.1.6.0",
}

var System_oids = []string{System_oid[`sysDescr`], System_oid[`sysObjectID`], System_oid[`sysUpTime`], System_oid[`sysContact`], System_oid[`sysName`], System_oid[`sysLocation`]}

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

type InterfaceStatus int32

//The current operational state of the interface. The testing(3) state indicates that no operational packets can be passed.
//If ifAdminStatus is down(2) then ifOperStatus should be down(2).
//If ifAdminStatus is changed to up(1) then ifOperStatus should change to up(1)
//if the interface is ready to transmit and receive network traffic; it should change to dormant(5)
//if the interface is waiting for external actions (such as a serial line waiting for an incoming connection); it should remain in the down(2) state
//if and only if there is a fault that prevents it from going to the up(1) state; it should remain in the notPresent(6) state
//if the interface has missing (typically, hardware) components.
const (
	UP         InterfaceStatus = 1
	DOWN       InterfaceStatus = 2
	Testing    InterfaceStatus = 3
	Unknown    InterfaceStatus = 4
	Dormant    InterfaceStatus = 5
	NotPresent InterfaceStatus = 6
)

var InterfaceStatus_name = map[int32]string{
	1: "UP",
	2: "DOWN",
	3: "Testing",
	4: "Unknown",
	5: "Dormant",
	6: "NotPresent",
}
var InterfaceStatus_value = map[string]int32{
	"UP":         1,
	"DOWN":       2,
	"Testing":    3,
	"Unknown":    4,
	"Dormant":    5,
	"NotPresent": 6,
}

func (x InterfaceStatus) String() string {
	s, ok := InterfaceStatus_name[int32(x)]
	if ok {
		return s
	}
	return utils.Transform.AnyToString(x)
}

type Interface struct {
	Name       string          `json:"name,omitempty"`
	MacAddress string          `json:"mac_addr,omitempty"`
	IPv4       string          `json:"ip,omitempty"`
	NetMask    string          `json:"netmask,omitempty"`
	Status     InterfaceStatus `json:"status,omitempty"`
}

func (intf *Interface) ToString() string {
	return utils.Transform.AnyToString(intf)
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
		logs.Error(err.Error())
		return nil, err
	}
	res := &System{}
	for _, v := range pack.Variables {

		switch v.Name {
		case System_oid[`sysDescr`]:
			res.Description = v.Value.(string)
		case System_oid[`sysLocation`]:
			res.Location = v.Value.(string)
		case System_oid[`sysContact`]:
			res.Contract = v.Value.(string)
		case System_oid[`sysObjectID`]:
			//TODO
		case System_oid[`sysUpTime`]:
			res.UpTime = v.Value.(int)
		case System_oid[`sysName`]:
			res.Hostname = v.Value.(string)
		}

	}
	return res, nil
}

const INTERFACE_IP_LIST_OID = `.1.3.6.1.2.1.4.20.1.2`

func (wrap *SNMPWrapper) Interfaces() ([]Interface, error) {
	//get interface index first
	prefix := `.1.3.6.1.2.1.2.2.1`
	index_str := prefix + `.1`
	res, err := wrap.Client.Walk(index_str)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	//get ip

	ip_str := INTERFACE_IP_LIST_OID
	ips, err := wrap.Client.Walk(ip_str)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	hash := make(map[int]string)
	for _, v := range ips {

		hash[v.Value.(int)] = v.Name[len(ip_str)+1:]
	}

	intfs := make([]Interface, 0)

	for _, v := range res {

		intf := Interface{}

		index := v.Value.(int)
		intf.IPv4 = hash[index]

		name_str := prefix + `.2.` + utils.Transform.IntToString(index)
		mac_str := prefix + `.6.` + utils.Transform.IntToString(index)
		status_str := prefix + `.8.` + utils.Transform.IntToString(index)
		mask_str := `.1.3.6.1.2.1.4.20.1.3.` + intf.IPv4

		var intfValues []string
		if len(intf.IPv4) != 0 {
			intfValues = []string{name_str, mac_str, status_str, mask_str}
		} else {
			intfValues = []string{name_str, mac_str, status_str}
		}

		res, err := wrap.Client.GetMulti(intfValues)
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}

		for _, v := range res.Variables {

			switch v.Name {
			case name_str:
				intf.Name = v.Value.(string)
			case mac_str:
				intf.MacAddress = v.Value.(string)
			case status_str:
				intf.Status = InterfaceStatus(v.Value.(int))
			case mask_str:
				mask := v.Value.(net.IP)
				intf.NetMask = mask.To4().String()
			}

		}

		intfs = append(intfs, intf)
	}

	return intfs, err
}
