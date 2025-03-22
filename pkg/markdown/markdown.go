package markdown

import (
	"fmt"
	"os"
)

// Note: I learned that it makes sense to return AST to the user,
// as they might want to make changes to the AST before render, etc.
func ParseFile(path string) (*Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Could not open file %s: %w", path, err)
	}
	defer file.Close()

	return parseFile(file)
}
