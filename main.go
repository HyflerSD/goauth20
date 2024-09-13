package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
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

type Validator interface {
	Validate() (bool, error)
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
	clientID     string
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

	http.HandleFunc("/hello", postGenToken)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server is closed bro")
	} else if err != nil {
		fmt.Printf("Error happened: %s\n", err)
		os.Exit(1)
	}

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
	var a OauthClient = OauthClient{clientID: "1", clientSecret: "secret1"}
	a.Validate(db)
}

func postGenToken(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request Received\n")
	p := r.URL.Query()
	//Parse url and check if its valid then handle the rest
	v := p.Get("name")
	fmt.Printf(v)
	io.WriteString(w, "hello world\n")
}

/** CLIENT STUFF*/

func (a *OauthClient) Validate(db *sql.DB) (bool, error) {
	users, err := db.Query("SELECT COUNT(*) FROM oauth_clients where id = ? AND secret = ?", a.clientID, a.clientSecret)
	if err != nil {
		log.Fatal("Not Found")
		return false, err
	}
	defer users.Close()

	var count int

	for users.Next() {
		if err := users.Scan(&count); err != nil {
			log.Fatal("Not Found")
			return false, err
		}
	}
	fmt.Printf("Count=%d\n", count)
	return true, nil
}

func (a *OauthClient) Create(fd *os.File) (bool, error) {
	//Check id client exists before creating
	return true, nil
}

func (a *OauthClient) Authorize(fd *os.File) (bool, error) {
	//just check if this account is there and valid
	return true, nil
}

func (a *OauthClient) Revoke(fd *os.File) bool {
	return true
}

func (a *OauthClient) Delete(fd *os.File) (bool, error) {
	return true, nil
}

/** USER STUFF*/

func (a *OauthUser) Authorize(fd *os.File) bool {
	//Simiate the sign in popup flow
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
