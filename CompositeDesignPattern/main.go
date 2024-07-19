package main

import (
	"fileSystem/application"
	"fmt"
)

func main() {
	// return filesystem with root directory: "/"
	fileSystem := application.NewLinuxFileSystem()

	// ls inside '/'
	files, _ := fileSystem.Ls()
	fmt.Println(files)

	// create directory 'var'
	fileSystem.Mkdir("var")

	// ls inside 'var'
	files, _ = fileSystem.Ls()
	fmt.Println(files)

	// change directory to `var`
	fileSystem, _ = fileSystem.Cd("var")

	// ls inside 'var'
	files, _ = fileSystem.Ls()
	fmt.Println(files)

	// create file inside 'var'
	fileSystem.Touch("config.ini")

	// ls inside 'var'
	files, _ = fileSystem.Ls()
	fmt.Println(files)

}
