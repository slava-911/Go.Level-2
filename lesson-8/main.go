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
	// workDir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(workDir)
	// fmt.Println(os.Executable())

	// var inputDir string
	// fmt.Println("Enter path to directory:")
	// fmt.Scanf("%s\n", &inputDir)
	// fmt.Println(inputDir)

	// fmt.Println("Enter path to directory:")
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// if err := input.Err(); err != nil {
	// 	fmt.Println("Input error:", err)
	// }
	// inputDir := input.Text()
	fmt.Println("input directory:", *dir, ", concurrency mode:", *c, ", remove duplicate files:", *r)

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

	// files := make([]*duplicates.FileInfo, 0)
	// err := duplicates.FillAllFileList(&files, *dir)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("all file list:")
	// for _, file := range files {
	// 	fmt.Println(file.Name, file.Size)
	// }

	// file1 0
	// file1 0
	// file1 0
	// file1.txt 0
	// file1.txt 0
	// file1.txt 0
	// file2 0
	// file2 0
	// file2.txt 87
	// file2.txt 87
	// file3 0
	// file3 0
	// file3.txt 87
	// file3.txt 87

}
