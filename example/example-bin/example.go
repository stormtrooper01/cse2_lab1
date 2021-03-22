package example_coverage

import "fmt"

func Hello() string {
	return "<h1>May the force be with you</h1>"
}

func main() {
	fmt.Println(Hello())
}
