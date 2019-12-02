package main

import (
	"io/ioutil"
	"os"
)

func GetPuzzleInput(path string) []byte {
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/" + path)
	if err != nil {
		panic(err)
	}

	conts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return conts
}
