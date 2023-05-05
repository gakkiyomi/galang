// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/beevik/etree"
	str "github.com/gakkiyomi/galang/string"
	"github.com/songtianyi/rrframework/logs"
)

type GalangFile byte

var File GalangFile

// GetFileName getFileName from a full filePath
func (*GalangFile) GetFileName(filePath string) (name string) {
	name = ""
	if str.String.IsBlank(filePath) {
		return
	}
	name = path.Base(filePath)
	return
}

// GetFileSize get the length in bytes of file of the specified path.
func (*GalangFile) GetFileSize(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		logs.Error(err)
		return -1
	}
	return file.Size()
}

// IsJSONFile check one file is json file
func (file *GalangFile) IsJSONFile(path string) bool {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Error(err)
		return false
	}
	return file.IsJSONByte(b)
}

// IsJSONByte check one []byte is json file
func (*GalangFile) IsJSONByte(b []byte) bool {
	return json.Valid(b)
}

// IsXmlFile check one file is xml file
func (file *GalangFile) IsXmlFile(path string) bool {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	return file.IsXmlByte(b)
}

// IsXmlByte check one []byte is xml file
func (*GalangFile) IsXmlByte(b []byte) bool {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		return false
	}
	return true
}
