// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package config

import (
	"strings"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	config, err := ReadConfigFile("./example.json")
	if err != nil {
		t.Error(err.Error())
	}
	str, err := config.ToString()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(str)
}

func TestReadConfig(t *testing.T) {
	json := `{
		"rpcConfig":[
			{
				"path":"auth.RoleService",
				"host": "192.168.1.146:8187",
				"methods":["Create","Update","List","Delete","UpdateGroup"]
				
			},
			{
				"path":"auth.UserService",
				"host": "192.168.1.146:8187",
				"methods":["List","ListGroup"]
			},
			{
				"path":"cmdb.NodeService",
				"host": "192.168.1.146:8583",
				"methods":["Create","Update","List","Search","Delete","Get"]
				
			},
			{
				"path":"cmdb.ModelService",
				"host": "192.168.1.146:8583",
				"methods":["Create","Update","List","Delete","Get","GetByLabel","BindingProperty","DeleteBindingProperty","GetBindingProperties"]
				
			}
		 ]
	}
	`
	config, err := ReadConfig(strings.NewReader(json))
	if err != nil {
		t.Error(err.Error())
	}
	str, err := config.ToString()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(str)
}
