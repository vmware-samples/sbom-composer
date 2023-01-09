// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"testing"

	"github.com/spdx/tools-golang/spdx"
	"github.com/stretchr/testify/assert"
)

func TestGenerateComposedDoc(t *testing.T) {
	fmt.Println("Testing Unmarshal JSON Config")
	t.Run("Unmarshal Config:", func(t *testing.T) {
		input := []byte(`spdxVersion: "SPDX-2.2"
documentName: "composed-1.0"
packageName: "top-level-artifact"
spdxID: "SPDXRef-DOCUMENT"
packageVersion: "1.0"
packageDownloadLocation: ""
filesAnalyzed: false
packageChecksum:
  sha256: "<checksum>"
packageLicenseConcluded: "BSD-3-Clause"
packageLicenseDeclared: "BSD-3-Clause"
packageCopyrightText: ""
packageSupplier: "somesupplier"
packageComment: "<text>somecomment</text>"`)

		want := Document{
			SPDXDocRef: &SPDXDocRef{
				Doc2_2: &spdx.Document2_2{
					DataLicense:       "CC0-1.0",
					SPDXIdentifier:    "DOCUMENT",
					DocumentName:      "top-level-artifact",
					DocumentNamespace: "https://spdx.org/spdxdocs/top-level-artifact-",
				}},
			ConfigDataRef: &Config{
				SPDXVersion:             "SPDX-2.2",
				SPDXConfigRef:           SPDXConfigReference,
				PackageName:             "top-level-artifact",
				DocumentName:            "composed-1.0",
				SPDXID:                  "SPDXRef-DOCUMENT",
				PackageVersion:          "1.0",
				PackageDownloadLocation: NOASSERTION,
				PackageChecksum: PackageChecksum{
					SHA256: "<checksum>",
				},
				FilesAnalyzed:           false,
				PackageLicenseConcluded: "BSD-3-Clause",
				PackageLicenseDeclared:  "BSD-3-Clause",
				PackageCopyrightText:    NOASSERTION,
				PackageSupplier:         "somesupplier",
				PackageComment:          "<text>somecomment</text>",
			},
		}

		loadedConfig := createConfig(input)
		doc, err := Build("../example_data/micro_sboms/tag_value", loadedConfig)
		assert.Equal(t, nil, err)

		assert.Equal(t, want.ConfigDataRef.SPDXVersion, doc.SPDXDocRef.Doc2_2.SPDXVersion)
		assert.Equal(t, want.SPDXDocRef.Doc2_2.DataLicense, doc.SPDXDocRef.Doc2_2.DataLicense)
		assert.Equal(t, want.SPDXDocRef.Doc2_2.SPDXIdentifier, doc.SPDXDocRef.Doc2_2.SPDXIdentifier)
		assert.Equal(t, want.SPDXDocRef.Doc2_2.DocumentName, doc.SPDXDocRef.Doc2_2.DocumentName)
		assert.Contains(t, doc.SPDXDocRef.Doc2_2.DocumentNamespace, want.SPDXDocRef.Doc2_2.DocumentNamespace)
		assert.Equal(t, want.ConfigDataRef.DocumentName, doc.ConfigDataRef.DocumentName)
		assert.Equal(t, want.ConfigDataRef.SPDXID, doc.ConfigDataRef.SPDXID)
		assert.Equal(t, want.ConfigDataRef.PackageVersion, doc.ConfigDataRef.PackageVersion)
		assert.Equal(t, want.ConfigDataRef.PackageLicenseConcluded, doc.ConfigDataRef.PackageLicenseConcluded)
	})
}
