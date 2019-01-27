package main
import "fmt"
type Obj struct {}

func a() *Obj {
	r := Obj{}
	return &r
}
func b() {
	o := a()
	fmt.Printf("%#v\n", o)
}
func main() {
	b()
}
