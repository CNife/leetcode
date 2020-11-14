package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	packages := collectPackages()
	runGoCommand("test", packages)
	runGoCommand("fmt", packages)
}

func collectPackages() []string {
	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panicf("list current directory: %v", err)
	}
	var result []string
	for _, dir := range dirs {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
			relativeDirName := fmt.Sprintf(".%c%v", os.PathSeparator, dir.Name())
			result = append(result, relativeDirName)
		}
	}
	return result
}

func runGoCommand(commandName string, packages []string) {
	cmdArgs := append([]string{commandName}, packages...)
	cmd := exec.Command("go", cmdArgs...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		log.Panicf("%v packages: %v", commandName, err)
	}
}
