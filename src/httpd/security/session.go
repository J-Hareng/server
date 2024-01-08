package security

import (
	"fmt"
	"server/src/helper"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateRAWToken(username string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time
	})
	return token
}
func cryptToken(token *jwt.Token, secretKey []byte) (string, error) {
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateUserToken(userName string, userID primitive.ObjectID, tm *TokenMap) (string, error) {
	key := []byte(helper.GetEnvVar("SECRETKEY"))
	fmt.Println("______________")

	rawT := generateRAWToken(userName)

	T, err := cryptToken(rawT, key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	UT := UserToken{
		USERID:    userID,
		USERNAME:  userName,
		RAW_TOKEN: rawT,
		TOKEN:     T,
	}
	if tm.USERS == nil {
		tm.USERS = make([]UserToken, 0)
	}

	tm.USERS = append(tm.USERS, UT)
	tm.AddUserByUsername(&UT)

	fmt.Println(UT)
	fmt.Println(*tm)

	return UT.TOKEN, nil
}
