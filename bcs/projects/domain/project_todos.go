package domain

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileTodo struct {
	FilePath    string
	ParentTitle string
	Todo        string
	Done        bool
}

func ProjectTodos(searchString string) ([]FileTodo, error) {
	todos := make([]FileTodo, 0)
	path, err := ProjectPath(searchString)
	if err != nil {
		return todos, err
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		fmt.Println(path)

		if strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			fmt.Println("Markdown")
			thisFileTodos, err := fileTodos(path)
			fmt.Println(thisFileTodos)
			if err != nil {
				return err
			}
			todos = append(todos, thisFileTodos...)
		}
		return nil
	})

	return todos, nil
}

func fileTodos(path string) ([]FileTodo, error) {
	fileTodos := make([]FileTodo, 0)

	file, err := os.Open(path)
	if err != nil {
		return fileTodos, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	todoRegex := regexp.MustCompile(`^- \[ ] .*$`)
	titleRegex := regexp.MustCompile(`^\s*#+ .*$`)

	lastTitle := ""
	for scanner.Scan() {
		line := scanner.Text()
		if titleRegex.Match([]byte(line)) {
			lastTitle = line
		} else if todoRegex.Match([]byte(line)) {
			fileTodos = append(fileTodos, FileTodo{
				FilePath:    path,
				ParentTitle: lastTitle,
				Todo:        line,
				Done:        false,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return fileTodos, err
	}

	return fileTodos, nil
}
