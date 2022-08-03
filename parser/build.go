// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"os"

	"github.com/spdx/tools-golang/builder"
	"github.com/spdx/tools-golang/tvsaver"
)

func Build(dirRoot string, conf *Config) (*Document, error) {
	spdxDocRef, err := builder.Build2_2(conf.PackageName, dirRoot, conf.SPDXConfigRef)
	if err != nil {
		fmt.Printf("error while building spdx document reference for path %v with config %v, %v: %v\n", dirRoot, conf.PackageName, conf.SPDXConfigRef, err)
	}
	fmt.Printf("HERE: %+v\n", spdxDocRef.DocumentName)
	doc := &Document{
		SPDXDocRef:    spdxDocRef,
		ConfigDataRef: conf,
	}
	return doc, nil
}

func Save(doc *Document, composedDoc string) error {
	w, err := os.Create(composedDoc)
	if err != nil {
		fmt.Printf("error while opening %v for writing: %v\n", composedDoc, err)
		return err
	}
	defer w.Close()

	err = tvsaver.Save2_2(doc.SPDXDocRef, w)
	if err != nil {
		fmt.Printf("error while saving %v: %v", composedDoc, err)
		return err
	}
	return nil
}

func GenerateComposedDoc(dirRoot string, composedDoc string, confFile string) error {
	conf := LoadConfig(confFile)

	doc, err := Build(dirRoot, conf)
	if err != nil {
		return err
	}

	err = Save(doc, composedDoc)
	if err != nil {
		fmt.Printf("failed to save composed document %v: %v", composedDoc, err)
		return err
	}
	return nil
}
