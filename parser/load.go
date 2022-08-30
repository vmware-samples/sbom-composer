// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"os"

	spdx_json "github.com/spdx/tools-golang/json"
	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
)

func LoadFile(file string) *Document {
	r, err := os.Open(file)
	if err != nil {
		fmt.Printf("error  opening %v: %v", file, err)
		return nil
	}
	defer r.Close()

	res := &Document{}

	doc := &Document{}

	if isJSON(file) {
		doc.SPDXDocRef, err = spdx_json.Load2_2(r)
	} else {
		doc.SPDXDocRef, err = tvloader.Load2_2(r)
	}

	if err != nil {
		fmt.Printf("error parsing %v: %v\n", file, err)
		return nil
	}

	// verify if the SPDX file describes at least one package
	pkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(doc.SPDXDocRef)
	if err != nil {
		fmt.Printf("couldn't find package description in the SPDX document: %v\n", err)
		return nil
	}

	if len(pkgIDs) == 0 {
		return nil
	}

	res.SPDXDocRef = doc.SPDXDocRef

	return res
}

func LoadAll(dir string) []*Document {

	allSBOMs := []string{}
	if len(dir) > 0 {
		allSBOMs = readDir(dir)
	}

	var loaded []*Document

	for _, doc := range allSBOMs {
		loaded = append(loaded, LoadFile(dir+"/"+doc))
	}
	return loaded
}
