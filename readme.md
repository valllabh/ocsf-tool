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

## Configuration File Documentation

This configuration file is written in YAML format and contains settings for extensions, profiles, and schema.

`config.yaml` will be automatically generated with following default structure.

```
extensions:
    discovery:
        paths:
            - $CWD/example/extensions
    selected: []
profiles:
    selected: []
schema:
    loading:
        strategies:
            repository:
                branch:
                    name: main
                directory:
                    path: $CWD/schema/git
                url: https://github.com/ocsf/ocsf-schema
        strategy: repository
    path: $CWD/schema
```

### Extensions

The `extensions` section is used to configure the extensions for the application.

- `discovery.paths`: This is an array of paths where the application should look for extensions. Refer section "Path variables"
- `selected`: This is an array of selected extensions. If empty, all extensions will be selected.

### Profiles

The `profiles` section is used to configure the profiles for the application.

- `selected`: This is an array of selected profiles. If empty, all profiles will be selected.

### Schema

The `schema` section is used to configure the schema for the application.

- `loading`: This section contains settings for loading the schema.
  - `strategies`: This section contains settings for different loading strategies.
    - `repository`: This section contains settings for loading the schema from a repository.
      - `branch.name`: The name of the branch to load the schema from.
      - `directory.path`: The path where the schema should be saved. Refer section "Path variables"
      - `url`: The URL of the repository to load the schema from.
  - `strategy`: The loading strategy to use. In this case, it's set to `repository`.
- `path`: The path where the schema is located. Refer section "Path variables"

### Path variables
- `$CWD` Replaces with current working directory (Recommended Option)
- `$HOME` Replaces with user home directory
- `$TMP` Replaces with system temporary directory

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
This command generates proto files for the mentioned classes in the arguments. There are more options available to the command use `--help` for more information or follow [documentation link](docs/ocsf-tool_generate-proto.md)

[List of all possible OCSF classes](https://schema.ocsf.io/1.1.0-dev/classes?extensions=)

#### To generate proto files for OCSF classes
```
ocsf-tool generate-proto file_activity security_finding
```

## Feedback and Contributions
We eagerly welcome your valuable feedback, bug reports, and contributions to the OCSF-Tool project. If you encounter any issues or have suggestions for enhancements, kindly create an issue on GitHub.

## License
OCSF-Tool is distributed under the [Apache 2 License](LICENSE). Your usage and contributions are subject to the terms outlined in this license.
