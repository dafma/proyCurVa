package main 

import "fmt"


//parecido a la herencia en po
type Human struct{
	name string
}

func (this Human) hablar() string{
	return "Bla bla bla"
}

type Dummy struct{
	name string
}

type Tutor struct{
	Human
	Dummy
}

func main() {
	tutor := Tutor{Human{"Uriel"}, Dummy{"Marco"}}
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.hablar())
}