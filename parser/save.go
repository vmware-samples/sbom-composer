// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"io"
	"os"

	spdx_json "github.com/spdx/tools-golang/json"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/tvsaver"
)

func Save(doc *Document, composableDocs []*Document, output string, outFormat string) error {

	output = updateFileExtension(output, outFormat)

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

	for _, cdoc := range composableDocs {
		if cdoc != nil {
			AppendComposableDocument(doc, cdoc, w, outFormat)
		}
	}

	switch outFormat {
	case "tv":
		err = tvsaver.Save2_2(doc.SPDXDocRef, w)
	case "json":
		err = spdx_json.Save2_2(doc.SPDXDocRef, w)
	default:
		fmt.Printf("warn: %s is not proper output format; saving to default\n", outFormat)
		err = tvsaver.Save2_2(doc.SPDXDocRef, w)
	}
	if err != nil {
		fmt.Printf("error while saving %v: %v\n", output, err)
		return err
	}
	return nil
}

// RenderComposableDocument processes a composable document
// and renders it to the composed document
func AppendComposableDocument(res *Document, cdoc *Document, w io.Writer, outFormat string) {

	res.SPDXDocRef.Annotations = append(res.SPDXDocRef.Annotations, cdoc.SPDXDocRef.Annotations...)

	for _, e := range cdoc.SPDXDocRef.ExternalDocumentReferences {
		if isNotDuplicate(e.Checksum.Value, res, EDR) {
			res.SPDXDocRef.ExternalDocumentReferences = append(res.SPDXDocRef.ExternalDocumentReferences, e)
		}
	}
	for _, f := range cdoc.SPDXDocRef.Files {
		if areNotIdenticalChecksums(f.Checksums, res, FL) {
			res.SPDXDocRef.Files = append(res.SPDXDocRef.Files, f)
		}
	}
	for _, ol := range cdoc.SPDXDocRef.OtherLicenses {
		if isNotDuplicate(ol.LicenseIdentifier, res, OL) {
			res.SPDXDocRef.OtherLicenses = append(res.SPDXDocRef.OtherLicenses, ol)
		}
	}
	for _, p := range cdoc.SPDXDocRef.Packages {
		if areNotIdenticalChecksums(p.PackageChecksums, res, PKG) {
			res.SPDXDocRef.Packages = append(res.SPDXDocRef.Packages, p)
		}
	}
	for _, r := range cdoc.SPDXDocRef.Relationships {
		relStr := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s_%s",
			r.RefA.DocumentRefID, r.RefA.ElementRefID, r.RefA.SpecialID,
			r.Relationship,
			r.RefB.DocumentRefID, r.RefB.ElementRefID, r.RefB.SpecialID,
			r.RelationshipComment)
		if isNotDuplicate(relStr, res, RL) {
			res.SPDXDocRef.Relationships = append(res.SPDXDocRef.Relationships, r)
		}
	}
	res.SPDXDocRef.Reviews = append(res.SPDXDocRef.Reviews, cdoc.SPDXDocRef.Reviews...)
	res.SPDXDocRef.Snippets = append(res.SPDXDocRef.Snippets, cdoc.SPDXDocRef.Snippets...)
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
		if cdoc != nil && len(cdoc.SPDXDocRef.Packages) > 0 {
			elId := spdx.MakeDocElementID("",
				fmt.Sprintf("%s-%s", cdoc.SPDXDocRef.Packages[0].PackageName, cdoc.SPDXDocRef.Packages[0].PackageVersion))
			newRelationship := &spdx.Relationship2_2{
				RefA:         rootDocElID,
				RefB:         elId,
				Relationship: "DESCRIBES",
			}
			doc.SPDXDocRef.Relationships = append(doc.SPDXDocRef.Relationships, newRelationship)
		}
		if cdoc != nil && len(cdoc.SPDXDocRef.Relationships) > 0 {
			cdoc.SPDXDocRef.Relationships = cdoc.SPDXDocRef.Relationships[1:]
		}
	}

	return doc, composableDocs
}
