package parser

import (
	"yabt/generator"
)

type Config interface {
	GetProjectName() (string, error)
	GetVersion() string
	GetCPPStandard() int
	GetOutputDirectory() string
	GetOutputType() string
	GetLibraryType() string
	GetIncludeDirs() []string
	GetLibraries() []string
	GetLibraryDirs() []string
	GetBuildFlags() []string
	GetSources() []string
	GetExternalProjects() []ExternalProject
}

func New(filePath string) (Config, error) {
	config, err := YamlReadFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GenerateFromConfig(config Config, generatorType string, outDir string) error {
	generatorInstance, err := generator.New(generatorType)
	if err != nil {
		return err
	}

	projectName, err := config.GetProjectName()
	if err != nil {
		return err
	}
	generatorInstance.GenerateProjectName(projectName)
	generatorInstance.GenerateVersion(config.GetVersion())
	generatorInstance.GenerateCPPStandard(config.GetCPPStandard())
	generatorInstance.GenerateBuildDir(config.GetOutputDirectory())

	for _, dir := range config.GetIncludeDirs() {
		generatorInstance.AddIncludeDir(dir)
	}

	for _, dir := range config.GetLibraryDirs() {
		generatorInstance.AddLibDir(dir)
	}

	for _, lib := range config.GetLibraries() {
		generatorInstance.AddLibrary(lib)
	}

	for _, flag := range config.GetBuildFlags() {
		generatorInstance.AddFlag(flag)
	}

	for _, source := range config.GetSources() {
		generatorInstance.AddSource(source)
	}

	err = generatorInstance.GenerateFiles(outDir)
	if err != nil {
		return err
	}

	return nil
}
