package domain

import "errors"

func findProject(searchString string) (string, error) {
	paths, err := SearchProject(searchString)
	if err != nil {
		return "", err
	}
	if len(paths) > 1 {
		return "", errors.New("Multiple projects match")
	}
	if len(paths) == 0 {
		return "", errors.New("No project match")
	}
	return paths[0], nil
}
