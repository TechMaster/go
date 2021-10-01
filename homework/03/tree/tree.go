package tree

import (
	"io/fs"
	"io/ioutil"
	"strings"
)

var sb []string

func PrintDirectoryTree(path string) (string) {
	indent := 0
	printDirectoryTree(path, indent)
	return strings.Join(sb, "")
}

func printDirectoryTree(path string, indent int) error {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	sb = append(sb, getIndentString(indent))
	sb = append(sb, "|--")
	sb = append(sb, path + "/")
	sb = append(sb, "\n")

	for _, entry := range entries {
		if entry.IsDir() {
			printDirectoryTree(entry.Name(), indent + 1)
		} else {
			printFile(entry, indent + 1)
		}
	}

	return nil
}

func printFile(file fs.FileInfo, indent int) {
	sb = append(sb, getIndentString(indent))
	sb = append(sb, "|--")
	sb = append(sb, file.Name())
	sb = append(sb, "\n")
}

func getIndentString(indent int) string {
	sbb := []string{}

	for i := 0; i < indent; i++ {
        sbb = append(sbb, "|  ")
    }
    return strings.Join(sbb,"")
}