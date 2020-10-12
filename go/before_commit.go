package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var testPackages []string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panicf("list current dir: %v", err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() != ".idea" {
			testPackages = append(testPackages,
				fmt.Sprintf(".\\%v\\", file.Name()))
		}
	}

	for _, tp := range testPackages {
		cmd := exec.Command("go", "test", tp)
		cmd.Stdout = os.Stdout
		if err = cmd.Run(); err != nil {
			log.Panicf("test %v: %v", tp, err)
		}
	}

	for _, tp := range testPackages {
		cmd := exec.Command("go", "fmt", tp)
		if err = cmd.Run(); err != nil {
			log.Panicf("format %v: %v", tp, err)
		}
	}
}
