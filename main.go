package main

import "github.com/graniticio/granitic/v2"
import "scrumize/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
