package main

func add(a, b int) int {
	sum := a + b
	return sum
}
func getElement() (string, string, string) { // in golang we can return multiple values

	return "Javascript", "Java", "Golang"
}
func processIt(fn func(a int) int) { //here fn is a name of a function it is passed as an argument
	//fn is a func is a type

	fn(1)
}
func main() {
	//fmt.Println(add(5, 7))
	//fmt.Println(getElement())

	fn := func(a int) int {
		return 2
	}
	processIt(fn)

}
