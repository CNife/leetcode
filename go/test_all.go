package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var testDirs []string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panicf("list current dir: %v", err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() != ".idea" {
			testDirs = append(testDirs, file.Name())
		}
	}

	for _, testcase := range testDirs {
		cmd := exec.Command("go", "test", ".\\"+testcase+"\\")
		cmd.Stdout = os.Stdout
		if err = cmd.Run(); err != nil {
			log.Panicf("testing %v: %v", testcase, err)
		}
	}
}
