package main

import (
	"fmt"
	"mutilGo/find"
	"time"
)

func findFliesV1(path, query string) {
	startAt := time.Now()
	res := find.Files(path, query)
	fmt.Printf("[%d] matchs\n", len(res))
	fmt.Printf("Cost %ds\n", time.Since(startAt).Seconds())
}

func findFliesV2(path, query string) {
	startAt := time.Now()
	res := find.FilesV2(path, query)
	fmt.Printf("[%d] matchs\n", len(res))
	fmt.Printf("Cost %ds\n", time.Since(startAt).Seconds())
}

func main() {

	path := "<path>"
	query := "<file>"

	// findFliesV1(path, query)

	fmt.Println()
	findFliesV2(path, query)

}
