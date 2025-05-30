/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package test

import (
	"github.com/stretchr/testify/assert"
	"mep-agent/src/util"
	"os"
	"testing"
)

func TestClearByteArray(t *testing.T) {
	data1 := []byte{'a', 'b', 'c'}
	util.ClearByteArray(data1)
	data2 := []byte{0, 0, 0}
	assert.Equal(t, data2, data1)

	util.ClearByteArray(nil)

}

func TestReadTokenFromEnvironment1(t *testing.T) {
	os.Setenv("ak", "ZXhhbXBsZUFL")
	os.Setenv("sk", "ZXhhbXBsZVNL")
	err := util.ReadTokenFromEnvironment()
	assert.EqualValues(t, 0, len(os.Getenv("ak")))
	assert.EqualValues(t, 0, len(os.Getenv("sk")))
	assert.NoError(t, err, "No error is expected")
}

func TestReadTokenFromEnvironment2(t *testing.T) {
	os.Setenv("ak", "ZXhhbXBsZUFL")
	err := util.ReadTokenFromEnvironment()
	Expected := "ak and sk keys should be set in env variable"
	assert.EqualError(t, err, Expected)
}


func TestGetAppInstanceIdDecodeFailed(t *testing.T) {
	os.Setenv("APPINSTID", "b1fe5b4d-76a7-4a52-b60f-932fde7c8d57")
	_, err := util.GetAppInstanceID()
	assert.Equal(t, err, nil)
}

func TestGetAppInstanceIdNotSet(t *testing.T) {
	_, err := util.GetAppInstanceID()
	Expected := "appInstanceId should be set in env variable"
	assert.EqualError(t, err, Expected)
}
