package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	lines int
)

func init() {
	flag.IntVar(&lines, "n", 10, "The location is number lines.")
	flag.Parse()
}

func tail(filename string, n int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ファイル %s を開けませんでした。", filename)
	}
	defer file.Close()

	var stack []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(stack) == n {
			stack = stack[1:]
		}
		stack = append(stack, scanner.Text())
	}
	for _, line := range stack {
		fmt.Println(line)
	}
	return nil
}

func main() {
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "ファイル名を指定してください。")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	err := tail(filename, lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
