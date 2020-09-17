package authentication

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/tinnaing347/interval-flux/models"
	"github.com/tinnaing347/interval-flux/query"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccountHandler(c *gin.Context) {
	var input CreateAccount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account_exist, err := input.ValidateUsername()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if account_exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
		return
	}

	account := &Account{
		UserName: input.UserName,
		Password: HashedPassword(input.Password),
		Role:     input.Role,
	}

	tags, fields, time_ := account.TagField()

	models.CreateBatchPoint("ivdb", "user_data", tags, fields, time_)
	c.JSON(http.StatusOK, gin.H{"data": "user created"})
}

func (r *CreateAccount) ValidateUsername() (bool, error) {
	var b bytes.Buffer
	b.WriteString("SELECT * FROM user_data ")

	check_empty := r == &CreateAccount{}
	if check_empty {
		return false, errors.New("CreateAccount struct is empty.")
	}

	b, _ = query.HandleTagFieldQuery(b, "username", r.UserName, true)

	q := client.NewQuery(b.String(), "ivdb", "")

	response, err := models.DB.Query(q)

	if err != nil {
		return false, err
	}

	if response.Error() != nil {
		return false, err
	}

	if len(response.Results[0].Series) == 0 {
		return false, nil
	}

	return true, nil
}

func LoginHanlder(username, password string) (*User, error) {

	account, err := GetAccount(username)
	if err != nil {
		return &User{}, err
	}
	flag := CheckPasswordHash(password, account.Password)
	if flag {
		return &User{
			UserName: account.UserName,
			Role:     account.Role,
			Created:  account.Created,
		}, nil
	}

	return &User{}, errors.New("Password is incorrect.")
}

func GetAccount(username string) (*Account, error) {
	var b bytes.Buffer
	b.WriteString("SELECT * FROM user_data ")

	b, _ = query.HandleTagFieldQuery(b, "username", username, true)
	q := client.NewQuery(b.String(), "ivdb", "")
	response, err := models.DB.Query(q)

	if err != nil {
		return &Account{}, err
	}

	if response.Error() != nil {
		return &Account{}, err
	}

	if len(response.Results[0].Series) == 0 {
		return &Account{}, errors.New("Username does not exist")
	}

	account := NewAccount(response.Results[0].Series[0].Columns, response.Results[0].Series[0].Values[0])

	return account, nil
}

func HashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
