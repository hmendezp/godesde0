package users

import (
	"fmt"
	"time"

	"github.com/hmendezp/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(11, "Checo", time.Now(), true)
	fmt.Println(u)
}
