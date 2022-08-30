// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package parser

import (
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
