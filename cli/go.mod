module github.com/ivanayov/sbom-composer-go/cli

go 1.16

require (
	github.com/spdx/tools-golang v0.3.0
	github.com/spf13/cobra v1.5.0
)

replace sbom-composer-go/parser => ../parser
