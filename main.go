package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Enter one argument only")
		os.Exit(1)
	}

	arg := args[1]
	if len(arg) > 9 {
		fmt.Println("Arg are too long")
		os.Exit(1)
	}

	if arg == "help" {
		PrintHelp()
		return
	}

	if len(arg) == 3 {
		arg += strings.Repeat(arg, 2)
	}
	if len(arg) < 9 {
		arg += strings.Repeat("-", 9-len(arg))
	}

	var (
		owner, group, other int
	)

	for i := 0; i < 3; i++ {
		r := string(arg[i*3])
		w := string(arg[i*3+1])
		x := string(arg[i*3+2])
		for i, permission := range []string{r, w, x} {
			if i == 0 && permission != "r" && permission != "-" {
				fmt.Println("Invalid read permission")
				os.Exit(1)
			}
			if i == 1 && permission != "w" && permission != "-" {
				fmt.Println("Invalid write permission")
				os.Exit(1)
			}
			if i == 2 && permission != "x" && permission != "-" {
				fmt.Println("Invalid execute permission")
				os.Exit(1)
			}
		}
		switch i {
		case 0:
			owner = CalcDigit(r, w, x)
		case 1:
			group = CalcDigit(r, w, x)
		case 2:
			other = CalcDigit(r, w, x)
		}
	}

	chmod := owner*100 + group*10 + other

	file_permissions := "-" + arg
	directory_permissions := "d" + arg

	fmt.Println()
	fmt.Println("File:", file_permissions)
	fmt.Println("Directory:", directory_permissions)
	fmt.Println("chmod", chmod)
	fmt.Println()
}

func CalcDigit(r, w, x string) int {
	digit := 0
	if r == "r" {
		digit += 4
	}
	if w == "w" {
		digit += 2
	}
	if x == "x" {
		digit += 1
	}
	return digit
}

func PrintHelp() {
	fmt.Println(`Chmod Calculator

usage:
  chmod-calculator [permission]

example:
  chmod-calculator --x
  chmod-calculator rwxrw-rw-
  `)
}
