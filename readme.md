# OCSF-Tool: Open Cybersecurity Schema Framework Command Line Utility

## Introduction
OCSF-Tool is a command-line utility designed for developers working with the [Open Cybersecurity Schema Framework (OCSF)](https://github.com/ocsf/).
It provides a set of utilities to process OCSF schemas, including a generator for creating Proto files.
Utility uses https://schema.ocsf.io/export/schema to download latest OCSF schema.
This README file aims to provide an overview of OCSF-Tool, its features, installation instructions, and basic usage examples.


## Features
### Proto File Generation (beta)
The tool includes a generator that simplifies the creation of Proto files from OCSF schemas, helping you generate code for various programming languages.

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
👍 Downloaded OCSF-Tool v0.1.3 (latest)
👍 Verified downloaded files
👍 Extracted the Tar in ./ocsf-tool directory
👍 Tar and Checksums removed

🎉 Download Complete!

Go to "/workspaces/ocsf-tool"
And Run "./ocsf-tool"

Usage:
  ocsf-tool [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate
  help        Help about any command

Flags:
  -h, --help   help for ocsf-tool

Use "ocsf-tool [command] --help" for more information about a command.
```

### Option 2 - Manually
- Go to the [releases section](https://github.com/valllabh/ocsf-tool/releases) of the OCSF-Tool repository.
- Download the latest binary release suitable for your platform (e.g., Windows, Linux, macOS).

## Usage
For detailed information on using OCSF-Tool and its commands, refer to the [Command Documentation](docs/ocsf-tool.md)

## Feedback and Contributions
We eagerly welcome your valuable feedback, bug reports, and contributions to the OCSF-Tool project. If you encounter any issues or have suggestions for enhancements, kindly create an issue on GitHub.

## License
OCSF-Tool is distributed under the [Apache 2 License](LICENSE). Your usage and contributions are subject to the terms outlined in this license.