package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	//ID    *int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}

var Db *sqlx.DB

func InitDB() {
	var err error
	Db, err = sqlx.Connect("postgres", "host=172.17.0.2 port=5432 user=spondon password=1234 dbname=temp_db sslmode=disable")
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
	}

}

func Insert(user User) error {

	_, err := Db.NamedExec(`INSERT INTO users (name, email, age) VALUES (:name, :email, :age)`, user)

	return err
}

func Select() ([]User, error) {

	user, err := Db.Query(`SELECT * FROM users`)

	fmt.Printf("%v", user)

	var all_user []User
	for user.Next() {
		var temp_user User
		err = user.Scan(&temp_user.Name, &temp_user.Email, &temp_user.Age)
		all_user = append(all_user, temp_user)
	}
	fmt.Println(" ")
	fmt.Printf("%v", all_user)
	fmt.Println(" ")
	return all_user, err
}
