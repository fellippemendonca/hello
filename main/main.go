package main

import (
	"fmt";
	"github.com/fellippemendonca/hello/lib";
)

func main() {
	if 10 < 100 {
		fmt.Printf("%d - %b - %x \n", 42, 42, 42);
		lib.DistanceMeter(2);
	} else {
		fmt.Printf("hello, world 2\n");
	}
	
};
