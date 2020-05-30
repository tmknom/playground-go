package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	err = proverbs("sample.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var g Grid
	err = g.Set(8, 0, 5)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Printf("エラーです: %v \n", err)
		default:
			fmt.Printf("An error occured: %v \n", err)
		}
		os.Exit(1)
	}

	err = g.Set2(10, 0, 10)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("エラーです: %v \n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e.Error())
			}
		}
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("Recover!: %v\n", e)
			}
		}()
		panic("Panic!")
	}
}

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g Grid) Set2(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if digit < 1 || digit > 9 {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

var ErrBounds = errors.New("out of bounds")
var ErrDigit = errors.New("invalid digit")

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (g Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{writer: f}
	sw.writeln("Errors are values")
	sw.writeln("Handle them gracefully")
	sw.writeln("Errors are values")
	sw.writeln("Handle them gracefully")
	return sw.err
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.writer, s)
}

type safeWriter struct {
	writer io.Writer
	err    error
}
