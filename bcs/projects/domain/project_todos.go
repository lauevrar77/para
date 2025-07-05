package domain

import (
	"io/fs"
	"path/filepath"
	"strings"

	"para.evrard.online/bcs/shared"
)

func ProjectTodos(searchString string) ([]shared.MdDocument, error) {
	documents := make([]shared.MdDocument, 0)
	path, err := ProjectPath(searchString)
	if err != nil {
		return documents, err
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			filteredDocument, err := filterDocument(path)
			if err != nil {
				return err
			}
			documents = append(documents, *filteredDocument)
		}
		return nil
	})

	return documents, nil
}

func filterDocument(path string) (*shared.MdDocument, error) {
	doc, err := shared.ParseMd(path)
	if err != nil {
		return nil, err
	}

	parent := doc.Root
	doc.Root = *elementTodos(&parent)

	return doc, nil
}

func elementTodos(parent *shared.MdElement) *shared.MdElement {
	todos := make([]*shared.MdElement, 0)
	for _, child := range parent.Childrens {
		if child.Type == "Title" {
			child = elementTodos(child)
			todos = append(todos, child)
		} else if child.Type == "TodoItem" {
			todos = append(todos, child)
		}
	}
	parent.Childrens = todos

	return parent
}
