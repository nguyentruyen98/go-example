package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type logWriter struct {
}

func main() {
	// resp, err := http.Get("http://google.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// bs := make([]byte, 99999)
	// amb, v := resp.Body.Read(bs)
	// fmt.Println(amb)
	// fmt.Println(v)
	// fmt.Println(string(bs))

	// lw := logWriter{}
	// io.Copy(lw, resp.Body)

	triangle := triangle{10, 5}
	square := square{10}
	printArea(triangle)
	printArea(square)

}
func printArea(s shape) {
	fmt.Println(s.getArea())

}
func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (logWriter) Write(bs []byte) (int, error) {
	// fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
