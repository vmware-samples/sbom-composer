// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigYAML(t *testing.T) {
	fmt.Println("Testing Loading YAML Config")
	t.Run("Loading YAML Config:", func(t *testing.T) {

		conf := LoadConfig("../config/example_config.yaml")

		assert.Equal(t, "SPDX-2.2", conf.SPDXVersion)
		assert.Equal(t, "composed-1.0", conf.DocumentName)
		assert.Equal(t, "top-level-artifact", conf.PackageName)
		assert.Equal(t, "SPDXRef-DOCUMENT", conf.SPDXID)
		assert.Equal(t, "1.0", conf.PackageVersion)
		assert.Equal(t, NOASSERTION, conf.PackageDownloadLocation)
		assert.Equal(t, "<checksum>", conf.PackageChecksum.SHA256)
		assert.Equal(t, false, conf.FilesAnalyzed)
		assert.Equal(t, NOASSERTION, conf.PackageLicenseConcluded)
		assert.Equal(t, NOASSERTION, conf.PackageLicenseDeclared)
		assert.Equal(t, NOASSERTION, conf.PackageCopyrightText)
		assert.Equal(t, "Example supplier", conf.PackageSupplier)
		assert.Equal(t, "<text>Example comment</text>", conf.PackageComment)

	})
}

func TestLoadConfigJSON(t *testing.T) {
	fmt.Println("Testing Loading JSON Config")
	t.Run("Loading JSON Config:", func(t *testing.T) {

		conf := LoadConfig("../config/example_config.json")
		fmt.Println(conf)
		assert.Equal(t, "composed-1.0", conf.DocumentName)
		assert.Equal(t, "top-level-artifact", conf.PackageName)
		assert.Equal(t, "SPDXRef-composed-sbom-product", conf.SPDXID)
		assert.Equal(t, "1.0", conf.PackageVersion)
		assert.Equal(t, "Input from user or NOASSERTION", conf.PackageDownloadLocation)
		assert.Equal(t, "<checksum>", conf.PackageChecksum.SHA256)
		assert.Equal(t, false, conf.FilesAnalyzed)
		assert.Equal(t, "license, licenseRef or NOASSERTION", conf.PackageLicenseConcluded)
		assert.Equal(t, "license, licenseRef or NOASSERTION", conf.PackageLicenseDeclared)
		assert.Equal(t, "text or NOASSERTION", conf.PackageCopyrightText)
		assert.Equal(t, "Organization or recognized author of product. Optional", conf.PackageSupplier)
		assert.Equal(t, "<text>Any relevant comment</text>. Optional", conf.PackageComment)

	})
}
