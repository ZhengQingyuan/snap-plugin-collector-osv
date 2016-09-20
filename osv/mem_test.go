//
// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package osv

import (
	"strconv"
	"testing"

	"github.com/intelsdi-x/snap-plugin-collector-osv/osv/httpmock"

	"github.com/intelsdi-x/snap/core"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemPlugin(t *testing.T) {
	httpmock.Mock = true

	Convey("getMemstat Should return memory amount value", t, func() {

		defer httpmock.ResetResponders()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/free", "20000", 200)
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/total", "10000", 200)

		memFree, err := getMemStat("http://192.168.192.200:8000", "free")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(memFree, 10), ShouldResemble, "20000")
		memTotal, err := getMemStat("http://192.168.192.200:8000", "total")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(memTotal, 10), ShouldResemble, "10000")

	})
	Convey("MemStat Should return pluginMetricType Data", t, func() {

		defer httpmock.ResetResponders()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/free", "20000", 200)
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/total", "10000", 200)

		ns := core.NewNamespace("intel", "osv", "memory", "free")
		ns2 := core.NewNamespace("intel", "osv", "memory", "total")
		memFree, err := memStat(ns, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(memFree.Namespace(), ShouldResemble, ns)
		So(memFree.Data_, ShouldResemble, "20000")
		memTotal, err := memStat(ns2, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(memTotal.Namespace(), ShouldResemble, ns2)
		So(memTotal.Data_, ShouldResemble, "10000")

	})
}
