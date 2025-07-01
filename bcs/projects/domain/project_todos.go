package domain

import (
	"io/fs"
	"path/filepath"
	"strings"

	"para.evrard.online/bcs/shared"
)

type MdTitle struct {
	Title      string
	Level      int
	LineNumber int
}

type MdTodo struct {
	FilePath    string
	ParentTitle *MdTitle
	Todo        string
	Done        bool
	LineNumber  int
}

func ProjectTodos(searchString string) ([]MdTodo, error) {
	todos := make([]MdTodo, 0)
	path, err := ProjectPath(searchString)
	if err != nil {
		return todos, err
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			thisFileTodos, err := fileTodos(path)
			if err != nil {
				return err
			}
			todos = append(todos, thisFileTodos...)
		}
		return nil
	})

	return todos, nil
}

func fileTodos(path string) ([]MdTodo, error) {
	fileTodos := make([]MdTodo, 0)

	doc, err := shared.ParseMd(path)
	if err != nil {
		return fileTodos, err
	}

	var lastTitle *MdTitle = nil
	for _, element := range doc.Childrens {
		if element.Type == "Title" {
			lastTitle = &MdTitle{
				Title:      element.Value,
				Level:      element.Metadata["Level"].(int),
				LineNumber: element.LineNumber,
			}
		} else if element.Type == "TodoItem" {
			todo := MdTodo{
				FilePath:    path,
				ParentTitle: lastTitle,
				Todo:        element.Value,
				Done:        element.Metadata["Done"].(bool),
				LineNumber:  element.LineNumber,
			}
			fileTodos = append(fileTodos, todo)
		}

	}

	return fileTodos, nil
}
