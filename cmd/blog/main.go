package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/Sc4ramouche/blog/pkg/markdown"
)

type Preview struct {
	Title string
	Link  string
}

type Post struct {
	Body  template.HTML
	Title string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	outputDir := "public"

	os.Mkdir(outputDir, os.ModePerm)

	generatePost("articles/hello-world.md", outputDir, "hello-world.html")
	generatePost("articles/we-are-not-writing-enough-software.md", outputDir, "we-are-not-writing-enough-software.html")
}

func generatePost(path string, outputDir string, filename string) {
	document, err := markdown.ParseFile(path)
	if err != nil {
		log.Fatalf("Failed to parse markdown file: %s", path)
	}
	htmlContent := document.Render()
	post := Post{Body: template.HTML(htmlContent), Title: document.Title}
	postFile, _ := os.Create(filepath.Join(outputDir, filename))
	defer postFile.Close()

	templates.ExecuteTemplate(postFile, "post.html", post)
}
