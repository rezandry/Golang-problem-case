package main

import (
	"fmt"
	"io/ioutil"
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
				f1, _ := ioutil.ReadFile(sourceDir + sourceList[i])
				f2, _ := ioutil.ReadFile(targetDir + targetList[j])

				if !compare(string(f1), string(f2)) {
					sourceList[i] = sourceList[i] + " MODIFIED"
					targetList[j] = targetList[j] + " MODIFIED"
				} else {
					sourceList[i] = ""
					targetList[j] = ""
				}
			}
		}
	}

	for _, file := range sourceList {
		if file != "" && !strings.Contains(file, "MODIFIED") {
			fmt.Println(sourceDir+file, " NEW")
		} else if strings.Contains(file, "MODIFIED") {
			fmt.Println(sourceDir + file)
		}
	}

	for _, file := range targetList {
		if file != "" && !strings.Contains(file, "MODIFIED") {
			fmt.Println(targetDir+file, " DELETED")
		} else if strings.Contains(file, "MODIFIED") {
			fmt.Println(targetDir + file)
		}
	}
}

func compare(a, b string) bool {
	return reflect.DeepEqual(a, b)
}
