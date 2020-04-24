package main

import (
	"fmt"
	"path"
)
func main() {
	filename := "CMakeLists.haha.txt"
	filenameall := path.Base(filename)
	filesuffix := path.Ext(filename)
	fileprefix := filenameall[0:len(filenameall) - len(filesuffix)]
	//fileprefix, err := strings.TrimSuffix(filenameall, filesuffix)

	fmt.Println("file name:", filenameall)
	fmt.Println("file prefix:", fileprefix)
	fmt.Println("file suffix:", filesuffix)
}