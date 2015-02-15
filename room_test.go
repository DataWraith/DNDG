package main

import (
	"testing"
)

func TestRoomsShouldDescribeThemselves(t *testing.T) {
	g := NewGamestate(1)

	r := Room{
		DescriptionFuncs: []DescriptionFunc{
			func(g *Gamestate) string {
				return "You are in a maze of twisty little passages, all alike"
			},
		}}

	if r.Description(g) != "You are in a maze of twisty little passages, all alike" {
		t.Fatal("expected the description of r to be `You are in a maze of twisty little passages, all alike`")
	}
}

func TestRoomsShouldDescribeThemselvesWithMultipleStanzas(t *testing.T) {
	g := NewGamestate(1)

	r := Room{
		DescriptionFuncs: []DescriptionFunc{
			func(g *Gamestate) string { return "foo bar baz quux" },
			func(g *Gamestate) string { return "xyzzy" },
		},
	}

	if r.Description(g) != "foo bar baz quux\n\nxyzzy" {
		t.Fatalf("expected description of testRoom with two stanzas to be `foo bar baz quux\\n\\nxyzzy`, got %q", r.Description(g))
	}
}
