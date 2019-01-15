package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	var (
		uid     uint32
		gid     uint32
		ino     uint64
		nlink   uint64
		device  uint64
		mtime   int64
		atime   int64
		ctime   int64
		blksize int64
		blocks  int64
		root    string
	)

	fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", "path", "basename", "owner", "group", "mode", "inode", "hardlinks", "size", "blocks", "blksize", "mtime", "ctime", "atime", "device", "is_dir")
	if len(os.Args) < 2 {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		root = dir
	} else {
		root = os.Args[1]
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// return if file no longer exists or go can't talk to it
		if info == nil {
			return nil
		}

		// linux specific metrics
		if info.Sys().(*syscall.Stat_t) != nil {
			sysinfo := info.Sys().(*syscall.Stat_t)
			uid = sysinfo.Uid
			gid = sysinfo.Gid
			ino = sysinfo.Ino
			nlink = sysinfo.Nlink
			mtime = sysinfo.Mtim.Sec
			atime = sysinfo.Atim.Sec
			ctime = sysinfo.Ctim.Sec
			blocks = sysinfo.Blocks
			blksize = sysinfo.Blksize
			device = sysinfo.Dev
		} else {
			mtime = info.ModTime().Unix()
		}

		fmt.Printf("\"%s\",\"%s\",%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%v\n", path, info.Name(), uid, gid, info.Mode(), ino, nlink, info.Size(), blocks, blksize, mtime, ctime, atime, device, info.IsDir())
		return nil
	})
	if err != nil {
		panic(err)
	}
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
}
