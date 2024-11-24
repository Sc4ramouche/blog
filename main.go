package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

type Preview struct {
	Title string
}

type Post struct {
	Body template.HTML
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/blog/", blogHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	files, err := filepath.Glob("articles/*.md")
	if err != nil {
		log.Fatal("Failed to read articles.")
	}

	posts := []Preview{}
	for _, file := range files {
		title := filepath.Base(file)
		posts = append(posts, Preview{Title: title[:len(title)-3]})
	}
	templates.ExecuteTemplate(w, "index.html", posts)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/blog/"):]
	content, err := os.ReadFile("articles/" + slug + ".md")
	fmt.Println("DEB: ", slug, err)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	htmlContent := blackfriday.Run(content)
	post := Post{Body: template.HTML(htmlContent)}

	templates.ExecuteTemplate(w, "post.html", post)
}
