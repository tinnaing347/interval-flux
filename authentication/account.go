package authentication

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type CreateAccount struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

//login info
type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//account for database
type Account struct {
	UserName string    `mapsturcture:"username" json:"time"`
	Password string    `mapsturcture:"password" json:"password"`
	Role     string    `mapsturcture:"role" json:"role"`
	Created  time.Time `mapsturcture:"time" json:"time"`
}

//passed around in jwt
type User struct {
	UserName string
	Role     string
	Created  time.Time
}

func (i *Account) TagField() (map[string]string, map[string]interface{}, time.Time) {

	tags := map[string]string{"role": i.Role}
	fields := map[string]interface{}{
		"username": i.UserName,
		"password": i.Password,
	}
	return tags, fields, time.Now()
}

func Serializer(columns []string, values []interface{}) map[string]interface{} {
	map_ := make(map[string]interface{})

	for i := 0; i < len(columns); i++ {
		map_[columns[i]] = values[i]
	}
	return map_
}

func NewAccount(columns []string, values []interface{}) *Account {

	account_map := Serializer(columns, values)

	var account Account

	if err := mapstructure.Decode(account_map, &account); err != nil {
		panic(err)
	}
	return &account
}
