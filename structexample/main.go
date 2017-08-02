package main


import "fmt"
import "./ops"
import "os"

var personname string = "Person"
var Age int = 10
var number string = "900812250"

func generatePeople(i int) (string, int, string){
    name := fmt.Sprintf("%s%d", personname, i)
    age := Age + i
    number := fmt.Sprintf("%s%d", number, i)

    return name, age, number
}


func writeToFile(people ops.People, ){
    f, _ := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
    defer f.Close()
    str1 := fmt.Sprintf("Name,Age,Tel")
    f.WriteString(str1)
    for _,p := range people {
        str1 = fmt.Sprintf("%v", p)
        f.WriteString(str1)
    }

}

func main() {
    entries := 4
    ppl := make(ops.People, entries)
    for i:=0; i< entries; i++ {
        name, age, tel := generatePeople(i)
        ppl[i] = ops.New(name, age, tel)
    }
    fmt.Printf("%v", ppl)
    writeToFile(ppl)
}

// use Fprintf , write to a writer