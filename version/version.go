package main

import "fmt"

// Version - version of app.
const Version = "v0.0.0" // TODO: set version of you app

func main() {
	fmt.Println(Version) //nolint:forbidigo // it is bin for print version
}
