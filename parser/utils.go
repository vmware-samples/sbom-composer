// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readDir(dir string) []string {
	res := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		res = append(res, f.Name())
	}
	return res
}

func isJSON(file string) bool {
	return strings.HasSuffix(file, ".json")
}

func updateFileExtension(file string, outFormat string) string {
	if outFormat == "json" && !isJSON(file) {
		file = file + ".json"
	}
	return file
}

func contains(l []string, e string) bool {
	for _, p := range l {
		if p == e {
			return true
		}
	}
	return false
}

func containsSubstring(l []string, e string) bool {
	for _, r := range l {
		if strings.Contains(e, r) {
			isPrefixOfAnotherPackage := false
			suff := strings.SplitAfter(e, r)
			for i := 1; i < 11; i++ {
				if len(suff) > 1 {
					if strings.HasPrefix(suff[1], fmt.Sprintf("%d", i)) {
						isPrefixOfAnotherPackage = true
					}
				}
			}
			return !isPrefixOfAnotherPackage
		}
	}
	return false
}
