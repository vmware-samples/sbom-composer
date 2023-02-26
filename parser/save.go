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

// TODO: Make configurable
var SPDX_VERSION = "2.2"

func Save(doc *Document, composableDocs []*Document, filters []string, output string, outFormat string) error {

	output = updateFileExtension(output, outFormat)

	w, err := os.Create(output)
	if err != nil {
		fmt.Printf("error while opening %v for writing: %v\n", output, err)
		return err
	}
	defer w.Close()

	// It's not necessary for the composed doc to
	// contain all merged documents as Files
	doc = cleanDocumentFileData(SPDX_VERSION, doc)

	composableDocs = filterComposableDocs(composableDocs, filters)

	updateRelationships(SPDX_VERSION, doc, composableDocs)

	for _, cdoc := range composableDocs {
		if cdoc != nil {
			AppendComposableDocument(SPDX_VERSION, doc, cdoc, filters, w, outFormat)
		}
	}

	err = SaveVersion(SPDX_VERSION, outFormat, doc, w)
	if err != nil {
		fmt.Printf("error while saving %v: %v\n", output, err)
		return err
	}
	return nil
}

func SaveVersion(version string, format string, doc *Document, w *os.File) error {
	switch format {
	case "tv":
		if version == "2.2" {
			return tvsaver.Save2_2(doc.SPDXDocRef.Doc2_2, w)
		}
	case "json":
		if version == "2.2" {
			return spdx_json.Save2_2(doc.SPDXDocRef.Doc2_2, w)
		}
	default:
		fmt.Printf("warn: %s is not proper output format; saving to default\n", format)
		if version == "2.2" {
			return tvsaver.Save2_2(doc.SPDXDocRef.Doc2_2, w)
		}
	}

	return nil
}

// RenderComposableDocument processes a composable document
// and renders it to the composed document
func AppendComposableDocument(spdxVersion string, res *Document, cdoc *Document, filters []string, w io.Writer, outFormat string) {

	switch spdxVersion {
	case "2.2":
		res.SPDXDocRef.Doc2_2.Annotations = append(res.SPDXDocRef.Doc2_2.Annotations, cdoc.SPDXDocRef.Doc2_2.Annotations...)
		res.SPDXDocRef.Doc2_2.ExternalDocumentReferences = append(res.SPDXDocRef.Doc2_2.ExternalDocumentReferences, cdoc.SPDXDocRef.Doc2_2.ExternalDocumentReferences...)
		res.SPDXDocRef.Doc2_2.Files = append(res.SPDXDocRef.Doc2_2.Files, cdoc.SPDXDocRef.Doc2_2.Files...)
		res.SPDXDocRef.Doc2_2.OtherLicenses = append(res.SPDXDocRef.Doc2_2.OtherLicenses, cdoc.SPDXDocRef.Doc2_2.OtherLicenses...)
		res.SPDXDocRef.Doc2_2.Packages = append(res.SPDXDocRef.Doc2_2.Packages, cdoc.SPDXDocRef.Doc2_2.Packages...)
		res.SPDXDocRef.Doc2_2.Relationships = append(res.SPDXDocRef.Doc2_2.Relationships, cdoc.SPDXDocRef.Doc2_2.Relationships...)
		res.SPDXDocRef.Doc2_2.Reviews = append(res.SPDXDocRef.Doc2_2.Reviews, cdoc.SPDXDocRef.Doc2_2.Reviews...)
		res.SPDXDocRef.Doc2_2.Snippets = append(res.SPDXDocRef.Doc2_2.Snippets, cdoc.SPDXDocRef.Doc2_2.Snippets...)
	}
}

func filterComposableDocs(cdocs []*Document, filters []string) []*Document {
	for i, c := range cdocs {
		cdocs[i].SPDXDocRef.Doc2_2.Packages = filterPackages(c.SPDXDocRef.Doc2_2.Packages, filters)
		cdocs[i].SPDXDocRef.Doc2_2.Relationships = filterRelationships(c.SPDXDocRef.Doc2_2.Relationships, filters)
	}
	return cdocs
}

func filterPackages(packages []*spdx.Package2_2, filters []string) []*spdx.Package2_2 {
	res := []*spdx.Package2_2{}

	identical := true
	for _, p := range packages {
		if contains(filters, p.PackageName) {
			identical = false
		} else {
			res = append(res, p)
		}
	}
	if !identical {
		return res
	}

	return packages
}

func filterRelationships(relationships []*spdx.Relationship2_2, filters []string) []*spdx.Relationship2_2 {
	res := []*spdx.Relationship2_2{}

	identical := true
	for _, r := range relationships {
		if containsSubstring(filters, string(r.RefA.ElementRefID)) ||
			containsSubstring(filters, string(r.RefB.ElementRefID)) {
			identical = false
		} else {
			res = append(res, r)
		}
	}
	if !identical {
		return res
	}

	return relationships
}

func cleanDocumentFileData(spdxVersion string, doc *Document) *Document {
	switch spdxVersion {
	case "2.2":
		doc.SPDXDocRef.Doc2_2.Files = []*spdx.File2_2{}

		for i := range doc.SPDXDocRef.Doc2_2.Packages {
			doc.SPDXDocRef.Doc2_2.Packages[i].Files = []*spdx.File2_2{}
		}
	}

	return doc
}

func updateRelationships(spdxVersion string, doc *Document, composableDocs []*Document) (*Document, []*Document) {

	rootDocElID := setDocElID(spdxVersion, doc)
	for _, cdoc := range composableDocs {
		switch spdxVersion {
		case "2.2":
			if cdoc != nil && len(cdoc.SPDXDocRef.Doc2_2.Packages) > 0 {
				elId := spdx.MakeDocElementID("",
					fmt.Sprintf("%s-%s", cdoc.SPDXDocRef.Doc2_2.Packages[0].PackageName, cdoc.SPDXDocRef.Doc2_2.Packages[0].PackageVersion))
				newRelationship := &spdx.Relationship2_2{
					RefA:         rootDocElID,
					RefB:         elId,
					Relationship: "DESCRIBES",
				}
				doc.SPDXDocRef.Doc2_2.Relationships = append(doc.SPDXDocRef.Doc2_2.Relationships, newRelationship)
			}
			if cdoc != nil && len(cdoc.SPDXDocRef.Doc2_2.Relationships) > 0 {
				cdoc.SPDXDocRef.Doc2_2.Relationships = cdoc.SPDXDocRef.Doc2_2.Relationships[1:]
			}
		}
	}

	return doc, composableDocs
}

func setDocElID(spdxVersion string, doc *Document) spdx.DocElementID {
	rootDocElID := spdx.DocElementID{}
	switch spdxVersion {
	case "2.2":
		if len(doc.SPDXDocRef.Doc2_2.Packages) > 0 {
			rootDocElID = spdx.MakeDocElementID("",
				fmt.Sprintf("%s-%s", doc.SPDXDocRef.Doc2_2.Packages[0].PackageName, doc.SPDXDocRef.Doc2_2.Packages[0].PackageVersion))
		}
	default:
		rootDocElID = spdx.MakeDocElementID("",
			fmt.Sprintf("%s-%s", doc.ConfigDataRef.PackageName, doc.ConfigDataRef.PackageVersion))
	}
	return rootDocElID
}
