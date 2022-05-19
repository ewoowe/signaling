package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Signaling map[string]string = make(map[string]string)

func init() {
	dir := flag.String("dir", "./", "please input signal dirs")
	flag.Parse()
	readDir, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}
	for _, fi := range readDir {
		if !fi.IsDir() {
			bytes, err := ioutil.ReadFile(strings.Join([]string{*dir, fi.Name()}, string(os.PathSeparator)))
			if err == nil {
				fmt.Printf("read %s...\n", fi.Name())
				Signaling[fi.Name()] = string(bytes)
			}
		}
	}
}
