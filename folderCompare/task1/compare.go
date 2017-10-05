package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {
	sourceDir := os.Args[1]
	targetDir := os.Args[2]

	sourceList := []string{}
	targetList := []string{}

	filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		path = strings.Replace(path, sourceDir, "", -1)
		sourceList = append(sourceList, path)
		return nil
	})

	filepath.Walk(targetDir, func(path string, f os.FileInfo, err error) error {
		path = strings.Replace(path, targetDir, "", -1)
		targetList = append(targetList, path)
		return nil
	})
	x := len(sourceList) - 1
	y := len(targetList) - 1

	for i := x; i >= 0; i-- {
		for j := y; j >= 0; j-- {
			if compare(sourceList[i], targetList[j]) {
				sourceList[i] = ""
				targetList[j] = ""
			}
		}
	}

	for _, file := range sourceList {
		if file != "" {
			fmt.Println(file, " NEW")
		}
	}

	for _, file := range targetList {
		if file != "" {
			fmt.Println(file, " DELETED")
		}
	}
}

func compare(a, b string) bool {
	return reflect.DeepEqual(a, b)
}
