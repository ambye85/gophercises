package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "gopher"
	password = "gopher"
	dbname = "gophercises"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = resetDB(db, dbname)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = createDB(db, dbname)
	//if err != nil {
	//	panic(err)
	//}
	//
	//db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = createPhoneNumbersTable(db)
	if err != nil {
		panic(err)
	}

	id, err := insertPhone(db, "1234567890")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted phone numner with ID #%d\n", id)
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return nil
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		);`
	_, err := db.Exec(statement)
	return err
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id;`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

//func normalize(phone string) string {
//	var buf bytes.Buffer
//	for _, ch := range phone {
//		if ch >= '0' && ch <= '9' {
//			buf.WriteRune(ch)
//		}
//	}
//
//	return buf.String()
//}
