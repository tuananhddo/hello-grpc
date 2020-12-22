package repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id    int    `db:"id"`
	First string `db:"email"`
	//CreatedAt time.Time `db:"created_at"`
}
type ToDo struct {
	Id      int    `db:"id"`
	Content string `db:"content"`
}

func CreateToDoRepo() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "anhdt:anhdt@123@(127.0.0.1:3306)/LOCALDB?parseTime=true&loc=Asia%2FHo_Chi_Minh")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Connected")
	//var user []User

	err = Transact(db, func(tx *sqlx.Tx) error {
		insertQuery := "INSERT INTO todo (content) VALUES (:content)"
		_, err := tx.NamedExec(insertQuery, ToDo{Content: "Water"})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
func Transact(db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	err = fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
