# AWS CloudWatch Viewer (aclv)

Aclv is a command-line interface (CLI) application that allows users to view AWS CloudWatch saved queries.
It provides a simple and intuitive way to list CloudWatch saved queries, select a particular query folder, and display the actual query stored in the folder.

## Features

- List CloudWatch saved query definition names.
- Display the actual query stored in the query definition.

## Installation

1. Clone the repository:
   ```git clone https://github.com/rahulmlokurte/aws-cloudwatch-logs-viewer.git```
2. Navigate to the project directory:
   ```cd aws-cloudwatch-logs-viewer```
3. Install dependencies:
   ```go mod tidy```
4. Build the application:
   ```go build -o aclv```

## Usage
To use Aclv CLI, run the binary file aclv from the command line. You can execute various commands to list saved queries.

## Commands
- `aclv saved-queries`:  List CloudWatch saved queries

## Flags
- `--startsWith`, `-s`: Filter queries starting with a specific name.

## Example
List CloudWatch saved queries starting with the name "test":
```aclv saved-queries --startsWith=test```

## Dependencies
- AWS SDK for Go: Go SDK for AWS services.
- Cobra: CLI library for Go.
- Lip Gloss: Styling for your Go text UI.