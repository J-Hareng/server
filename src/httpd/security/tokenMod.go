package security

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserToken struct {
	USERID    primitive.ObjectID
	USERNAME  string
	RAW_TOKEN *jwt.Token
	TOKEN     string
}

type TokenMap struct {
	USERS        []UserToken
	TOKENTOINDEX map[string]*UserToken
	USERINDEX    map[string]*UserToken
}

func (tm *TokenMap) AddUserByUsername(user *UserToken) {
	if tm.TOKENTOINDEX == nil {
		tm.TOKENTOINDEX = make(map[string]*UserToken)
	}

	if tm.USERINDEX == nil {
		tm.USERINDEX = make(map[string]*UserToken)
	}
	tm.TOKENTOINDEX[user.TOKEN] = user
	tm.USERINDEX[user.USERNAME] = user
}

type EmailTokenMap struct {
	Keys map[string]string
}

func (EKM *EmailTokenMap) ValidateEmail(key string, email string) bool {
	if EKM.Keys[key] == email {
		delete(EKM.Keys, key)
		return true
	}
	return false
}
