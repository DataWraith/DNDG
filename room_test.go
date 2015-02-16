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

	if r.Description(g) != "You are in a maze of twisty little passages, all alike\n\n" {
		t.Fatal("expected the description of r to be `You are in a maze of twisty little passages, all alike\n\n`")
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

	if r.Description(g) != "foo bar baz quux\n\nxyzzy\n\n" {
		t.Fatalf("expected description of testRoom with two stanzas to be `foo bar baz quux\\n\\nxyzzy\\n\\n`, got %q", r.Description(g))
	}
}

func TestRoomsShouldHaveAvailableActions(t *testing.T) {
	g := NewGamestate(1)

	r := Room{
		DescriptionFuncs: []DescriptionFunc{
			func(g *Gamestate) string { return "To the north is a cave." },
		},

		Actions: []Action{
			Action{
				Command: "go north",
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 1
					return true
				},
			},
		},
	}

	r.ExecuteAction("go north", g)

	if g.CurrentRoom != 1 {
		t.Fatalf("expected action 'go north' to change the current room to 001, got %3d", g.CurrentRoom)
	}
}
