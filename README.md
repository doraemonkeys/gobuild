# gobuild

`gobuild` is a custom command-line utility designed to extend the functionality of the standard `go build` command by allowing users to specify target operating systems and architectures directly through command-line options. This tool simplifies cross-compilation processes for Go applications and integrates seamlessly with existing Go toolchain environments.


## Installation

```shell
go install github.com/doraemonkeys/gobuild@latest
```

## Usage

To use `gobuild`, you can simply replace `go build` with `gobuild` and specify additional flags for the target OS and architecture.

### Command-Line Options

- `-t, -target [string]`: Set the target OS for the build (e.g., `linux`, `windows`, `darwin`).
- `-arch [string]`: Set the target architecture for the build (e.g., `amd64`, `arm64`).
- `-h, -help`: Show the help message and exit.

### Examples

**Build for Linux**

```bash
gobuild -t linux
```

**Build for Windows arm64:**

```bash
gobuild -t windows -arch arm64
```

