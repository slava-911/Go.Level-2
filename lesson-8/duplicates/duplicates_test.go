package duplicates

import (
	"errors"
	"fmt"
	"testing"
)

var (
	fileList  = make([]*FileInfo, 0)
	inputDir  = "bigdir"
	testCases = []struct {
		InputDir      string
		NoDuplicates  bool
		ExpectedError error
	}{
		{
			InputDir: "bigdir",
		},
		{
			InputDir: "smalldir",
		},
		{
			InputDir:      "testdir",
			ExpectedError: ErrNotExist,
		},
	}
)

func TestGetDuplicateFileList(t *testing.T) {
	for _, cs := range testCases {
		cs := cs
		t.Run(cs.InputDir, func(t *testing.T) {
			duplicateFiles, err := GetDuplicateFileList(cs.InputDir, false)
			if err != nil {
				if !errors.Is(err, cs.ExpectedError) {
					fmt.Println(cs.ExpectedError)
					t.Fatalf("unexpected error: %s", err.Error())
				}
				return
			}
			if cs.NoDuplicates && len(duplicateFiles) != 0 {
				t.Fatalf("wrong result, number of duplicates: %v, expected: 0", len(duplicateFiles))
			}
			if !checkDuplicateFileList(duplicateFiles) {
				t.Fatalf("wrong result, not all items are duplicates: %v", duplicateFiles)
			}
		})
	}
}

func checkDuplicateFileList(files []*FileInfo) bool {
	filesLen := len(files)
	for i := 1; i < filesLen; i++ {
		if i == 1 || i == filesLen-1 {
			if files[i].Name != files[i-1].Name || files[i].Size != files[i-1].Size {
				return false
			}
		} else if (files[i].Name != files[i-1].Name || files[i].Size != files[i-1].Size) &&
			(files[i-1].Name != files[i-2].Name || files[i-1].Size != files[i-2].Size) {
			return false
		}
	}
	return true
}

func BenchmarkFillAllFileList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FillAllFileList(&fileList, inputDir)
	}
}

func BenchmarkFillAllFileListConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FillAllFileListConcurrent(&fileList, inputDir)
	}
}

func BenchmarkGetAllFileList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAllFileList(inputDir)
	}
}

func BenchmarkGetAllFileListConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAllFileListConcurrent(inputDir)
	}
}

func BenchmarkGetDuplicateFileList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetDuplicateFileList(inputDir, false)
	}
}

func BenchmarkGetDuplicateFileListConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetDuplicateFileList(inputDir, true)
	}
}
