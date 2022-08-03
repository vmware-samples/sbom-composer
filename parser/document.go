// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import "github.com/spdx/tools-golang/spdx"

type Document struct {
	SPDXDocRef    *spdx.Document2_2
	ConfigDataRef *Config
}
