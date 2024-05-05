package tree_sitter_plantuml_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-plantuml"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_plantuml.Language())
	if language == nil {
		t.Errorf("Error loading Plantuml grammar")
	}
}
