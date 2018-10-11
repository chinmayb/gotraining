package ops

import "fmt"
//import "os"

type Person struct {
    Name string `json:"name"`
    Age int     `json:"age"`
    Telephone string `json:"phone_number"`
}

type People []Person

func New(name string, age int, tele string) Person{
    return Person{name, age, tele}
}

func (p Person)String() string{
    str1 := fmt.Sprintf("\n%s,%3d,%s", p.Name, p.Age, p.Telephone)
    return str1
}

//implemt Write
