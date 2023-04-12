package data

import (
	"database/sql"
	_ "io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (*sql.DB, error) {
	//uri := "postgres://hug58:grosss213@127.0.0.1:5432/microvideogame?sslmode=disable"
	//pgConnString := os.Getenv("URI")

	return sql.Open("mysql", "root:1234@tcp(localhost:3306)/rickandmorty")
}

/*
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}


	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
*/
