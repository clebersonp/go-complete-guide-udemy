// the main package is special because it tells to go where is the entry point of application to build it
package main

// packages groups go files
// module groups packages

// command to build: go build
// to build a go program we need to create a module
// to create a module: go mod init $name-project. i.e: go mod init example.com/first-app
// to run the exec file on linux and mac: ./first-app

import "fmt"

// To run in terminal: go run app.go

func main() {
	fmt.Println("Hello World!")
}
