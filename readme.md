# OCSF-Tool: Open Cybersecurity Schema Framework Command Line Utility

## Introduction
OCSF-Tool is a command-line utility designed for developers working with the [Open Cybersecurity Schema Framework (OCSF)](https://github.com/ocsf/).
It provides a set of utilities to process OCSF schemas, including a generator for creating Proto files.
Utility uses https://schema.ocsf.io/export/schema to download latest OCSF schema.
This README file aims to provide an overview of OCSF-Tool, its features, installation instructions, and basic usage examples.

## Features
### Proto File Generation (beta)
The tool includes a generator that simplifies the creation of Proto files from OCSF schemas, helping you generate code for various programming languages.

## Demo
[![asciicast](https://asciinema.org/a/2A26OaySGAIEoVHypgRR6NjRM.svg)](https://asciinema.org/a/2A26OaySGAIEoVHypgRR6NjRM)


## Download
To acquire OCSF-Tool, you have two options:

### Option 1 - Automatic Download
Execute the following command to automatically download OCSF-Tool acording to OS and Processor Architecture:

```shell
curl -sfL https://raw.githubusercontent.com/valllabh/ocsf-tool/main/download/download.sh | bash
```

Upon successful execution of the command, you will receive output similar to the following:

```bash
📦 OCSF-Tool Downloading

👍 OS and Architecture detected
👍 Detected latest version of OCSF-Tool
👍 Downloaded OCSF-Tool v0.1.0 (latest)
👍 Verified downloaded files
👍 Extracted the Tar in ./ocsf-tool directory
👍 Tar and Checksums removed

🎉 Download Complete!

Go to "/workspaces/ocsf-tool"
And Run "./ocsf-tool"

Usage:
  ocsf-tool [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  config            Set configuration values for extensions and profiles
  generate-proto    Generate a Proto file
  help              Help about any command
  schema-class-list List all classes in the OCSF schema

Flags:
  -h, --help   help for ocsf-tool

Use "ocsf-tool [command] --help" for more information about a command.
```

### Option 2 - Manually
- Go to the [releases section](https://github.com/valllabh/ocsf-tool/releases) of the OCSF-Tool repository.
- Download the latest binary release suitable for your platform (e.g., Windows, Linux, macOS).

## Command Documentation
For detailed information on using OCSF-Tool and its commands, refer to the [Command Documentation](docs/ocsf-tool.md)

## Example Use Case
### Setting default OCSF Extensions to use
The OCSF Schema is customizable through **extensions** that add new attributes, objects, and event classes, enabling vendor-specific customizations and maintaining a concise core schema.

[More Information on OCSF Extensions](https://github.com/ocsf/ocsf-schema/tree/main/extensions)

#### To set OCSF Extensions.
```
ocsf-tool config extensions linux
```
> All extensions will be active if no extensions are configured 

### Setting default OCSF Extensions to use
OCSF (Open Cybersecurity Schema Framework) profiles are predefined sets of data models and attributes within the OCSF Schema that cater to specific cybersecurity use cases or scenarios. These profiles help in standardizing the way cybersecurity data is structured and shared, ensuring compatibility and interoperability across different systems and tools in the cybersecurity landscape.

[More information on OCSF Profiles](https://schema.ocsf.io/1.1.0-dev/profiles?extensions=)

#### To set Profiles
```
ocsf-tool config profiles cloud container
```
> All profiles will be active if no profiles are configured 

### Generate Proto
This command generates proto files for the mentioned classes in the arguments. There are more options avaiable to the command use `--help` for more information or follow [documentation link](docs/ocsf-tool_generate-proto.md)

[List of all possible OCSF classes](https://schema.ocsf.io/1.1.0-dev/classes?extensions=)

#### To generate proto files for OCSF classes
```
ocsf-tool generate-proto file_activity security_finding
```

## Feedback and Contributions
We eagerly welcome your valuable feedback, bug reports, and contributions to the OCSF-Tool project. If you encounter any issues or have suggestions for enhancements, kindly create an issue on GitHub.

## License
OCSF-Tool is distributed under the [Apache 2 License](LICENSE). Your usage and contributions are subject to the terms outlined in this license.
