package datatype

import (
	"fmt"
	"reflect"
)

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func Test() {
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}

	var person PersonInfo
	var ok bool
	person, ok = personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}

	fmt.Println(reflect.TypeOf(person))

	person, ok = personDB["12345"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 12345.")
	} else {
		fmt.Println("Did not find person with ID 12345.")
	}

	person, ok = personDB["123456"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 123456.")
	} else {
		fmt.Println("Did not find person with ID 123456.")
	}
}
