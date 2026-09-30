package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

type Project struct {
	Name   string
	Module string
}

type Entry struct {
	Path     string
	Template string // vazio = pasta
}

func (e Entry) IsDir() bool { return e.Template == "" }

func (e Entry) Kind() string {
	if e.IsDir() {
		return "dir "
	}
	return "file"
}

func (p Project) Entries() []Entry {
	entries := []Entry{
		{Path: "cmd/api/main.go", Template: "main.go.tmpl"},
		{Path: "internal/handler/handler.go", Template: "handler.go.tmpl"},
		{Path: "internal/service/service.go", Template: "service.go.tmpl"},
		{Path: "internal/middleware/logging.go", Template: "middleware.go.tmpl"},
		{Path: "internal/server/server.go", Template: "server.go.tmpl"},
		{Path: "go.mod", Template: "go.mod.tmpl"},
	}

	for i := range entries {
		entries[i].Path = filepath.Join(p.Name, entries[i].Path)
	}
	return entries
}

func (p Project) Render(name string) ([]byte, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/"+name)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p Project) Generate() error {
	if _, err := os.Stat(p.Name); err == nil {
		return fmt.Errorf("a pasta %q já existe", p.Name)
	} else if !os.IsNotExist(err) {
		return err
	}

	for _, e := range p.Entries() {
		if e.IsDir() {
			if err := os.MkdirAll(e.Path, 0o755); err != nil {
				return err
			}
			continue
		}

		content, err := p.Render(e.Template)
		if err != nil {
			return fmt.Errorf("renderizando %s: %w", e.Template, err)
		}

		if err := os.MkdirAll(filepath.Dir(e.Path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(e.Path, content, 0o644); err != nil {
			return err
		}
		fmt.Println("  criado", e.Path)
	}
	return nil
}
