package analyzer

import (
	"github.com/tehsphinx/astrav"
)

// LoadPackage loads a go package from a folder
func LoadPackage(path string) (*astrav.Package, error) {
	folder := astrav.NewFolder(path)
	_, err := folder.ParseFolder()
	if err != nil {
		return nil, err
	}
	return folder.Package(folder.Pkg.Name()), nil
}
