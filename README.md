# sbom-composer

**This project is now officially migrated to [opensbom-generator/sbom-composer](https://github.com/opensbom-generator/sbom-composer).**

**Please visit the new official repository instead.**

## Overview
sbom-composer is a tool that serves for composing two or more micro SBOMs into a single SBOM document in SPDX format.

## Try it out

### Build & Run

1. `cd cli/`
2. `go build`
3. `./sbomcompose -d <path-to-dir-with-spdx-files-to-compose> [flags]`


* `flags`:
    - `-d`, `--dir`: Folder with micro SBOMs in SPDX format
    - `-s`, `--save`: Saves composed SBOM to a given file. `composed.spdx` by default
    - `-c`, `--conf`: Configuration for the composed document. `sbom-composer/config/example_config.yaml` by default
    - `-o`, `--out`: Output format of the composed document: `tv` or `json`. `tv` by default
    - `-f`, `--filters`: A list of packages to filter from the output

To filter a single, or a list of packages, use `-f <pkg1> -f <pkg2> [...]`.

### Testing changes

Run your local changes with:
```
cd cli/
go run sbom_compose.go -d <path-to-dir-with-spdx-files-to-compose> [flags]
```

If testing local changes to some of the sbom-composer's packages, e.g. the `parser`, modify `cli/sbom_compose.go` imports:
```
// "github.com/vmware-samples/sbom-composer/parser"
"sbom-composer/parser"
```
and `cli/go.mod` with:
```
replace sbom-composer/parser => ../parser
```
Then run:
```
cd cli
go mod tidy
```
## Documentation

To be added.

## Contributing

The sbom-composer project team welcomes contributions from the community. Before you start working with sbom-composer, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).


