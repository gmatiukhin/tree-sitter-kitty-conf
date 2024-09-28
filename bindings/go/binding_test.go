package tree_sitter_kittyconf_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_kittyconf "github.com/tree-sitter/tree-sitter-kittyconf/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_kittyconf.Language())
	if language == nil {
		t.Errorf("Error loading Kittyconf grammar")
	}
}
