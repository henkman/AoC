package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size uint
}

type Dir struct {
	Name            string
	Files           []File
	Dirs            []*Dir
	Parent          *Dir
	CachedTotalSize uint
}

func (dir *Dir) CacheTotalSize() uint {
	var sum uint
	for _, f := range dir.Files {
		sum += f.Size
	}
	for _, d := range dir.Dirs {
		sum += d.CacheTotalSize()
	}
	dir.CachedTotalSize = sum
	return dir.CachedTotalSize
}

func (dir *Dir) Dir(name string) *Dir {
	for _, d := range dir.Dirs {
		if d.Name == name {
			return d
		}
	}
	nd := &Dir{Name: name, Parent: dir}
	dir.Dirs = append(dir.Dirs, nd)
	return nd
}

func (dir *Dir) File(name string, size uint) File {
	for _, f := range dir.Files {
		if f.Name == name {
			return f
		}
	}
	f := File{name, size}
	dir.Files = append(dir.Files, f)
	return f
}

func list(dir *Dir, out io.Writer, level int) {
	for _, f := range dir.Files {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Println("-", f.Name, f.Size)
	}
	for _, d := range dir.Dirs {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Println("->", d.Name, d.CachedTotalSize)
		list(d, out, level+1)
	}
}

func (dir *Dir) List(out io.Writer) {
	list(dir, out, 0)
}

func sumDirsWithTotalSizeAtMost(cur *Dir, atMost uint) uint {
	var sum uint
	if cur.CachedTotalSize <= atMost {
		sum += cur.CachedTotalSize
	}
	for _, d := range cur.Dirs {
		sum += sumDirsWithTotalSizeAtMost(d, atMost)
	}
	return sum
}

func getAllDirsAboveSize(cur *Dir, size uint) []*Dir {
	dirs := []*Dir{}
	if cur.CachedTotalSize >= size {
		dirs = append(dirs, cur)
	}
	for _, d := range cur.Dirs {
		dirs = append(dirs, getAllDirsAboveSize(d, size)...)
	}
	return dirs
}

type SortDirsTotalSize []*Dir

func (s SortDirsTotalSize) Len() int {
	return len(s)
}

func (s SortDirsTotalSize) Less(i, j int) bool {
	return s[i].CachedTotalSize > s[j].CachedTotalSize
}

func (s SortDirsTotalSize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var reListing = regexp.MustCompile(`(dir|\d+) ([a-z.]+)`)

func main() {
	var root *Dir = &Dir{Name: "/"}
	{
		var (
			listing bool
			cur     *Dir
		)
		cur = root
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				if line[0] == '$' {
					if strings.HasPrefix(line[2:], "cd") {
						listing = false
						name := line[5:]
						if name == "/" {
							cur = root
						} else if name == ".." {
							cur = cur.Parent
						} else {
							cur = cur.Dir(name)
						}
					} else if strings.HasPrefix(line[2:], "ls") {
						listing = true
					}
				} else if listing {
					m := reListing.FindStringSubmatch(line)
					if m[1] == "dir" {
						cur.Dir(m[2])
					} else {
						size, err := strconv.ParseUint(m[1], 10, 64)
						if err != nil {
							panic(err)
						}
						cur.File(m[2], uint(size))
					}
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	root.CacheTotalSize()
	// root.List(os.Stdout)

	fmt.Println("first:", sumDirsWithTotalSizeAtMost(root, 100000))
	{
		const (
			DISK_SPACE        = 70000000
			NEEDED_FOR_UPDATE = 30000000
		)
		unused := DISK_SPACE - root.CachedTotalSize
		toSave := NEEDED_FOR_UPDATE - unused
		candidateDirs := getAllDirsAboveSize(root, toSave)
		sort.Sort(sort.Reverse(SortDirsTotalSize(candidateDirs)))
		fmt.Println("second:", candidateDirs[0].CachedTotalSize)
	}
}
