package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ExternalProject struct {
	Name    string `yaml:"name"`
	GitRepo string `yaml:"git_repo"`
	Tag     string `yaml:"tag,omitempty"`
}

type ConfigYaml struct {
	ProjectName      string            `yaml:"project_name"`
	Version          string            `yaml:"version,omitempty"`
	CPPStandard      int               `yaml:"cpp_standard,omitempty"`
	OutputDirectory  string            `yaml:"output_directory"`       // Directory where build files will be generated
	OutputType       string            `yaml:"output_type"`            // Can be 'library' or 'executable'
	LibraryType      string            `yaml:"library_type,omitempty"` // static or shared
	IncludeDirs      []string          `yaml:"include_dirs,omitempty"`
	Libraries        []string          `yaml:"libraries,omitempty"`
	LibDirs          []string          `yaml:"lib_dirs,omitempty"`
	BuildFlags       []string          `yaml:"flags,omitempty"`
	SourcePatterns   []string          `yaml:"source_patterns,omitempty"`
	Sources          []string          `yaml:"source,omitempty"` // Deprecated, use source_patterns instead
	ExternalProjects []ExternalProject `yaml:"external_projects,omitempty"`
}

func YamlReadFromFile(filePath string) (*ConfigYaml, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ConfigYaml
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (y *ConfigYaml) GetProjectName() (string, error) {
	if y.ProjectName == "" {
		return "", fmt.Errorf("project_name must be specified in the configuration file")
	}
	return y.ProjectName, nil
}

func (y *ConfigYaml) GetVersion() string {
	if y.Version == "" {
		return "local" // Default version if not specified
	}
	return y.Version
}

func (y *ConfigYaml) GetCPPStandard() int {
	return y.CPPStandard
}

func (y *ConfigYaml) GetOutputDirectory() string {
	return y.OutputDirectory
}

func (y *ConfigYaml) GetOutputType() string {
	if y.OutputType == "" {
		panic("output_type must be specified in the configuration file")
	}
	return y.OutputType
}

func (y *ConfigYaml) GetLibraryType() string {
	return y.LibraryType
}

func (y *ConfigYaml) GetIncludeDirs() []string {
	return y.IncludeDirs
}

func (y *ConfigYaml) GetLibraries() []string {
	return y.Libraries
}

func (y *ConfigYaml) GetLibraryDirs() []string {
	return y.LibDirs
}

func (y *ConfigYaml) GetBuildFlags() []string {
	return y.BuildFlags
}

func (y *ConfigYaml) GetSources() []string {
	if len(y.Sources) > 0 {
		return y.Sources
	}
	if len(y.SourcePatterns) > 0 {
		var sources []string
		for _, pattern := range y.SourcePatterns {
			files, err := filepath.Glob(pattern)
			if err != nil {
				continue // Ignore errors, just skip this pattern
			}
			sources = append(sources, files...)
		}
		return sources
	}
	return nil
}

func (y *ConfigYaml) GetExternalProjects() []ExternalProject {
	return y.ExternalProjects
}
