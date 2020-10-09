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
	"fmt"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/daba0007/golib/tools"
	"github.com/gakkiyomi/galang/utils"
	"github.com/songtianyi/rrframework/logs"
)

type ScannerStatus int32

const (
	Created ScannerStatus = 1
	Pending ScannerStatus = 2
	Running ScannerStatus = 3
	Success ScannerStatus = 4
	Error   ScannerStatus = 5
)

var ScannerStatus_name = map[int32]string{
	1: "Created",
	2: "Pending",
	3: "Running",
	4: "Success",
	5: "Error",
}
var ScannerStatus_value = map[string]int32{
	"Created": 1,
	"Pending": 2,
	"Running": 3,
	"Success": 4,
	"Error":   5,
}

func (x ScannerStatus) String() string {
	s, ok := ScannerStatus_name[int32(x)]
	if ok {
		return s
	}
	return utils.Transform.AnyToString(x)
}

type Scanner struct {
	Id        string        `json:"id"`
	Targets   []string      `json:"targets"`
	StartTime time.Time     `json:"start_at"`
	EndTime   time.Time     `json:"end_at"`
	Status    ScannerStatus `json:"status"`
	Message   string        `json:"message"`
	User      string        `json:"user"`
	Result    nmap.Run      `json:"result"`
	Warnings  []string      `json:"warnings"`
}

//NewScanner 新建一个扫描器
func (*GalangNMAP) NewScanner(target, user string) (*Scanner, error) {
	var targets []string
	subnetInfo, err := NewSubnetInfo(target)
	if err != nil {
		//说明不是一个网段格式 可能是一个IP
		ips, err2 := Network.IPSFormat(target)
		if err2 != nil {
			logs.Error(err2.Error())
			return nil, err2
		}
		targets = ips
	} else {
		targets = append(targets, subnetInfo.GetCidrSignature())
	}
	idGottor, _ := tools.NewSnowflake(0, 0)
	id, _ := idGottor.NextVal()
	idStr := utils.Transform.AnyToString(id)
	res := &Scanner{
		Id:      idStr,
		Targets: targets,
		User:    user,
		Status:  Created,
	}
	return res, nil
}

//Scanner will scan target ip or subnet
func (sc *Scanner) Scanner(options ...func(*nmap.Scanner)) (*Scanner, error) {

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(sc.Targets...),
	)

	for _, option := range options {
		option(scanner)
	}

	if err != nil {
		logs.Error("unable to create nmap scanner: %v", err)
		return nil, err
	}
	sc.StartTime = time.Now()
	result, warnings, err := scanner.Run()
	sc.EndTime = time.Now()

	if err != nil {
		logs.Error("unable to run nmap scan: %v", err)
		sc.Status = Error
		sc.Message = fmt.Sprintf("unable to run nmap scan: %v", err)
	} else {
		sc.Result = *result
		sc.Message = "ok"
		sc.Status = Success
	}

	if warnings != nil {
		logs.Warn("Warnings: \n %v", warnings)
		sc.Warnings = warnings
	}

	return sc, nil
}
