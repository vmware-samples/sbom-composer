# sbom-composer

## Overview
sbom-composer is a tool that serves for composing two or more micro SBOMs into a single SBOM document in SPDX format.

## Try it out

### Build & Run

1. `cd cli/`
2. `go build`
3. `./compose -d <path-to-dir-with-spdx-files-to-compose> [flags]`


* `flags`:
    - `-d`, `--dir`: Folder with micro SBOMs in SPDX format
    - `-s`, `--save`: Saves composed SBOM to a given file. `composed.spdx` by default
    - `-c`, `--conf`: Configuration for the composed document. `sbom-composer/config/example_config.yaml` by default

### Testing changes

Run your local changes with
```
cd cli/
go run sbom_compose.go -d <path-to-dir-with-spdx-files-to-compose> [flags]
```
  
## Documentation

To be added.

## Contributing

The sbom-composer project team welcomes contributions from the community. Before you start working with sbom-composer, please
read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be
signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on
as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).


