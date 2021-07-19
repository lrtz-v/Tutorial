package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %s\n", file)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %s\n", file)
			return err
		}

		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %s\n", file)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: go run file_read.go <file1> [<file2> ...]\n")
		return
	}
	for _, file := range flag.Args() {
		// err := lineByLine(file)
		// err := wordByWord(file)
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
