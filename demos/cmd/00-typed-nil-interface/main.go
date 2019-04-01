package main

import "log"

type Runner interface {
	Run()
}

type FooRunner struct{}

func (*FooRunner) Run() {}

func NilFooRunner() *FooRunner {
	return nil
}

func main() {
	var f Runner = NilFooRunner()
	if f != nil {
		log.Fatalf("Expected f == nil, was: %v", f)
	}
	log.Println("Everything is fine")
}
