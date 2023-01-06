package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
)

func ReadFileLines(file string, call func(line string, err error)) {
	f, err := os.Open(file)
	if err != nil {
		call("", err)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			call("", err)
			return
		}
	}(f)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		call(text, nil)
	}

	if err := scanner.Err(); err != nil {
		call("", err)
	}
}

func CreateFile(filePath string) (*os.File, error) {
	if len(filePath) == 0 {
		return nil, errors.New("没有路径")
	}
	dir, _ := path.Split(filePath)
	if len(dir) != 0 {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, errors.New("文件创建失败" + err.Error())
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("文件创建失败" + err.Error())
	}

	return file, nil
}

func WriteFileLines(filePath string, lines []string) error {

	file, err := CreateFile(filePath)

	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := write.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	err = write.Flush()

	if err != nil {
		return err
	}

	return nil
}
