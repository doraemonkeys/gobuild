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
			target = &os.Args[i+1]
			find = true
		case "-arch":
			arch = &os.Args[i+1]
			find = true
		case "-h", "-help":
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
	if target != nil {
		cmd.Env = append(os.Environ(), fmt.Sprintf("GOOS=%s", *target))
	}
	if arch != nil {
		cmd.Env = append(os.Environ(), fmt.Sprintf("GOARCH=%s", *arch))
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error executing go build: %s\n", err))
		os.Stderr.Write(output)
		os.Exit(1)
	}

	fmt.Print(strings.TrimSpace(string(output)))
}
