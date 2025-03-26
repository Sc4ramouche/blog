### Blog

* Write content in `/articles`
* Run `go run .` to update the content in `/public`
* To host content locally run `python3 -m http.server 9000 --directory ./public`

### Parser 

When I started the blog, the only dependecy I had was the markdown to html parser. I thought it'd be fun to roll my own parser and learn a thing or two along the way.

#### Parser TODOs

* Handle images
* Sanitise HTML
* Implement `Stringer` interface for AST nodes, so that they can be neatly printed.
* Maybe ordered lists
