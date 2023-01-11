package day7

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	size     int
	parent   *Directory
	children map[string]*Directory
}

type FileSystem struct {
	root *Directory
	pwd  *Directory
}

func createFileSystem() *FileSystem {
	root := &Directory{
		parent:   nil,
		children: make(map[string]*Directory),
		size:     0,
	}
	return &FileSystem{
		root: root,
	}
}

func (fs *FileSystem) cd(dir string) {
	if dir == ".." {
		fs.pwd = fs.pwd.parent
	} else if dir == "/" {
		fs.pwd = fs.root
	} else {
		if fs.pwd.children[dir] == nil {
			fs.pwd.children[dir] = &Directory{
				parent:   fs.pwd,
				children: make(map[string]*Directory),
				size:     0,
			}
		}
		fs.pwd = fs.pwd.children[dir]
	}
}

func (fs *FileSystem) registerSize(size int) {
	current := fs.pwd
	for current != nil {
		current.size += size
		current = current.parent
	}
}

func (fs *FileSystem) sizes() []int {
	sizes := make([]int, 0)
	return appendSizes(sizes, fs.root)
}

func appendSizes(sizes []int, dir *Directory) []int {
	sizes = append(sizes, dir.size)
	for _, child := range dir.children {
		sizes = appendSizes(sizes, child)
	}
	return sizes
}

func PartOne(sizes []int) int {
	var sum int
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func PartTwo(sizes []int, totalSize int) int {
	totalUsed := totalSize
	total := 70000000
	needed := 30000000
	sort.Ints(sizes)
	for _, size := range sizes {
		if total-totalUsed+size >= needed {
			return size
		}
	}
	return -1
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	fs := createFileSystem()
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != '$' {
			if line[0] == 'd' {
				continue
			}
			i := strings.IndexByte(line, ' ')
			size, err := strconv.Atoi(line[:i])
			if err != nil {
				panic(err)
			}
			fs.registerSize(size)
			continue
		}
		fields := strings.Fields(line)
		if fields[1] == "cd" {
			fs.cd(fields[2])
		}
	}

	sizes := fs.sizes()
	totalSize := fs.root.size

	if output := PartOne(sizes); output != 1432936 {
		panic(fmt.Errorf("PartOneDaySeven failed -> %d", output))
	}

	if output := PartTwo(sizes, totalSize); output != 272298 {
		panic(fmt.Errorf("PartTwoDaySeven failed -> %d", output))
	}
}
