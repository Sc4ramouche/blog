package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	articles := []markdown.Document{}
	err := filepath.Walk("articles", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Could not read /articles directory")
		}

		if !info.IsDir() {
			doc := parsePost(path)
			if doc.Draft != "" {
				return nil
			}

			articlePath := strings.ReplaceAll(strings.TrimPrefix(path, "articles/"), ".md", ".html")
			doc.Path = articlePath
			articles = append(articles, *doc)
			generatePost(doc, "public")
		}
		return nil
	})
	if err != nil {
		log.Fatal("Could not read /articles directory")
	}

	indexFile, _ := os.Create("public/index.html")
	defer indexFile.Close()
	templates.ExecuteTemplate(indexFile, "index.html", articles)
}

func parsePost(path string) *markdown.Document {
	document, err := markdown.ParseFile(path)
	if err != nil {
		log.Fatalf("Failed to parse markdown file: %s", path)
	}
	return document
}

func generatePost(document *markdown.Document, outputDir string) {
	htmlContent := document.Render()
	post := Post{Body: template.HTML(htmlContent), Title: document.Title}
	postFile, _ := os.Create(filepath.Join(outputDir, document.Path))
	defer postFile.Close()

	templates.ExecuteTemplate(postFile, "post.html", post)
}
