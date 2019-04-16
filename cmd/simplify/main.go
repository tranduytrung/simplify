package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tranduytrung/simplify"
)

var target int

func init() {
	flag.IntVar(&target, "t", 1024, "the number of faces in the output")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: simplify [-t TARGET] input.stl output.stl")
		os.Exit(-1)
	}
	fmt.Printf("Loading %s\n", args[0])
	mesh, err := simplify.LoadBinarySTL(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input mesh contains %d faces\n", len(mesh.Triangles))
	fmt.Printf("Simplifying to %d faces\n", int(target))
	for len(mesh.Triangles) > target {
		mesh = mesh.Simplify(target)
	}
	fmt.Printf("Output mesh contains %d faces\n", len(mesh.Triangles))
	fmt.Printf("Writing %s\n", args[1])
	mesh.SaveBinarySTL(args[1])
}
