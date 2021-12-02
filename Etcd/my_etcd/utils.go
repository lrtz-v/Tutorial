package main

import (
	"os"
	"regexp"
	"strings"
)

func generatePath(path, filename string) string {
	b := strings.Builder{}
	b.Write([]byte(path))
	b.WriteByte('/')
	b.Write([]byte(filename))
	return b.String()
}

func Save(path, filename string, data []byte) error {
	filePath := generatePath(path, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

func DirGenrate(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0766)
		if err != nil {
			return err
		}
	}
	return nil
}

func MatchString(pattern, str string) (bool, error) {
	return regexp.MatchString(pattern, str)
}

func DistinctStringList(stringList []string) []string {
	tmpMap := make(map[string]int, len(stringList))
	for _, v := range stringList {
		tmpMap[v] = 0
	}

	newArr := make([]string, len(tmpMap))
	index := 0
	for k := range tmpMap {
		newArr[index] = k
		index++
	}
	return newArr
}
