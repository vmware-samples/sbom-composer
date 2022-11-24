// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

const (
	EDR string = "external_document_reference"
	FL         = "file"
	OL         = "other_license"
	PKG        = "package"
	RL         = "relationship"
)

func areNotIdenticalChecksums(checksums []spdx.Checksum, doc *Document, t string) bool {

	switch t {
	case "file":
		for _, f := range doc.SPDXDocRef.Files {
			if len(checksums) > 0 && reflect.DeepEqual(checksums, f.Checksums) {
				return false
			}
		}

	case "package":
		for _, p := range doc.SPDXDocRef.Packages {
			if len(checksums) > 0 && reflect.DeepEqual(checksums, p.PackageChecksums) {
				return false
			}
		}
	default:
		return true
	}
	return true
}

func isNotDuplicate(data string, doc *Document, t string) bool {
	switch t {
	case "external_document_reference":
		for _, edr := range doc.SPDXDocRef.ExternalDocumentReferences {
			if len(data) > 0 && edr.Checksum.Value == data {
				return false
			}
		}

	case "other_license":
		for _, ol := range doc.SPDXDocRef.OtherLicenses {
			if len(data) > 0 && ol.LicenseIdentifier == data {
				return false
			}
		}

	case "relationship":
		for _, r := range doc.SPDXDocRef.Relationships {
			relStr := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s_%s",
				r.RefA.DocumentRefID, r.RefA.ElementRefID, r.RefA.SpecialID,
				r.Relationship,
				r.RefB.DocumentRefID, r.RefB.ElementRefID, r.RefB.SpecialID,
				r.RelationshipComment)
			if len(data) > 0 && relStr == data {
				return false
			}
		}

	default:
		return true
	}
	return true
}
