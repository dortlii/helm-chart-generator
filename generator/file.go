package generator

import (
	"log"
	"os"
	"path"

	"github.com/dortlii/helm-chart-generator/helm"
)

var helmFolders = []string{
	"templates",
	"crds",
	"charts",
}

// FileService defines the interface for file services
type FileService interface {
	Save() error
}

// NewHelmFileService is a factory function that returns the interface
func NewHelmFileService(path string, helm helm.Helm) FileService {
	return &HelmFiles{
		Path: path,
		Helm: helm,
	}
}

// HelmFiles is the implementation of the helm file structure
type HelmFiles struct {
	// Path where the folder should be placed
	Path string
	// Helm parts to create
	Helm helm.Helm
}

// Save writes the helm chart to the filesystem
func (hf HelmFiles) Save() error {
	// Create the root folder of the helm chart
	fullRootHelmPath := path.Join(hf.Path, hf.Helm.Chart.Name)
	createFolder(fullRootHelmPath)

	// Create all sub folders
	for _, helmFolder := range helmFolders {
		helmDestFolder := path.Join(fullRootHelmPath, helmFolder)
		createFolder(helmDestFolder)
	}

	// Save all files to the filesystem
	// TODO: Rethink the save cascading ...
	hf.Helm.Save(fullRootHelmPath)

	return nil
}

// createFolder helper function
func createFolder(folderPath string) {
	if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
