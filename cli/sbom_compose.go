// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vmware-samples/sbom-composer/parser"
)

var (
	dir    string
	save   string
	config string
)

func init() {
	composeCommand.Flags().StringVarP(&dir, "dir", "d", "spdx", "Folder with micro SBOMs")
	composeCommand.Flags().StringVarP(&save, "save", "s", "composed.sdpx", "Save composed data to")
	composeCommand.Flags().StringVarP(&config, "conf", "c", "config.yaml", "Configuration for the composed document")
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

	if len(dir) == 0 {
		fmt.Println("please provide a non-empty directory with documents...")
	}

	if len(save) > 0 {
		parser.GenerateComposedDoc(dir, save, config)
	}
	return nil
}

func main() {
	composeCommand.Execute()
}
