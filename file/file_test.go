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
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileName(t *testing.T) {
	a := File.GetFileName("a/b/c/ac/fangcong.txt")
	b := File.GetFileName("")

	assert.Equal(t, "fangcong.txt", a)
	assert.Equal(t, "", b)

}

func TestIsJSONFile(t *testing.T) {
	a := File.IsJSONFile("../examples/file/demo.json")
	b := File.IsJSONFile("../examples/file/demo.xml")
	assert.Equal(t, true, a)
	assert.Equal(t, false, b)
}

func TestIsXmlFile(t *testing.T) {
	b := File.IsXmlFile("../examples/file/demo.xml")
	assert.Equal(t, true, b)
}

func TestIsJSONByte(t *testing.T) {
	byte, _ := ioutil.ReadFile("../examples/file/demo.json")
	byte2, _ := ioutil.ReadFile("../examples/file/demo.xml")
	res := File.IsJSONByte(byte)
	assert.Equal(t, true, res)

	res2 := File.IsJSONByte(byte2)
	assert.Equal(t, false, res2)
}

func TestIsXmlByte(t *testing.T) {
	byte, _ := ioutil.ReadFile("../examples/file/demo.xml")
	res := File.IsXmlByte(byte)
	assert.Equal(t, true, res)
}
