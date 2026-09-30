package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ask(r *bufio.Reader, question, def string) string {
	fmt.Printf("%s (%s): ", question, def)

	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "" {
		return def
	}
	return text
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "App name?", "app")
	module := ask(reader, "Module path?", "github.com/you/"+name)

	project := Project{Name: name, Module: module}

	fmt.Printf("\nGerando %s...\n", project.Name)
	if err := project.Generate(); err != nil {
		fmt.Println("erro:", err)
		os.Exit(1)
	}

	fmt.Printf("\nPronto! Para rodar:\n  cd %s\n  make run\n", project.Name)
}
