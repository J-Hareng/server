package main

import (
	"fmt"
	email "server/src/api/Email"
	"server/src/api/db"
	"server/src/helper"
	"server/src/httpd"
	"server/src/httpd/security"
)

func main() {

	// * set Up buffer
	DB, err := db.New()
	if err != nil {
		helper.CustomError(err.Error())
	}
	EKM := security.EmailTokenMap{}
	E := email.GeneratEmail()
	TM := security.TokenMap{}

	fmt.Print(DB) //! just for testing

	s := httpd.Init(DB, &TM, E, &EKM)

	s.Run("8080")
}
