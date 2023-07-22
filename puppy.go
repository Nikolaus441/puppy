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
