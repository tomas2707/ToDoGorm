package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jaax2707/ToDoGorm/models"
	"time"
)

func (access *Db) Register(u *models.User) models.User {
	defer access.DB.Create(&u)
	return *u
}

func (access *Db) Login(username string, password string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Printf("%T", token)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}