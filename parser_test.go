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

func TestParse(t *testing.T) {
	got := Parse(`
		// TODO: We will fix this someday
	`)
	want := []string{"TODO: We will fix this someday"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
