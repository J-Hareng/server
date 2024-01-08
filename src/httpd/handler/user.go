package handler

import (
	"fmt"
	"net/http"
	"server/src/api/db"
	"server/src/helper"
	"server/src/httpd/bodymodels"
	"server/src/httpd/security"

	"github.com/gin-gonic/gin"
)

func GetUsers(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := db.GetAllUsers()
		if err != nil {
			// Handle the error appropriately
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		c.JSON(http.StatusOK, val)
	}
}
func AddUser(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user bodymodels.AddUserMod

		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		is_used, err := db.AvalabileEmail(user.Email)
		if err != nil {
			fmt.Println(err)
			fmt.Println(is_used)
		}

		fmt.Print("is used : ")
		fmt.Println(is_used)

		fmt.Print("user : ")
		fmt.Println(user)

		if is_used {
			fmt.Println(err)
			fmt.Println(is_used)

			c.JSON(http.StatusConflict, gin.H{"error": "Email already used"})
			return

		}

		db.AddUser(user.Name, user.Email, user.Password)

		c.JSON(http.StatusOK, user)
	}
}

// TODO: func checkUser(db *db.DB , id primitive.ObjectID){ a }
func RequestSessionToken(db *db.DB, TM *security.TokenMap) gin.HandlerFunc {

	return func(c *gin.Context) {
		var userLogin bodymodels.RequestSessionTokenMod

		if err := c.ShouldBindJSON(&userLogin); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			helper.CustomError(err.Error())
			return
		}

		user, err := db.GetUser(userLogin.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password wrong"})
			return
		}

		if userLogin.Password != user.PASSWORD {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password wrong"})
			return
		}
		token, err := security.CreateUserToken(user.NAME, user.ID, TM)
		if err != nil {
			//!NEED TO REMOVE TO LOG THE ERROR (because of security for future me u dumb ass )
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.SetCookie("token", token, 3600, "/", "", false, false)

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

		//!Just for testing
		fmt.Println(TM.TOKENTOINDEX)
	}

}

// TODO: need to Validate user Token
func ValidateUserToken(db *db.DB, TM *security.TokenMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		fmt.Println(token)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		fmt.Println(token)
		if TM.TOKENTOINDEX[token] != nil {
			fmt.Println(*TM.TOKENTOINDEX[token])
			c.JSON(http.StatusOK, gin.H{"message": "Save"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "SessionInvalid"})
	}
}

// //!needs to bee removed because of safty
// func CheckEmail(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context){
// 		db.AvalabileEmail("a")
// 	}

// }