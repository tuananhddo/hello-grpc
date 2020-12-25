// +build pro

package main

import "fmt"

func init() {
	features = append(features,
		"Pro Feature #1",
		"Pro Feature #2",
	)
	fmt.Print("HASAGI")

}
