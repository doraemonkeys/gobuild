# gobuild

`gobuild` is a custom command-line utility designed to extend the functionality of the standard `go build` command by allowing users to specify target operating systems and architectures directly through command-line options. This tool simplifies cross-compilation processes for Go applications and integrates seamlessly with existing Go toolchain environments.

## Features

- **Set Target OS**: Specify the target operating system for the build process, facilitating easy cross-compilation.
- **Set Target Architecture**: Define the target architecture, supporting diverse hardware platforms.

## Installation

```shell
go install github.com/doraemonkeys/gobuild
```

## Usage

To use `gobuild`, you can simply replace `go build` with `gobuild` and specify additional flags for the target OS and architecture.

### Command-Line Options

- `-t, -target [string]`: Set the target OS for the build (e.g., `linux`, `windows`, `darwin`).
- `-arch [string]`: Set the target architecture for the build (e.g., `amd64`, `arm64`).
- `-h, -help`: Show the help message and exit.

### Examples

**Build for Linux on amd64:**

```bash
gobuild -t linux -arch amd64
```

**Build for Windows on arm64:**

```bash
gobuild -t windows -arch arm64
```

