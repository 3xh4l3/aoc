package main

import (
	"fmt"
	"strings"

	"github.com/3xh4l3/aoc"
)

type dir struct {
	files_size int // size of all dir files
	total_size int // total size of dir with files and dirs
	dirs       []string
}

func main() {
	p1, p2 := day_7("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_7(input string) (p1, p2 int) {
	var (
		dir_stack  []string
		c_path     string
		files_size int
	)

	all_dirs := make(map[string]dir)
	dirs := make([]string, 0)

	f, scanner := aoc.GetFile(input)
	defer f.Close()

	// Parse files and dirs
	for scanner.Scan() {
		var c_dir string

		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "$ cd"):
			// Get current dir
			fmt.Sscanf(line, "$ cd %s", &c_dir)
			// Pop dir stack
			if c_dir == ".." {
				dir_stack = dir_stack[:len(dir_stack)-1]
				continue
			}
			// Store previous dir
			if entry, ok := all_dirs[c_path]; ok {
				entry.dirs = dirs
				entry.files_size = files_size
				all_dirs[c_path] = entry
			}
			// Reset files size and dir list of prev dir
			files_size = 0
			dirs = make([]string, 0)
			// Push new dir
			dir_stack = append(dir_stack, c_dir)
			c_path = strings.Join(dir_stack, "/")
			all_dirs[c_path] = dir{}
		case strings.HasPrefix(line, "$ ls"):
			// Skip ls
			continue
		default:
			// Get files size and dirs list for current dir
			if strings.HasPrefix(line, "dir") {
				var dir_name string
				fmt.Sscanf(line, "dir %s", &dir_name)
				dirs = append(dirs, fmt.Sprintf("%s/%s", c_path, dir_name))
			} else {
				var (
					file_name string
					file_size int
				)
				fmt.Sscanf(line, "%d %s", &file_size, &file_name)
				files_size += file_size
			}
		}

	}
	// Finally
	if entry, ok := all_dirs[c_path]; ok {
		entry.dirs = dirs
		entry.files_size = files_size
		all_dirs[c_path] = entry
	}
	// End of parse files and dirs

	// Calc sizes
	for k, v := range all_dirs {
		v.total_size = getDirTotalSize(all_dirs, v.dirs) + v.files_size
		all_dirs[k] = v

		// Part one answer
		if v.total_size <= 100000 {
			p1 += v.total_size
		}
	}

	// Part two answer
	needed_space := 30000000 - 70000000 + all_dirs["/"].total_size
	p2 = getSmallestDir(all_dirs, needed_space)

	return
}

func getSmallestDir(a_dirs map[string]dir, ns int) (size int) {
	first := true
	for _, v := range a_dirs {
		if v.total_size-ns > 0 {
			if first {
				size = v.total_size
				first = false
			}
			if v.total_size < size {
				size = v.total_size
			}
		}
	}
	return
}

// Recursively calculate dir size
func getDirTotalSize(a_dirs map[string]dir, dirs []string) (size int) {
	for _, v := range dirs {
		d := a_dirs[v]
		size += d.files_size
		size += getDirTotalSize(a_dirs, d.dirs)
	}
	return
}
