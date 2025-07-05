package shared

import "fmt"

func PrintMdDocument(doc MdDocument) {
	printPath(doc)
	printElement(doc.Root)
	fmt.Print("\n\n")
}

func printPath(doc MdDocument) {
	for i := 0; i < 10; i++ {
		fmt.Print("-")
	}
	fmt.Print(doc.Path)
	for i := 0; i < 10; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func printElement(element MdElement) {
	switch element.Type {
	case "Root":
		printRoot(element)
	case "Title":
		printTitle(element)
	case "TodoItem":
		printTodo(element)
	case "Text":
		printText(element)
	}
	for _, child := range element.Childrens {
		printElement(*child)
	}
}

func printRoot(element MdElement) {}

func printTitle(element MdElement) {
	titleLevel := element.Metadata["Level"].(int)
	fmt.Print("\n")
	for i := 0; i < titleLevel; i++ {
		fmt.Print("#")
	}
	fmt.Printf(" %s\n", element.Value)
}

func printTodo(element MdElement) {
	done := element.Metadata["Done"].(bool)
	doneChar := " "
	if done {
		doneChar = "X"
	}
	fmt.Printf("- [%s] %s\n", doneChar, element.Value)
}

func printText(element MdElement) {
	fmt.Println(element.Value)
}
