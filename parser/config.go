// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spdx/tools-golang/builder"
	"gopkg.in/yaml.v3"
)

// PackageChecksum is a unique identifier used to verify if all files
// in the orginal package are unchanged, exluding SPDX documents.
// Should be generated with SHA256.
type PackageChecksum struct {
	SHA256 string
}

var NOASSERTION string = "NOASSERTION"
var SPDXConfigReference *builder.Config2_2 = &builder.Config2_2{
	NamespacePrefix: "https://spdx.org/spdxdocs/", // TODO: move this to config
	CreatorType:     "Tool",
	Creator:         "sbom-composer-1.0", // TODO: automate taking the version
}

// Config is a collection of configuration settings for builder
// to create a composed document with.
type Config struct {
	SPDXConfigRef *builder.Config2_2

	// DocumentName is an SBOM-Composer report
	// for <top level product name>
	DocumentName string `json:"documentName"`

	// PackageName is the name of the
	// top level composed SBOM product
	PackageName string `json:"packageName"`

	// SPDXID is an ID of the type:
	// SPDXRef-composed-sbom-product
	SPDXID string `json:"spdxID"`

	// PackageVersion correspond to the composer's version
	PackageVersion string `json:"packageVersion"`

	// PackageDownloadLocation is the location the package is
	// downloaded to if specified, or NOASSERTION otherwise
	PackageDownloadLocation string `json:"packageDownloadLocation"`

	// PackageChecksum is a PackageChecksum type representation
	PackageChecksum PackageChecksum `json:"packageChecksum"`

	// FilesAnalyzed indicates if package level analysis has been performed
	FilesAnalyzed bool `json:"filesAnalyzed"`

	// PackageLicenseConcluded is either a license, licenseRef or NOASSERTION
	PackageLicenseConcluded string `json:"packageLicenseConcluded"`

	// PackageLicenseDeclared is either a license, licenseRef or NOASSERTION
	PackageLicenseDeclared string `json:"packageLicenseDeclared"`

	// PackageCopyrightText contains copyright text or NOASSERTION
	PackageCopyrightText string `json:"packageCopyrightText"`

	// PackageSupplier refers to an organization
	// or recognized author of product
	PackageSupplier string `json:"packageSupplier,omitempty"`

	// PackageComment is any relevant comment in a <text></text> section
	PackageComment string `json:"packageComment,omitempty"`
}

func UnmarshalJSONConfig(jsonData []byte) (*Config, error) {
	var c *Config
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println("unmarshal failed. continuing to update individual field\n", err)
	}

	var objmap map[string]json.RawMessage
	err = json.Unmarshal(jsonData, &objmap)
	if err != nil {
		fmt.Println("object-map unmarshal failed.\n", err)
	}
	c.PackageChecksum.SHA256 = strings.Trim(string(objmap["packageChecksum"]), "\"{}")

	return c, err
}

func LoadConfig(file string) *Config {

	conf := readConfFile(file)

	return createConfig(conf)
}

func readConfFile(file string) []byte {

	conf, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("error: failed reading yaml file. SPDX document not generated.", err)
		os.Exit(1)
	}
	return conf
}

func createConfig(loadedData []byte) *Config {
	conf := Config{}

	conf.SPDXConfigRef = SPDXConfigReference

	mapData := make(map[string]interface{})

	err := yaml.Unmarshal(loadedData, &mapData)
	if err != nil {
		fmt.Println("unmarshal failed after reading yaml file\n", err)
	}

	dataByte, _ := json.Marshal(mapData)
	_ = json.Unmarshal(dataByte, &conf)

	if len(conf.PackageDownloadLocation) == 0 {
		conf.PackageDownloadLocation = NOASSERTION
	}
	if len(conf.PackageLicenseConcluded) == 0 {
		conf.PackageLicenseConcluded = NOASSERTION
	}
	if len(conf.PackageLicenseDeclared) == 0 {
		conf.PackageLicenseDeclared = NOASSERTION
	}
	if len(conf.PackageCopyrightText) == 0 {
		conf.PackageCopyrightText = NOASSERTION
	}
	conf.PackageChecksum.SHA256 = strings.Trim(conf.PackageChecksum.SHA256, "\"{}")

	return &conf
}
