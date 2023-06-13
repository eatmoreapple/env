package env

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	var entity struct {
		String string  `env:"string"`
		Bool   bool    `env:"bool"`
		Int    int     `env:"int"`
		Uint   uint    `env:"uint"`
		Float  float64 `env:"float"`
		Struct struct {
			String string `env:"string"`
		}
	}
	if err := os.Setenv("string", "string"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("bool", "true"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("int", "1"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("uint", "1"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("float", "1.1"); err != nil {
		t.Fatal(err)
	}
	if err := Decode(&entity); err != nil {
		t.Fatal(err)
	}
	if entity.String != "string" {
		t.Errorf("expected %q, got %q", "string", entity.String)
	}
	if entity.Bool != true {
		t.Errorf("expected %t, got %t", true, entity.Bool)
	}
	if entity.Int != 1 {
		t.Errorf("expected %d, got %d", 1, entity.Int)
	}
	if entity.Uint != 1 {
		t.Errorf("expected %d, got %d", 1, entity.Uint)
	}
	if entity.Float != 1.1 {
		t.Errorf("expected %f, got %f", 1.1, entity.Float)
	}
	if entity.Struct.String != "string" {
		t.Errorf("expected %q, got %q", "string", entity.Struct.String)
	}
	t.Logf("entity: %+v", entity)
}

func TestDefault(t *testing.T) {
	var entity struct {
		String string `env:"string" default:"hello"`
	}
	if err := Decode(&entity); err != nil {
		t.Fatal(err)
	}
	if entity.String != "hello" {
		t.Errorf("expected %q, got %q", "hello", entity.String)
	}
}
