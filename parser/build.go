// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"

	"github.com/spdx/tools-golang/builder"
)

func Build(spdxVersion string, dirRoot string, conf *Config) (*Document, error) {
	spdxDocRef := BuildVersion(spdxVersion, dirRoot, conf)

	UpdatePackages(SPDX_VERSION, &spdxDocRef, conf)
	doc := CreateDocument(&spdxDocRef, conf)
	return doc, nil
}

func BuildVersion(spdxVersion string, dirRoot string, conf *Config) SPDXDocRef {
	res := SPDXDocRef{}
	switch spdxVersion {
	case "2.2":
		var err error
		res.Doc2_2, err = builder.Build2_2(conf.PackageName, dirRoot, conf.SPDXConfigRef)
		if err != nil {
			fmt.Printf("error while building spdx document reference for path %v with config %v, %v: %v\n", dirRoot, conf.PackageName, conf.SPDXConfigRef, err)
		}
	}
	return res
}

func GenerateComposedDoc(dirRoot string, output string, outFormat string, confFile string) error {
	conf := LoadConfig(confFile)

	doc, err := Build(SPDX_VERSION, dirRoot, conf)
	if err != nil {
		return err
	}

	composableDocs := LoadAll(dirRoot)

	err = Save(doc, composableDocs, output, outFormat)
	if err != nil {
		fmt.Printf("failed to save composed document %v: %v", output, err)
		return err
	}
	return nil
}
