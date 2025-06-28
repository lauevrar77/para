package domain

import "errors"

func ProjectPath(searchString string) (string, error) {
	paths, err := SearchProject(searchString)
	if err != nil {
		return "", err
	}

	if len(paths) == 1 {
		return paths[0], nil
	} else if len(paths) == 0 {
		return "", errors.New("No project match")
	} else {
		return "", errors.New("Multiple project match")
	}
}
