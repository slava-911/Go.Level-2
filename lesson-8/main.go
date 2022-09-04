package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/slava-911/Go.Level-2/lesson-8/duplicates"
)

var (
	r   *bool
	c   *bool
	dir *string
)

func init() {
	r = flag.Bool("r", false, "remove duplicate files (default false)")
	c = flag.Bool("c", false, "use concurrency mode (default false)")
	dir = flag.String("dir", "testdir", "input directory")
	flag.Parse()
}

func main() {
	fmt.Println("input directory:", *dir, ", concurrent mode:", *c, ", remove duplicate files:", *r)
	duplicateFiles, err := duplicates.GetDuplicateFileList(*dir, *c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("duplicate files:")
	for _, file := range duplicateFiles {
		fmt.Println(file.Name, file.Size)
	}
	if *r {
		if err := duplicates.RemoveFiles(duplicateFiles, *c); err != nil {
			log.Fatal(err)
		}
	}
}
