package colon

import "testing"

func TestEmptyTemplate(t *testing.T) {
	render := Compile("")
	s := render(map[string]interface{}{"foo": "bar", "aaa": "bbb"})
	if s != "" {
		t.Fatalf("unexpected string returned %q. Expected empty string", s)
	}
}

func TestSimpleTemplate(t *testing.T) {
	render := Compile("hi :name")
	s := render(map[string]interface{}{"name": "matt"})
	if s != "hi matt" {
		t.Fatalf("unexpected string returned %q", s)
	}
}

func TestMultipleTemplate(t *testing.T) {
	render := Compile("hi :name & :name")
	s := render(map[string]interface{}{"name": "matt"})
	if s != "hi matt & matt" {
		t.Fatalf("unexpected string returned %q", s)
	}
}

func TestIgnoresTemplate(t *testing.T) {
	render := Compile("hi ::name & :name")
	s := render(map[string]interface{}{"name": "matt"})
	if s != "hi ::name & matt" {
		t.Fatalf("unexpected string returned %q", s)
	}
}

func TestIgnores2Template(t *testing.T) {
	render := Compile("hi \\:name & :name")
	s := render(map[string]interface{}{"name": "matt"})
	if s != "hi \\:name & matt" {
		t.Fatalf("unexpected string returned %q", s)
	}
}

func TestNumbersTemplate(t *testing.T) {
	render := Compile("i am :age years old")
	s := render(map[string]interface{}{"age": 27})
	if s != "i am 27 years old" {
		t.Fatalf("unexpected string returned %q", s)
	}
}

func TestNestedTemplate(t *testing.T) {
	render := Compile("i prefer :settings.color")
	s := render(map[string]interface{}{"settings": map[string]interface{}{"color": "red"}})
	if s != "i prefer red" {
		t.Fatalf("unexpected string returned %q", s)
	}
}
