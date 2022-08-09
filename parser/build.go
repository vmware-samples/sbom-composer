// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"

	"github.com/spdx/tools-golang/builder"
)

func Build(dirRoot string, conf *Config) (*Document, error) {
	spdxDocRef, err := builder.Build2_2(conf.PackageName, dirRoot, conf.SPDXConfigRef)
	if err != nil {
		fmt.Printf("error while building spdx document reference for path %v with config %v, %v: %v\n", dirRoot, conf.PackageName, conf.SPDXConfigRef, err)
	}

	for i := range spdxDocRef.Packages {
		if spdxDocRef.Packages[i].PackageName == conf.PackageName &&
			len(spdxDocRef.Packages[i].PackageVersion) == 0 {
			spdxDocRef.Packages[i].PackageVersion = conf.PackageVersion
		}
	}
	doc := &Document{
		SPDXDocRef:    spdxDocRef,
		ConfigDataRef: conf,
	}
	return doc, nil
}

func GenerateComposedDoc(dirRoot string, output string, confFile string) error {
	conf := LoadConfig(confFile)

	doc, err := Build(dirRoot, conf)
	if err != nil {
		return err
	}

	composableDocs := LoadAll(dirRoot)

	err = Save(doc, composableDocs, output)
	if err != nil {
		fmt.Printf("failed to save composed document %v: %v", output, err)
		return err
	}
	return nil
}
