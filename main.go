package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var target *string
	var arch *string
	// fmt.Println(os.Args)
	// fmt.Println(os.Environ())

	for i := 1; i < len(os.Args); {
		find := false
		switch os.Args[i] {
		case "-t", "-target":
			value := os.Args[i+1]
			target = &value
			find = true
		case "-arch":
			value := os.Args[i+1]
			arch = &value
			find = true
		case "-h", "-help", "--help":
			fmt.Println("Usage: gobuild [options] [arguments], equivalent to go build [arguments]")
			fmt.Println("Options:")
			fmt.Println("  -t, -target")
			fmt.Println("        Set the target OS for the build")
			fmt.Println("  -arch string")
			fmt.Println("        Set the target architecture for the build")
			fmt.Println("  -h    Show this help message")
			os.Exit(0)
		}
		if find {
			os.Args = append(os.Args[:i], os.Args[i+2:]...)
		} else {
			i++
		}
	}

	if target != nil {
		err := os.Setenv("GOOS", *target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if arch != nil {
		err := os.Setenv("GOARCH", *arch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("go", append([]string{"build"}, os.Args[1:]...)...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error executing go build: %s\n", err))
		os.Stderr.Write(output)
		os.Exit(1)
	}

	fmt.Print(strings.TrimSpace(string(output)))
}
