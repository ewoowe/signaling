package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Signallings map[string]string = make(map[string]string)

func Init(dir *string) {
	readDir, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}
	for _, fi := range readDir {
		if !fi.IsDir() {
			bytes, err := ioutil.ReadFile(strings.Join([]string{*dir, fi.Name()}, string(os.PathSeparator)))
			if err == nil {
				fmt.Printf("read %s...\n", fi.Name())
				Signallings[fi.Name()] = string(bytes)
			}
		}
	}
}