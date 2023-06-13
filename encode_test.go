package env

import (
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	var entity struct {
		String string `env:"HELLO"`
	}
	entity.String = "hello"
	if err := Encode(&entity); err != nil {
		t.Fatal(err)
	}
	if os.Getenv("HELLO") != "hello" {
		t.Errorf("expected %q, got %q", "hello", os.Getenv("HELLO"))
	}
}
