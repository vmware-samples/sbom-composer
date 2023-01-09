// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateJSONTConfigExample(spdxVersion string, documentName string, packageName string, spdxID string, packageVersion string,
	packageDownloadLocation string, packageChecksum PackageChecksum, filesAnalyzed bool, packageLicenseConcluded string,
	packageLicenseDeclared string, packageCopyrightText string, packageSupplier string, packageComment string) []byte {
	jsonStr := fmt.Sprintf("{\"spdxVersion\":\"%s\",\"documentName\":\"%s\",\"packageName\":\"%s\",\"spdxID\":\"%s\",\"packageVersion\":\"%s\",\"packageDownloadLocation\":\"%s\",\"filesAnalyzed\":\"%t\",\"packageChecksum\":\"%s\", \"packageLicenseConcluded\":\"%s\", \"packageLicenseDeclared\":\"%s\", \"packageCopyrightText\":\"%s\", \"packageSupplier\":\"%s\", \"packageComment\":\"%s\"}",
		spdxVersion, documentName, packageName, spdxID, packageVersion, packageDownloadLocation, filesAnalyzed, packageChecksum,
		packageLicenseConcluded, packageLicenseDeclared, packageCopyrightText, packageSupplier, packageComment)

	return []byte(jsonStr)
}

func generateSHA256(data string) string {
	h := sha256.New()
	h.Write([]byte("somedata"))
	sha := h.Sum(nil)
	checksum := hex.EncodeToString(sha)
	return checksum
}

func TestUnmarshalJSONConfig(t *testing.T) {
	fmt.Println("Testing Unmarshal JSON Config")
	t.Run("Unmarshal Config:", func(t *testing.T) {
		checksum := generateSHA256("somedata")

		pch := PackageChecksum{SHA256: checksum}

		jsonData := generateJSONTConfigExample("SPDX-2.2", "composed-1.0", "top-level-artifact", "SPDXRef-DOCUMENT", "1.0", NOASSERTION, pch,
			false, "BSD-3-Clause", "BSD-3-Clause", NOASSERTION,
			"somesupplier", "somecomment")
		conf, _ := UnmarshalJSONConfig(jsonData)

		assert.Equal(t, "SPDX-2.2", conf.SPDXVersion)
		assert.Equal(t, "composed-1.0", conf.DocumentName)
		assert.Equal(t, "top-level-artifact", conf.PackageName)
		assert.Equal(t, "SPDXRef-DOCUMENT", conf.SPDXID)
		assert.Equal(t, "1.0", conf.PackageVersion)
		assert.Equal(t, NOASSERTION, conf.PackageDownloadLocation)
		assert.Equal(t, checksum, conf.PackageChecksum.SHA256)
		assert.Equal(t, false, conf.FilesAnalyzed)
		assert.Equal(t, "BSD-3-Clause", conf.PackageLicenseConcluded)
		assert.Equal(t, "BSD-3-Clause", conf.PackageLicenseDeclared)
		assert.Equal(t, NOASSERTION, conf.PackageCopyrightText)
		assert.Equal(t, "somesupplier", conf.PackageSupplier)
		assert.Equal(t, "somecomment", conf.PackageComment)
	})
}
