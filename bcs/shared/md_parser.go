package shared

import (
	"bufio"
	"os"
	"regexp"
)

type MdDocument struct {
	Childrens []MdElement
}

type MdElement struct {
	Type       string
	Value      string
	LineNumber int
	Metadata   map[string]any
}

func ParseMd(path string) (*MdDocument, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	todoRegex := regexp.MustCompile(`^- \[ \] (.*)$`)
	doneRegex := regexp.MustCompile(`^- \[[xX]\] (.*)$`)
	titleRegex := regexp.MustCompile(`^\s*(#+) (.*)$`)

	doc := &MdDocument{}

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if titleRegex.Match([]byte(line)) {
			matches := titleRegex.FindStringSubmatch(line)
			titleLevel := len(matches[1])
			title := matches[2]
			doc.Childrens = append(doc.Childrens, MdElement{
				Type:       "Title",
				Value:      title,
				LineNumber: lineNumber,
				Metadata: map[string]any{
					"Level": titleLevel,
				},
			})
		} else if todoRegex.Match([]byte(line)) {
			matches := todoRegex.FindStringSubmatch(line)
			todo := matches[1]
			doc.Childrens = append(doc.Childrens, MdElement{
				Type:       "TodoItem",
				Value:      todo,
				LineNumber: lineNumber,
				Metadata: map[string]any{
					"Done": false,
				},
			})
		} else if doneRegex.Match([]byte(line)) {
			matches := doneRegex.FindStringSubmatch(line)
			todo := matches[1]
			doc.Childrens = append(doc.Childrens, MdElement{
				Type:       "TodoItem",
				Value:      todo,
				LineNumber: lineNumber,
				Metadata: map[string]any{
					"Done": true,
				},
			})
		}
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return doc, nil
}
