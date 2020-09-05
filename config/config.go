// Galang - Golang common utilities
// Copyright (c) 2019-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"std/file"
	"strings"
	"text/template"

	"github.com/beevik/etree"
)

type Config interface {
	ToString() (string, error)
	GetBytes() []byte
}

type JsonConfig struct {
	Json  map[string]interface{}
	Jsonb []byte
}

type XmlConfig struct {
	Xml  *etree.Document
	Xmlb []byte
}

func ReadConfigFile(path string) (Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if file.File.IsJSONByte(b) {
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		result := &JsonConfig{
			Json:  m,
			Jsonb: b,
		}
		return result, nil
	}

	doc := etree.NewDocument()
	if file.File.IsXmlByte(b) {
		result := &XmlConfig{
			Xml:  doc,
			Xmlb: b,
		}
		return result, nil
	}

	return nil, errors.New("only supoort json and xml config file")
}

func (jc *JsonConfig) ToString() (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, jc.Jsonb, "", "\t"); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func (jc *JsonConfig) GetBytes() []byte {
	return jc.Jsonb
}

func (m *XmlConfig) ToString() (string, error) {
	var buf bytes.Buffer
	template.HTMLEscape(&buf, m.Xmlb)
	str := string(buf.Bytes())
	str = strings.Replace(str, "&lt;", "<", -1)
	str = strings.Replace(str, "&gt;", ">", -1)
	str = strings.Replace(str, "&#34;", "\"", -1)
	return str, nil
}

func (xml *XmlConfig) GetBytes() []byte {
	return xml.Xmlb
}
