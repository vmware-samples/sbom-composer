// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"fmt"

	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
	"github.com/vmware-samples/sbom-composer/parser"
)

const figletStr = `
     _
 ___| |__   ___  _ __ ___         ___ ___  _ __ ___  _ __   ___  ___  ___ _ __
/ __| '_ \ / _ \| '_ '_  \ _____ / __/ _ \| '_ ' _ \| '_ \ / _ \/ __|/ _ \ '__|
\__ \ |_) | (_) | | | | | |_____| (_| (_) | | | | | | |_) | (_) \__ \  __/ |
|___/_.__/ \___/|_| |_| |_|      \___\___/|_| |_| |_| .__/ \___/|___/\___|_|
                                                    |_|

`

var (
	dir     string
	save    string
	config  string
	out     string
	filters []string
)

func init() {
	composeCommand.Flags().StringVarP(&dir, "dir", "d", "spdx", "Folder with micro SBOMs")
	composeCommand.Flags().StringVarP(&save, "save", "s", "composed.spdx", "Save composed data to")
	composeCommand.Flags().StringVarP(&config, "conf", "c", "config.yaml", "Configuration for the composed document")
	composeCommand.Flags().StringVarP(&out, "out", "o", "tv", "Output format of the composed document")
	var defaultFilterList []string
	composeCommand.Flags().StringArrayVarP(&filters, "filters", "f", defaultFilterList, "A list of packages to filter from the output")
}

var composeCommand = &cobra.Command{
	Use:     "sbomcompose",
	Short:   "Compose micro SBOM documents",
	Long:    "TDB",
	Example: "TDB",
	RunE:    runSBOMCompose,
}

func printFiglet() {
	fmt.Printf(aec.YellowF.Apply(figletStr))
}

func runSBOMCompose(cmd *cobra.Command, args []string) error {
	printFiglet()

	if len(dir) == 0 {
		fmt.Println("please provide a non-empty directory with documents...")
	}

	if len(save) > 0 {
		fmt.Printf("...generating composed SPDX document from directory: %s\n", dir)
		fmt.Printf("...using config: %s\n", config)
		err := parser.GenerateComposedDoc(dir, save, out, filters, config)
		if err != nil {
			fmt.Printf("failed to generate composed document: %v/n", err)
		}
		fmt.Printf("...document saved to %s in \"%s\" format\n\n", save, out)
	}
	return nil
}

func main() {
	composeCommand.Execute()
}
