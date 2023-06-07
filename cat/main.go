package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		r := io.TeeReader(os.Stdin, os.Stdout)
		io.ReadAll(r)
		os.Stdin.Close()
		os.Stdout.Close()
	} else if len(os.Args) == 2 {
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			printer("ERROR: open " + os.Args[1] + ": no such file or directory")
			z01.PrintRune('\n')
			os.Exit(1)
		} else {
			str := string(file)
			printer(str)
		}
	}
	if len(os.Args) > 2 {
		args := os.Args[1:]
		for _, v := range args {
			_, err := os.ReadFile(v)
			if err != nil {
				// printer("ERROR: " + v + " : no such file or directory\n")
				printer("ERROR: " + err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printer(getContent(v))
		}
	}
}

func printer(err string) {
	for _, v := range err {
		z01.PrintRune(v)
	}
}

func getContent(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		return "ERROR"
	} else {
		return string(file)
	}
}
