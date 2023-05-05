// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package utils

import (
	uuid "github.com/satori/go.uuid"
)

// NewUUID return a random generated UUID
func (*GalangUUID) NewUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}

// FromString return a UUID instance from string
func (*GalangUUID) FromString(str string) (uuid.UUID, error) {
	return uuid.FromString(str)
}

// Check check input is valid UUID String
func (id *GalangUUID) Check(str string) bool {
	_, err := id.FromString(str)
	if err != nil {
		return false
	}
	return true
}
