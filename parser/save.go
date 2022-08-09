// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvsaver"
)

func Save(doc *Document, composableDocs []*Document, output string) error {
	w, err := os.Create(output)
	if err != nil {
		fmt.Printf("error while opening %v for writing: %v\n", output, err)
		return err
	}
	defer w.Close()

	// It's not necessary for the composed doc to
	// contain all merged documents as Files
	doc = cleanDocumentFileData(doc)

	updateRelationships(doc, composableDocs)

	err = tvsaver.Save2_2(doc.SPDXDocRef, w)
	if err != nil {
		fmt.Printf("error while saving %v: %v", output, err)
		return err
	}

	for _, cdoc := range composableDocs {
		RenderComposableDocument(cdoc, w)
	}
	return nil
}

// RenderComposableDocument processes a composable document
// and renders it to the composed document
func RenderComposableDocument(doc *Document, w io.Writer) error {

	// Merged documents should not contain any head data. It's
	// only generated for the composed document
	doc = cleanDocumentHeadData(doc)

	err := tvsaver.Save2_2(doc.SPDXDocRef, w)
	if err != nil {
		fmt.Printf("error while saving doc %v: %v", doc, err)
		return err
	}

	return nil
}

func cleanDocumentHeadData(doc *Document) *Document {
	doc.SPDXDocRef.SPDXVersion = ""
	doc.SPDXDocRef.DataLicense = ""
	doc.SPDXDocRef.SPDXIdentifier = ""
	doc.SPDXDocRef.DocumentName = ""
	doc.SPDXDocRef.DocumentNamespace = ""
	doc.SPDXDocRef.DocumentComment = ""
	doc.SPDXDocRef.CreationInfo.LicenseListVersion = ""
	doc.SPDXDocRef.CreationInfo.Creators = []spdx.Creator{}
	doc.SPDXDocRef.CreationInfo.CreatorComment = ""
	doc.SPDXDocRef.CreationInfo.Created = ""

	return doc
}

func cleanDocumentFileData(doc *Document) *Document {
	doc.SPDXDocRef.Files = []*spdx.File2_2{}

	for i := range doc.SPDXDocRef.Packages {
		doc.SPDXDocRef.Packages[i].Files = []*spdx.File2_2{}
	}

	return doc
}

func updateRelationships(doc *Document, composableDocs []*Document) (*Document, []*Document) {

	rootDocElID := spdx.DocElementID{}
	if len(doc.SPDXDocRef.Packages) > 0 {
		rootDocElID = spdx.MakeDocElementID("",
			fmt.Sprintf("%s-%s", doc.SPDXDocRef.Packages[0].PackageName, doc.SPDXDocRef.Packages[0].PackageVersion))
	} else {
		rootDocElID = spdx.MakeDocElementID("",
			fmt.Sprintf("%s-%s", doc.ConfigDataRef.PackageName, doc.ConfigDataRef.PackageVersion))
	}
	for _, cdoc := range composableDocs {
		if len(cdoc.SPDXDocRef.Packages) > 0 {
			elId := spdx.MakeDocElementID("",
				fmt.Sprintf("%s-%s", cdoc.SPDXDocRef.Packages[0].PackageName, cdoc.SPDXDocRef.Packages[0].PackageVersion))
			newRelationship := &spdx.Relationship2_2{
				RefA:         rootDocElID,
				RefB:         elId,
				Relationship: "DESCRIBES",
			}
			doc.SPDXDocRef.Relationships = append(doc.SPDXDocRef.Relationships, newRelationship)
		}
		if len(cdoc.SPDXDocRef.Relationships) > 0 {
			cdoc.SPDXDocRef.Relationships = cdoc.SPDXDocRef.Relationships[1:]
		}
	}

	return doc, composableDocs
}
