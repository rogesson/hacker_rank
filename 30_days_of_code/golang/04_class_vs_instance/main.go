package main
import "fmt"

type person struct {
    age int
}

func (p person) NewPerson(initialAge int) person {
    if initialAge < 0 {
        fmt.Println("Age is not valid, setting age to 0.")
        initialAge = 0
    }

    p.age = initialAge
    return p
}

func (p person) amIOld() (message string) {
    switch age := p.age; {
    case age < 13:
        message = "You are young."
    case age >= 13 && age < 18:
        message = "You are a teenager."
    default:
        message = "You are old."
    }

    return
}

func (p person) yearPasses() person {
    p.age += 1

    return p
}

func main() {
    var T, age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()

        for j := 0; j < T; j++ {
            p = p.yearPasses()
        }

        message := p.amIOld()
        fmt.Println(message)
    }
}
