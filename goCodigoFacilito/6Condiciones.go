package main 

import (
	"fmt")

/*
== igual a
!= diferente de
<< menor que
> mayor que
<= menor e igual que
>=  mayor o igual que
&& and
 */

func main(){
	x := 10
	y := 10
	if x >= y{
			fmt.Printf("%d es mayor que %d \n",x, y)
	}else if x < y{
		fmt.Println("Entre al else if")
	}	
	else{
		fmt.Println("Entre al else")
	}
	
}