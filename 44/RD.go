package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}

	var total int64
	for _, entry := range contents {
		if entry.IsDir() {
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			total += info.Size()
		}
	}
	return total, nil
}

func main() {

	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Need a <directory")
		return
	}

	root, err := filepath.EvalSymlinks(arg[1])
	fileinfo, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
	}
	fileinfo, _ = os.Lstat(root)
	mode := fileinfo.Mode()
	if !mode.IsDir() {
		fmt.Println(root, "is not a directory")
		return
	}
	i, err := GetSize(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Total size:", i)
}
