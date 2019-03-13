package analyzer

import "github.com/tehsphinx/astrav"

// LoadPackage loads a go package from a folder
func LoadPackage(path string) *astrav.Package {
	folder := astrav.NewFolder(path)
	return folder.Package(folder.Pkg.Name())
}
