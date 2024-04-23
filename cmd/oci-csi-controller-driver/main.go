// Copyright 2019 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"runtime"

	"github.com/oracle/oci-cloud-controller-manager/cmd/oci-csi-controller-driver/csi-controller"
	"github.com/oracle/oci-cloud-controller-manager/cmd/oci-csi-controller-driver/csioptions"
	"github.com/oracle/oci-cloud-controller-manager/pkg/util/signals"

	"bitbucket.oci.oraclecorp.com/cryptography/go_ensurefips"
)

func main() {
	// Ensure AMD service is FIPS Compliant
	if runtime.GOARCH == "amd64" {
		go_ensurefips.Compliant()
	}

	csioptions := csioptions.NewCSIOptions()
	flag.Parse()
	stopCh := signals.SetupSignalHandler()
	csicontroller.Run(*csioptions, stopCh)
}
