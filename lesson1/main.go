package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	// Задание 1, 2
	a, b := 10, 0
	fmt.Println("a =", a, ", b =", b)
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("a / b =", res)
	}

	// Задание 3
	if _, err := os.Stat("testdir"); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("testdir", 0777)
			if err != nil {
				fmt.Println("Не удалось создать директорию")
				return
			}
		}
	}
	numOfFiles := 1000000
	numOfFilesCreated := 0
	openedFiles := make([]*os.File, 0, numOfFiles)
	for numOfFilesCreated != numOfFiles {
		numOfFilesCreated, err = createFiles(&openedFiles, numOfFilesCreated, numOfFiles)
		if err != nil {
			fmt.Println(err)
		}
		openedFiles = openedFiles[len(openedFiles):]
	}
	fmt.Println("Создано файлов всего:", numOfFilesCreated)

	// Задание 4
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}

func divide(a int, b int) (res int, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("При делении возникла ошибка\n%w, occurred at: %v", v, time.Now())
		}
	}()
	return a / b, err
}

func createFiles(openedFiles *[]*os.File, numFrom int, numTo int) (numOfFilesCreated int, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("При создании файлов возникла ошибка\n%w, occurred at: %v", v, time.Now())
		}
		for _, file := range *openedFiles {
			//fmt.Println("file", file.Name(), "closed")
			file.Close()
		}
	}()
	for i := numFrom + 1; i <= numTo; i++ {
		file, err := os.Create("testdir/file" + fmt.Sprintf("%v", i))
		if err != nil {
			panic(err)
		}
		numOfFilesCreated = i
		*openedFiles = append(*openedFiles, file)
		// т.к. ОС не выдала проблем, то чтобы смоделировать аварийную ситуацию добавил панику при i = 10 и i = 50
		if i == 10 {
			panic(errors.New("10 files created"))
		}
		if i == 50 {
			panic(errors.New("50 files created"))
		}
	}
	return numOfFilesCreated, err
}
