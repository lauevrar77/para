package shared

import (
	"bufio"
	"os"
	"regexp"
)

var todoRegex = regexp.MustCompile(`^- \[ \] (.*)$`)
var doneRegex = regexp.MustCompile(`^- \[[xX]\] (.*)$`)
var titleRegex = regexp.MustCompile(`^\s*(#+) (.*)$`)

type MdDocument struct {
	Path string
	Root MdElement
}

type MdElement struct {
	Type       string
	Value      string
	LineNumber int
	Metadata   map[string]any
	Parent     *MdElement
	Childrens  []*MdElement
}

func ParseMd(path string) (*MdDocument, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := MdElement{
		Type:       "Root",
		LineNumber: -1,
		Parent:     nil,
	}

	lineNumber := 1
	var parent *MdElement = &root
	for scanner.Scan() {
		line := scanner.Text()
		parent = parseLine(parent, line, lineNumber)
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	doc := &MdDocument{
		Path: path,
		Root: root,
	}

	return doc, nil
}

func parseLine(parent *MdElement, line string, lineNumber int) *MdElement {
	if titleRegex.Match([]byte(line)) {
		parent = parseTitle(parent, line, lineNumber)
	} else if todoRegex.Match([]byte(line)) {
		parent = parseTodo(parent, line, lineNumber)
	} else if doneRegex.Match([]byte(line)) {
		parent = parseDone(parent, line, lineNumber)
	} else {
		parent = parseText(parent, line, lineNumber)
	}
	return parent
}

func parseText(parent *MdElement, line string, lineNumber int) *MdElement {
	parent.Childrens = append(parent.Childrens, &MdElement{
		Type:       "Text",
		Value:      line,
		LineNumber: lineNumber,
	})
	return parent
}

func parseDone(parent *MdElement, line string, lineNumber int) *MdElement {
	matches := doneRegex.FindStringSubmatch(line)
	todo := matches[1]
	parent.Childrens = append(parent.Childrens, &MdElement{
		Type:       "TodoItem",
		Value:      todo,
		LineNumber: lineNumber,
		Metadata: map[string]any{
			"Done": true,
		},
	})
	return parent
}

func parseTodo(parent *MdElement, line string, lineNumber int) *MdElement {
	matches := todoRegex.FindStringSubmatch(line)
	todo := matches[1]
	parent.Childrens = append(parent.Childrens, &MdElement{
		Type:       "TodoItem",
		Value:      todo,
		LineNumber: lineNumber,
		Metadata: map[string]any{
			"Done": false,
		},
	})
	return parent
}

func parseTitle(parent *MdElement, line string, lineNumber int) *MdElement {
	matches := titleRegex.FindStringSubmatch(line)
	titleLevel := len(matches[1])
	title := matches[2]
	titleElement := MdElement{
		Type:       "Title",
		Value:      title,
		LineNumber: lineNumber,
		Metadata: map[string]any{
			"Level": titleLevel,
		},
	}

	titleParent := findParent(parent, titleElement)
	titleElement.Parent = titleParent
	titleParent.Childrens = append(titleParent.Childrens, &titleElement)
	return &titleElement
}

func findParent(parent *MdElement, titleElement MdElement) *MdElement {
	if parent.Type == "Root" {
		return parent
	}

	if parent.Type != "Title" {
		panic("Should never get here")
	}

	titleLevel := titleElement.Metadata["Level"].(int)
	parentTitleLevel := parent.Metadata["Level"].(int)
	if parentTitleLevel < titleLevel {
		return parent
	} else {
		return findParent(parent.Parent, titleElement)
	}
}
