package main

import (
	"fmt"
	"server/src/api/db"
	"server/src/helper"
	"server/src/httpd"
	"server/src/httpd/security"
)

func main() {

	DB, err := db.New()

	if err != nil {
		helper.CustomError(err.Error())
	}

	fmt.Print(DB) //! just for testing

	TM := security.TokenMap{}

	s := httpd.Init(DB, &TM)

	s.Run("8080")
}
