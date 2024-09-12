package main

import (
	//"errors"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"os"
)

type Auth interface {
	Authorize(fd *os.File) error
	Revoke(fd *os.File) error
}

type Creator interface {
	Create(fd *os.File) error
	Delete(fd *os.File) error
}

type User struct {
	email    string
	password string
	id       int32
}

type TempAuthCodes struct {
	clientID int32
	userID   int32
	scopes   []string
	authCode string
	expires  string
}

type OauthClient struct {
	clientID     int32
	clientSecret string
	isActive     bool
	scopes       []string
	redirectURI  string
}

type OauthUser struct {
	userID      int32
	clientID    int32
	accessToken string
	scope       []string
	expires     string
}

var db *sql.DB

func main() {

	cfg := mysql.Config{
		User:   os.Getenv("GODBUSER"),
		Passwd: os.Getenv("GODBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("GOAUTHDB"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connected!")
}

/** CLIENT STUFF*/
func (a *OauthClient) Create(fd *os.File) bool {
	return true
}

func (a *OauthClient) Authorize(fd *os.File) bool {

	return true
}

func (a *OauthClient) Revoke(fd *os.File) bool {
	return true
}

func (a *OauthClient) Delete(fd *os.File) bool {
	return true
}

/** USER STUFF*/

func (a *OauthUser) Authorize(fd *os.File) bool {
	return true
}

func (a *OauthUser) Revoke(fd *os.File) bool {
	return true
}

/** GENERAL SERVER NEEDS */

func (a *OauthUser) GenerateToken() string {

	token := uuid.New().String()

	return token
}
