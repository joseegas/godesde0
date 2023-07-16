package users

import (
	"fmt"
	"godesde0/modelos"
	"time"
)

func AltaUsuario(){
	u:=new(modelos.User)
	u.AddUser(10,"Pablo",time.Now(),true)
	fmt.Println(u)
}