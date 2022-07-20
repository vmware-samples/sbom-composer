// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	// "github.com/vmware-samples/sbom-composer-go/parser"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spf13/cobra"
)

var (
	dir    string
	load   bool
	save   string
	search string
)

func init() {
	composeCommand.Flags().StringVar(&dir, "dir", "d", "Folder with micro SBOMs")             //for now
	composeCommand.Flags().BoolVarP(&load, "load", "l", true, "Load all available documents") // to start just take the components inf and put in the composed sbom
	composeCommand.Flags().StringVarP(&save, "save", "s", "./", "Save composed data to")      //TODO: generate default folder
	composeCommand.Flags().StringVar(&search, "search", "r", "pattern to search for")         //separate library - we would need it internally
}

var composeCommand = &cobra.Command{
	Use:     "compose",
	Short:   "Compose micro SBOM documents",
	Long:    "TDB",
	Example: "TDB",
	RunE:    runSBOMCompose,
}

func runSBOMCompose(cmd *cobra.Command, args []string) error {
	fmt.Println("running compose")

	allSBOMs := []string{}
	if len(dir) > 0 {
		allSBOMs = readDir(dir)
	}

	var loaded []*spdx.Document2_2
	if load {
		for _, doc := range allSBOMs {
			fmt.Println(doc)
			// Run parser.Load
		}
	}

	if len(save) > 0 {
		for _, doc := range loaded {
			if doc != nil {
				// Run parser.Save
			}
		}
	}
	return nil
}

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

func main() {
	composeCommand.Execute()
}
