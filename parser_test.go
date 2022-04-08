package parser

import (
	"reflect"
	"testing"
)

func TestToLines_with_single_word(t *testing.T) {
	got := ToLines("ToLines")
	want := []string{"ToLines"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestToLines_should_strip_single_world(t *testing.T) {
	got := ToLines(`
		Pursue Wealth
	`)
	want := []string{"Pursue Wealth"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestToLines_should_strip_each_line(t *testing.T) {
	got := ToLines(`
		Pursue Wealth
		Not Money
	`)
	want := []string{"Pursue Wealth", "Not Money"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestToLines_should_handle_empty(t *testing.T) {
	got := ToLines("")
	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestParse_single_comment(t *testing.T) {
	got := Parse(`
		// TODO: We will fix this someday
	`)
	want := []string{"TODO: We will fix this someday"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestParse_multiple_comment(t *testing.T) {
	got := Parse(`
		/*
			TODO: We will fix this someday
			I mean, really.
		*/
		let wrong = 1;

		// TODO: Yes we will fix this. Soon.
		wrong = 2;

		if (!wrong) return null; // TODO: you might want to throw an error here?

		/* TODO: Stringify output */
		console.log(wrong);
	`)
	want := []string{
		"TODO: We will fix this someday", "TODO: Yes we will fix this. Soon.", "TODO: you might want to throw an error here?", "TODO: Stringify output",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
