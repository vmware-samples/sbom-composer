// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import "github.com/spdx/tools-golang/spdx"

type Document struct {
	SPDXDocRef    *SPDXDocRef
	ConfigDataRef *Config
}
type SPDXDocRef struct {
	Doc2_2 *spdx.Document2_2
}

func CreateDocument(spdxDocRef *SPDXDocRef, conf *Config) *Document {
	doc := &Document{
		SPDXDocRef:    spdxDocRef,
		ConfigDataRef: conf,
	}
	return doc
}

func CreateDocumentWithSPDXRef() *Document {
	spdxDocRef := &SPDXDocRef{Doc2_2: &spdx.Document2_2{}}
	doc := &Document{
		SPDXDocRef: spdxDocRef,
	}
	return doc
}

func UpdatePackages(spdxVersion string, spdxDocRef *SPDXDocRef, conf *Config) {
	switch spdxVersion {
	case "2.2":
		for i := range spdxDocRef.Doc2_2.Packages {
			if spdxDocRef.Doc2_2.Packages[i].PackageName == conf.PackageName &&
				len(spdxDocRef.Doc2_2.Packages[i].PackageVersion) == 0 {
				spdxDocRef.Doc2_2.Packages[i].PackageVersion = conf.PackageVersion
			}
		}
	}
}
