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

	cmd := exec.Command("go", append([]string{"build"}, os.Args[1:]...)...)
	env := os.Environ()
	if target != nil {
		env = append(env, fmt.Sprintf("GOOS=%s", *target))
	}
	if arch != nil {
		env = append(env, fmt.Sprintf("GOARCH=%s", *arch))
	}
	cmd.Env = env

	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error executing go build: %s\n", err))
		os.Stderr.Write(output)
		os.Exit(1)
	}

	fmt.Print(strings.TrimSpace(string(output)))
}
