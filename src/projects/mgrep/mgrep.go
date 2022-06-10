package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func SearchFile(filename string, substring string, wg *sync.WaitGroup) {
	file, _ := os.Open(filename)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	line := 0
	for fileScanner.Scan() {
		line++
		if strings.Contains(fileScanner.Text(), substring) {
			fmt.Printf("%s [%d] %s\n", filename, line, fileScanner.Text())
		}
	}
	wg.Done()
}

func ListAndSearchDirectory(directory string, substring string, wg *sync.WaitGroup) {
	files, _ := os.ReadDir(directory)
	for _, file := range files {
		wg.Add(1)
		if file.IsDir() {
			ListAndSearchDirectory(directory+"/"+file.Name(), substring, wg)
		}
		filename := file.Name()
		go SearchFile(directory+"/"+filename, substring, wg)
	}
}

func main() {
	var wg sync.WaitGroup
	substring, directory := os.Args[1], os.Args[2]
	ListAndSearchDirectory(directory, substring, &wg)
	wg.Wait()
}
