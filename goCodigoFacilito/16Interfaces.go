package main 
import "fmt"


type User interface{
	Permissions () int //1-5
	Name () string
}

type Admin struct{
	name string
}

func (this Admin) Permissions() int{
	return 5
}

func (this Admin) Name() string{
	return this.name
}

func auth(user User) string{
	if user.Permissions() > 4{
		return user.Name() + "Tiene permisos de administrador"
	}
	return ""
}


func main() {
	admin := Admin{"Marco"}
	fmt.Println(auth(admin))
}