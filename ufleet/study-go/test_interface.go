package main
import (
    "fmt"
    "strconv"
)
type Human struct {
    name string
    age int
    phone string
}
type Dog struct {
    name string
    color string
}

func (d Dog) String() string {
    return "the dog, name is:" + d.name + ",color is:{}"
 
}
// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
    return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}
func main() {
    Bob := Human{"Bob", 39, "000-7777-XXX"}
    fmt.Println("This Human is : ", Bob)
    dog := Dog{"haqi", "red"}
    fmt.Println("Anminal ", dog)
}
