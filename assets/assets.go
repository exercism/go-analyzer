package assets

// PatternDirs returns a list of folders with patterns for given exercise slug.
func PatternDirs(exercise string) ([]string, error) {
	dir, err := Patterns.Open(exercise)
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}
