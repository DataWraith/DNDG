package main

import (
	"testing"
)

func TestGamestateHasFlagsThatCanBeManipulated(t *testing.T) {
	g := NewGamestate(1)

	if g.HasFlag("foo") {
		t.Fatal("empty gamestate is not supposed to have flag 'foo'")
	}

	if g.HasFlag("bar") {
		t.Fatal("empty gamestate is not supposed to have flag 'bar'")
	}

	g.SetFlag("foo")

	if !g.HasFlag("foo") {
		t.Fatal("expected SetFlag(\"foo\") to set the 'foo' flag")
	}

	g.UnsetFlag("foo")

	if g.HasFlag("foo") {
		t.Fatal("expected UnsetFlag(\"foo\") too unset the 'foo' flag")
	}

	g.ToggleFlag("bar")

	if !g.HasFlag("bar") {
		t.Fatal("expected ToggleFlag(\"bar\") to set the 'bar' flag")
	}

	g.ToggleFlag("bar")

	if g.HasFlag("bar") {
		t.Fatal("expected second invocation of ToggleFlag(\"bar\") to unset the 'bar' flag")
	}
}
