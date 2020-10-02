package find

import (
	"io/ioutil"
	"log"
)

// Files search files in path
func Files(path, query string) []string {
	if len(path) == 0 || query == "" {
		log.Fatal("path and query string can not be empty.")
		return []string{}
	}
	n := len(path)
	if path[n-1] != '/' {
		path += "/"
	}
	return search(path, query)
}

func search(path, query string) []string {
	result := []string{}

	f, err := ioutil.ReadDir(path)
	if err != nil {
		return result
	}
	for _, file := range f {
		name := file.Name()

		if file.IsDir() {
			r := search(path+file.Name()+"/", query)
			result = merge(result, r)
		}

		if name == query {
			result = append(result, path+name)
		}
	}
	return result
}

func merge(s1, s2 []string) []string {
	newSlice := []string{}
	for _, i := range s1 {
		newSlice = append(newSlice, i)
	}
	for _, i := range s2 {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
