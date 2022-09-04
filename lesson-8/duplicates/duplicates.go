// Package duplicates implements functions to find duplicate files in drectory and its subdirectories.
//
// It is possible to remove duplicate files if the appropriate flag is passed.
//
// There are two modes: Normal and Concurrent.

package duplicates

import (
	"errors"
	"io/fs"
	"os"
	"runtime"
	"sort"
	"sync"

	"golang.org/x/sync/errgroup"
)

type FileInfo struct {
	Path string
	Name string
	Size int64
}

var ErrNotExist = errors.New("directory does not exist")
var gMutex sync.Mutex

// FillAllFileList takes a slice and fills it with all files in a directory and its subdirectories.
func FillAllFileList(fileList *[]*FileInfo, dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	dirList, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, el := range dirList {
		name := el.Name()
		path := dirPath + "/" + el.Name()
		if el.IsDir() {
			FillAllFileList(fileList, path)
		} else {
			*fileList = append(*fileList, &FileInfo{path, name, el.Size()})
		}
	}
	return nil
}

// FillAllFileListConcurrent takes a slice and using concurrent mode fills it with all files
// in a directory and its subdirectories.
func FillAllFileListConcurrent(fileList *[]*FileInfo, dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	dirList, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	var (
		wg   sync.WaitGroup
		pool = make(chan struct{}, runtime.NumCPU())
	)
	for _, el := range dirList {
		pool <- struct{}{}
		wg.Add(1)
		el := el
		go func() {
			defer func() {
				wg.Done()
				<-pool
			}()
			name := el.Name()
			path := dirPath + "/" + el.Name()
			if el.IsDir() {
				FillAllFileListConcurrent(fileList, path)
			} else {
				gMutex.Lock()
				*fileList = append(*fileList, &FileInfo{path, name, el.Size()})
				gMutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return nil
}

// GetAllFileList returns a slice of all files in a directory and its subdirectories.
func GetAllFileList(dirPath string) ([]*FileInfo, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	dirList, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var dirfileList []*FileInfo
	for _, el := range dirList {
		name := el.Name()
		path := dirPath + "/" + el.Name()
		if el.IsDir() {
			subdirfileList, err := GetAllFileList(path)
			if err != nil {
				return nil, err
			}
			dirfileList = append(dirfileList, subdirfileList...)
		} else {
			dirfileList = append(dirfileList, &FileInfo{path, name, el.Size()})
		}
	}
	return dirfileList, nil
}

// GetAllFileListConcurrent using concurrent mode returns a slice of all files
// in a directory and its subdirectories.
func GetAllFileListConcurrent(dirPath string) ([]*FileInfo, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	dirList, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	var (
		dirfileList []*FileInfo
		wg          sync.WaitGroup
		pool        = make(chan struct{}, runtime.NumCPU())
		mutex       sync.Mutex
	)
	for _, el := range dirList {
		pool <- struct{}{}
		wg.Add(1)
		go func(el fs.FileInfo) {
			defer func() {
				mutex.Unlock()
				wg.Done()
				<-pool
			}()
			name := el.Name()
			path := dirPath + "/" + el.Name()
			if el.IsDir() {
				subdirfileList, err := GetAllFileListConcurrent(path)
				if err != nil {
					return
				}
				mutex.Lock()
				dirfileList = append(dirfileList, subdirfileList...)
			} else {
				mutex.Lock()
				dirfileList = append(dirfileList, &FileInfo{path, name, el.Size()})
			}
		}(el)
	}
	wg.Wait()
	return dirfileList, nil
}

// GetDuplicateFileList takes directory path and a flag for using parallel mode,
// and returns a slice of duplicate files and an error if something went wrong.
// First, a slice of all files is formed, then it is sorted, and then duplicates are searched for in it.
func GetDuplicateFileList(inputDir string, cMode bool) ([]*FileInfo, error) {
	if _, err := os.Stat(inputDir); err != nil {
		if os.IsNotExist(err) {
			// return nil, fmt.Errorf("directory %s is not exist: %w", inputDir, err)
			return nil, ErrNotExist
		}
	}

	// fList := make([]*FileInfo, 0)
	// if err := FillAllFileList(&fList, inputDir); err != nil {
	// 	return nil, err
	// }
	var fList []*FileInfo
	var err error
	if cMode {
		fList, err = GetAllFileListConcurrent(inputDir)
	} else {
		fList, err = GetAllFileList(inputDir)
	}
	if err != nil {
		return nil, err
	}

	// fmt.Println("All file list:")
	// for _, file := range filesList {
	// 	fmt.Println(file)
	// }

	sort.Slice(fList, func(i, j int) bool {
		return (fList[i].Name < fList[j].Name) ||
			(fList[i].Name == fList[j].Name && fList[i].Size < fList[j].Size)
	})
	// fmt.Println("sorted all file list:")
	// for _, file := range fList {
	// 	fmt.Println(file.Name, file.Size)
	// }

	fListLen := len(fList)
	duplicateFiles := make([]*FileInfo, 0, fListLen)
	for i := 1; i < fListLen; i++ {
		if fList[i].Name == fList[i-1].Name && fList[i].Size == fList[i-1].Size {
			duplicateFiles = append(duplicateFiles, fList[i-1])
			if i == fListLen-1 {
				duplicateFiles = append(duplicateFiles, fList[i])
			}
		} else if i > 1 && fList[i-1].Name == fList[i-2].Name && fList[i-1].Size == fList[i-2].Size {
			duplicateFiles = append(duplicateFiles, fList[i-1])
		}
	}
	return duplicateFiles, nil
}

// RemoveFiles takes a slice of files and removes each one.
func RemoveFiles(fileList []*FileInfo, cMode bool) error {
	if cMode {
		var (
			// wg   sync.WaitGroup
			eg   errgroup.Group
			pool = make(chan struct{}, runtime.NumCPU())
		)
		for _, file := range fileList {
			pool <- struct{}{}
			// wg.Add(1)
			file := file
			// go func() {
			eg.Go(func() error {
				defer func() {
					// wg.Done()
					<-pool
				}()
				if err := os.Remove(file.Path); err != nil {
					// log.Println(err)
					return err
				}
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
		// wg.Wait()
	} else {
		for _, file := range fileList {
			if err := os.Remove(file.Path); err != nil {
				return err
			}
		}
	}
	return nil
}
