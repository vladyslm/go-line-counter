package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var numberOfLines int
var extensions []string

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			return count, nil
		} else if err != nil {
			return count, err
		}
	}
}

func openFileAndCountLines(path string) int {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("This path does not exists: ", path)
		fmt.Println("Unable open the file. ", err)
	}
	nLines, err := lineCounter(f)
	if err != nil {
		fmt.Println(err)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	return nLines
}

func readDirectory(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		if f.IsDir() {
			nPath := filepath.Join(path, f.Name())
			readDirectory(nPath)
		} else {
			s := strings.Split(f.Name(), ".")
			fType := s[len(s)-1]
			for _, t := range extensions {
				if fType == t {
					fPath := filepath.Join(path, f.Name())
					l := openFileAndCountLines(fPath)
					numberOfLines += l
				}
			}
		}
	}

}

func getPath(i string) string {
	absPath, err := filepath.Abs("")
	if err != nil {
		fmt.Println(err)
	}

	if strings.Contains(i, "./") {
		sp := strings.Split(i, ".")[1]
		p := filepath.Join(absPath, sp)
		return p
	}
	if i == "" {
		p := filepath.Join(absPath, "")
		return p
	}
	return i
}
