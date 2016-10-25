package main
import (
	"fmt"
)


func main() {
	a := 10
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
	//simular el ciclo while
	for i<10{
		fmt.Println(i)
		i++

		if i == 2{
			i++
			continue
		}


		if i>10{
			break
		}
	}

}