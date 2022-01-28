package main

// func Increase() func() int {
// 	n := 0
// 	return func() int {
// 		n++
// 		return n
// 	}
// }

// func Add() *int {
// 	c := 0
// 	return &c
// }

func main() {
	var l *int
	for i := 0; i < 10; i++ {
		l = new(int)
		*l = i
	}
	// println(*l)
}

// func main() {

// 	var l *int
// 	func() {
// 		l = new(int)
// 		*l = 1
// 	}()

// 	// Increase()

// 	// in()

// 	// Add()

// }
