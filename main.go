package main

import (
	"flag"
	"yabt/parser"
)

func main() {

	filePath := flag.String("y", "yabt.yaml", "Path to the YAML configuration file")
	outDir := flag.String("o", "tmp", "Output directory for generated files")
	generatorType := flag.String("generator", "Makefile", "Type of generator to use (e.g., Makefile, CMake, Ninja)")

	flag.Parse()

	configInstance, err := parser.New(*filePath)
	if err != nil {
		panic(err)
	}

	err = parser.GenerateFromConfig(configInstance, *generatorType, *outDir)
	if err != nil {
		panic(err)
	}
	println("Files generated successfully in", *outDir)

}
