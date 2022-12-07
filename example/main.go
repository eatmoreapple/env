package main

import (
	"fmt"
	"github.com/eatmoreapple/env"
	"log"
	"os"
)

func main() {
	if err := os.Setenv("hello", "world"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("bool", "true"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(env.Name("hello").String()) // world

	fmt.Println(env.Name("bool").Bool()) // true

	fmt.Println(env.Name("undefined").StringOrElse("undefined")) // undefined

	var entity struct {
		Hello string `env:"hello"`
		Bool  bool   `env:"bool"`
	}
	if err := env.Decode(&entity); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", entity) // {Hello:world Bool:true}
}
