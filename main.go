package main

import "fmt"

func main() {
	fmt.Println("expression calculator: ")

	var expr string
	for {
		_, err := fmt.Scanln(&expr)
		if err != nil {
			fmt.Printf("failed to read expression: %s\n", err.Error())
			continue
		}

		// exit
		if expr == "exit" {
			return
		}

		parser := NewParser(expr)
		ast := parser.Next()
		fmt.Printf("expr: %s, result: %d\n", expr, Visit(ast))
	}
}
