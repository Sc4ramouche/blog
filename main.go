package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

type Preview struct {
	Title string
	Link  string
}

type Post struct {
	Body template.HTML
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	files, _ := filepath.Glob("articles/*.md")
	outputDir := "public"

	os.Mkdir(outputDir, os.ModePerm)
	generateHomepage(files, outputDir)

	for _, file := range files {
		generatePost(file, outputDir)
	}
}

func generateHomepage(files []string, outputDir string) {
	previews := []Preview{}
	for _, file := range files {
		title := filepath.Base(file)
		title = strings.TrimSuffix(title, ".md")
		link := title + ".html"
		previews = append(previews, Preview{Title: title, Link: link})
	}

	indexFile, _ := os.Create(filepath.Join(outputDir, "index.html"))
	defer indexFile.Close()

	templates.ExecuteTemplate(indexFile, "index.html", previews)
}

func generatePost(file string, outputDir string) {
	content, _ := os.ReadFile(file)
	htmlContent := blackfriday.Run(content)
	title := filepath.Base(file)
	titleHtml := strings.TrimSuffix(title, ".md") + ".html"

	post := Post{Body: template.HTML(htmlContent)}
	postFile, _ := os.Create(filepath.Join(outputDir, titleHtml))
	defer postFile.Close()

	templates.ExecuteTemplate(postFile, "post.html", post)
}
