package main

import (
	"fmt"

	"github.com/wagnersilvafilho/aprendagolang/imersao/db"
	"github.com/wagnersilvafilho/aprendagolang/imersao/users"
)

func main() {
	app := users.NewApp("aprenda-golang")
	fmt.Println(app)
	u := users.NewUser("Wagner Filho", "wagnerfilho", "", 14)
	fmt.Println(u)

	db.Save(app)
	err := db.Save(u)
	fmt.Println(err)
}
