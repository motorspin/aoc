package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name, size}
}

func (f *File) Size() int {
	return f.size
}

type Directory struct {
	name        string
	parent      *Directory
	directories map[string]*Directory
	files       map[string]*File
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name, parent, make(map[string]*Directory, 0), make(map[string]*File, 0)}
}

func (d *Directory) Size() int {
	size := 0

	for _, sd := range d.directories {
		size += sd.Size()
	}

	for _, f := range d.files {
		size += f.Size()
	}

	return size
}

func (d *Directory) AddFile(file *File) {
	d.files[file.name] = file
}

func (d *Directory) AddDirectory(directory *Directory) {
	d.directories[directory.name] = directory
}

func (d *Directory) DoesDirectoryExist(name string) bool {
	_, found := d.directories[name]
	return found
}

func (d *Directory) DoesFileExist(name string) bool {
	_, found := d.files[name]
	return found
}

func (d *Directory) List() {
	d.internalList(0)
}

func (d *Directory) internalList(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}

	fmt.Printf("- %s (dir, total size=%d)\n", d.name, d.Size())

	for _, sd := range d.directories {
		sd.internalList(level + 1)
	}

	for _, f := range d.files {
		for i := 0; i < level+1; i++ {
			fmt.Printf("  ")
		}
		fmt.Printf("- %s (file, size=%d)\n", f.name, f.size)
	}
}

type Queue struct {
	items []*Directory
}

func NewQueue() *Queue {
	return &Queue{make([]*Directory, 0)}
}

func (q *Queue) Push(directory *Directory) {
	q.items = append(q.items, directory)
}

func (q *Queue) Pop() (directory *Directory) {
	directory = q.items[0]
	q.items = q.items[1:]
	return
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) List() {
	for _, d := range q.items {
		fmt.Printf("%s\t", d.name)
	}
	fmt.Println("")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	root := NewDirectory("/", nil)
	cd := root

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "$" && parts[1] == "ls" {
			continue
		} else if parts[0] == "dir" {
			cd.AddDirectory(NewDirectory(parts[1], cd))
		} else if parts[0] == "$" && parts[1] == "cd" {
			if parts[2] == "/" {
				cd = root
			} else if parts[2] == ".." {
				cd = cd.parent
			} else {
				cd = cd.directories[parts[2]]
			}
		} else {
			// We have a file
			val, _ := strconv.ParseInt(parts[0], 10, 64)
			cd.AddFile(NewFile(parts[1], int(val)))
		}
	}

	//root.List()

	queue := NewQueue()
	queue.Push(root)

	sum := 0

	for !queue.IsEmpty() {
		dir := queue.Pop()

		for _, sd := range dir.directories {
			queue.Push(sd)
		}

		if dir.Size() <= 100000 {
			sum += dir.Size()
		}
	}

	//fmt.Println(root.Size())
	fmt.Printf("Part 1: %d\n", sum)

	totDiskSpace := 70000000
	free := totDiskSpace - root.Size()
	amountNeededToFree := 30000000 - free

	queue.Push(root)
	smallestDir := root
	smallestDirSize := root.Size()

	for !queue.IsEmpty() {
		dir := queue.Pop()

		for _, sd := range dir.directories {
			queue.Push(sd)
		}

		if dir.Size() > amountNeededToFree && dir.Size() < smallestDirSize {
			smallestDir = dir
			smallestDirSize = dir.Size()
		}
	}

	fmt.Printf("Part 2: %d\n", smallestDir.Size())
}
