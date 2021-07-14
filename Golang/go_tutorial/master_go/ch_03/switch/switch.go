package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: switch number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("This value is not an integer: ", number)
	} else {

		switch {
		case number < 0:
			fmt.Println("less than zero")
		case number > 0:
			fmt.Println("bigger than zero")
		default:
			fmt.Println("zero")
		}
	}

	aString := args[1]
	switch aString {
	case "5":
		fmt.Println("Five")
	case "0":
		fmt.Println("zero")
	default:
		fmt.Println("Do not care!")
	}

	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	var mail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	switch {
	case negative.MatchString(aString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(aString):
		fmt.Println("floating point")
	case mail.MatchString(aString):
		fmt.Println("email address")
	default:
		fmt.Println("SOmething else!")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface")
	default:
		fmt.Println("Not nil interface!")
	}
}
