package puppy

import "fmt"

func main() {
	fmt.Println("See comments for details on module")

}

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Test_item() string {
	fmt.Println("Running test item")
	return "test"
}

func From13() {
	fmt.Println("I am from v1.3.0")
}
