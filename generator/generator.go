package generator

import "fmt"

type Generator interface {
	Name() string
	GenerateProjectName(string)
	GenerateVersion(string)
	GenerateCPPStandard(int)
	GenerateBuildDir(string)
	AddIncludeDir(string)
	AddFlag(string)
	AddLibrary(string)
	AddLibDir(string)
	AddSource(string)
	GenerateFiles(string) error
}

func New(outputType string) (Generator, error) {
	switch outputType {
	case "Makefile":
		return makefile(), nil
	// case "CMake":
	// 	return cmake(), nil
	// case "Ninja":
	// 	return ninja(), nil
	default:
		return nil, fmt.Errorf("unsupported output type: %s", outputType)
	}
}
