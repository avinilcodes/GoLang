package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Open a file
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	//closing file before close of main()
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	//read the file

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	//shorter way to read a file
	bs, err = ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str = string(bs)
	fmt.Println(str)
	//Creating a file
	file1, err := os.Create("Go.txt")
	if err != nil {
		return
	}
	defer file1.Close()

	file1.WriteString("Writing content for testing purpose")

	//Opening a directory

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) //by passing -1 we instruct it return all the entries
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
