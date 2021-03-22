package example_coverage

import "fmt"

func Force() string {
	return "<h1>May the force be with you</h1>"
}

func main() {
	fmt.Println(Force())
}
