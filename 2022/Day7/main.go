package main

import (
	"aoc/util"
	"os"
	"strings"
)

const maxSize = 100000

var totalCount = 0

type Dir struct {
	name            string
	size            int
	parent          *Dir
	subDirectoryMap map[string]*Dir
}

func sumDirectorySizes(filename string) (result int, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	rootDir := &Dir{
		name:            "/",
		parent:          nil,
		subDirectoryMap: make(map[string]*Dir),
	}
	buildFileSystem(file, rootDir)

	sumSizes(rootDir)
	traverse(rootDir, rootDir)

	return totalCount, err
}

func buildFileSystem(file []byte, rootDir *Dir) {
	currentDir := rootDir
	for _, line := range strings.Split(string(file), "\n") {
		if line == "$ cd /" {
			// we deal with this above
			continue
		}

		lineFields := strings.Fields(line)
		elem1 := lineFields[0]
		elem2 := lineFields[1]
		if elem1 == "$" {
			currentDir = handleUserInput(elem2, lineFields, currentDir)
		} else {
			handleOutput(elem1, currentDir, elem2)
		}
	}
}

func handleUserInput(elem2 string, lineFields []string, currentDir *Dir) *Dir {
	command := elem2
	switch command {
	case "ls":
		break
	case "cd":
		elem3 := lineFields[2]
		if elem3 == ".." {
			currentDir = currentDir.parent
		} else {
			currentDir = currentDir.subDirectoryMap[elem3]
		}
	}
	return currentDir
}

func handleOutput(elem1 string, currentDir *Dir, elem2 string) {
	if elem1 == "dir" {
		currentDir.subDirectoryMap[elem2] = &Dir{
			name:            elem2,
			parent:          currentDir,
			subDirectoryMap: make(map[string]*Dir),
		}
	} else {
		fileSize := util.MustAtoI(elem1)
		currentDir.size += fileSize
	}
}

func traverse(dir, rootDir *Dir) int {
	for _, v := range dir.subDirectoryMap {
		traverse(v, rootDir)
	}
	empty := diskSpace - rootDir.size
	if empty+dir.size >= unusedSpace && dir.size < minimumSize {
		minimumSize = dir.size
	}
	return dir.size
}

func sumSizes(dir *Dir) int {
	for _, v := range dir.subDirectoryMap {
		dir.size += sumSizes(v)
	}
	if dir.size < maxSize {
		totalCount += dir.size
	}
	return dir.size
}
