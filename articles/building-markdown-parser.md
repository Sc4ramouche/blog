# Building a parser for my blog

When I started this blog, I was quite satisfied with how simple it turned out.

However, the fact that I had to rely on a single third party dependency to 

In the same sentiment as in "We are not building enough software" (TODO: add link!), I thought that building a markdown to HTML parser would be fun idea!

I evaluated that I should be able to build a trivial parser which would cover the needs of my blog in under a thousand lines of code, so it seemed like a feasible project.

I did not want to vibe-code my way through, I wanted to learn more about parsers. I wanted to use LLM assistance though, but I did so very carefully.

First, I asked if it is possible to build markdown to html parser in under a thousand lines of code, and got the affirmative response.

Then I asked LLM to suggest a project structure which would align with idiomatic go approach to structure code. I must say that `/cmd` and `/pkg` was a bit confusing for me before, but in the example of building a parser, I think, it finally clicked with me, how these two folder work together. Small win!

I was suggested to do a `/pkg/parses.go` and `/pkg/renderer.go`. At this point, my understanding was that I will have to traverse the markdown file, i.e. parse, however I was not sure what `renderer.go` would be responsible for. Which lead me to another prompt.

I asked what would be very high level picture of how this parser would work, and LLM outlined it for me as `Markdown -> parser -> AST -> renderer -> HTML`. 

And having these prompts for starters, I felt like I am now confident enought to start with the code!

I started with the parser and pushed another prompt, where I wanted to confirm that my approach for parser makes sense. It was something like "I am going to build a document node. Then I traverse the file line by line. I will check the type of line (i.e. Heading, or Paragraph) and then construct a node and push it to the Document array". 

I asked for hint on how would one organise types in Go, and it made polymorphism click for me!

Another is parser looked simple at first, but then I realised that besides blocks, there are smaller inline markdown componets, like bold text or a link.

I got affirmative response!

Something I did not see through at the start is that I'll have to parse the inline contents, line bild text or links. I checked with LLMs what high-level approaches it suggests for that, as I was not sure what shall a "token" be in this case. I did not want to go for regular expressions, so I tried to implement a state machine, it was fun!

Markdown AST can be represented as a tree, and Text nodes are kind of the final leaves, as they have no descendants. That's an insight which I could only acquire by working through the AST implementation.

I also better understood the go ways of doing some things and it made sense for me the pattern of declaring a struct type and then having a constructor function for it.
 
I also learned a bit more about go memory management. For instance, I could declare something like `Children []InlinNode`, and when trying to access the last child I'd get it from slice. What I did not realise is that Go would create a copy of a child here. I wrote the code as I had children declared as array of pointers `Children []*InlineNode`. Nice learn.

What amazed me was also how I built up from simple parsing logic to handling as complex as nested lists. At first I thought it would be really hard to do and involve a ton of changes to make nested lists work, however I have built a solid foundation and I managed to go from simple lists to nested lists quite quickly.

I also liked the experience of working on something across the span of multiple days. I was learning new Go nuances along the way and was stumbling upon code I wrote earlier and noticed how I can improve it. It feels good.

I used claude 3.5.

It has been great coding experience for me. I believe I managed to get a lot of value out of using LLM, but I also learned a lot by not generating any code by LLM and practically only using the high-level blueprint provided by the LLM. Now I want to implement more projects like that, these programming sessions felt really productive and fun!

