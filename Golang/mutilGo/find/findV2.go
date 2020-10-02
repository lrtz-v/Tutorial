package find

import (
	"io/ioutil"
	"log"
	"runtime"
)

var maxWorkerNums = runtime.NumCPU() * 2
var workNum int

func FilesV2(path, query string) []string {
	if len(path) == 0 || len(query) == 0 {
		log.Fatal("path and query string can not be empty.")
		return []string{}
	}
	n := len(path)
	if path[n-1] != '/' {
		path += "/"
	}

	resultCh := make(chan string, 500)
	workerDoneCh := make(chan bool, 10)
	searchRequestCh := make(chan string, 100)
	workNum = 1

	go searchV2(path, query, resultCh, searchRequestCh, workerDoneCh, true)

	return waitDone(query, resultCh, searchRequestCh, workerDoneCh)
}

func waitDone(query string, resultCh, searchRequestCh chan string, workerDoneCh chan bool) []string {
	result := []string{}
	for {
		select {
		case path := <-resultCh:
			result = append(result, path)
		case req := <-searchRequestCh:
			go searchV2(req, query, resultCh, searchRequestCh, workerDoneCh, true)
			workNum++
			// log.Printf("workNum: %d\n", workNum)
		case <-workerDoneCh:
			workNum--
			if workNum == 0 {
				return result
			}
		}
	}
}

func searchV2(path, query string, resultCh, searchRequestCh chan string, workerDoneCh chan bool, worker bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			newPath := path + file.Name() + "/"
			if workNum < maxWorkerNums {
				searchRequestCh <- newPath
			} else {
				searchV2(newPath, query, resultCh, searchRequestCh, workerDoneCh, false)
			}
		}
		if query == file.Name() {
			resultCh <- path + file.Name()
		}
	}

	if worker {
		workerDoneCh <- true
	}

}
