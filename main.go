package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var upperNext = regexp.MustCompile(`[^a-zA-Z0-9_]`)

type FileInfo struct {
	Path string
	Name string
}

func getDataDir() string {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Missing data dir\n")
		os.Exit(1)
	}
	return flag.Arg(0)
}

func getFilesInDir(dir string) []FileInfo {
	var files []FileInfo
	filesInDir, _ := ioutil.ReadDir(dir)

	for _, file := range filesInDir {
		var fileInfo FileInfo
		fileInfo.Path, _ = filepath.Abs(filepath.Join(dir, file.Name()))
		fileInfo.Name = getName(file.Name())
		files = append(files, fileInfo)
	}
	return files
}

func getName(name string) string {
	var inBytes, outBytes []byte
	var toUpper bool = true

	inBytes = []byte(strings.ToLower(name))

	for i := 0; i < len(inBytes); i++ {
		if upperNext.Match([]byte{inBytes[i]}) {
			toUpper = true
		} else if toUpper {
			outBytes = append(outBytes, []byte(strings.ToUpper(string(inBytes[i])))...)
			toUpper = false
		} else {
			outBytes = append(outBytes, inBytes[i])
		}
	}

	return string(outBytes)
}

func main() {
	files := getFilesInDir(getDataDir())
	if err := writeCode(files); err != nil {
		fmt.Println("bindata: %sn", err.Error())
		os.Exit(1)
	}
}
