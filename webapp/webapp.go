package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:12345678@/webapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usersHandler := func(w http.ResponseWriter, r *http.Request) {
		var (
			id    int
			login string
		)
		rows, err := db.Query("select id, login from users where status = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &login)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w, strconv.Itoa(id)+" "+login+"\n")
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	slashHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "slash.html")
	}

	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		key, present := query["login"]
		if !present || len(key) == 0 {
			http.HandleFunc("/", slashHandler)
		}
		rows, err := db.Query("select * from users inner join pwds on users.id = pwds.id where login = ?", key[0])
		if err != nil {
			log.Fatal(err)
		}
		var (
			id           int
			login        string
			money_amount int
			card_number  string
			status       int
			pwd          string
		)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &login, &money_amount, &card_number, &status, &id, &pwd)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w, strconv.Itoa(id)+" "+login+" "+strconv.Itoa(money_amount)+" "+
				card_number+" "+strconv.Itoa(status)+" "+pwd+"\n")
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	idHandler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		key, present := query["id"]
		if !present || len(key) == 0 {
			http.HandleFunc("/", slashHandler)
		}
		rows, err := db.Query("select * from users inner join pwds on users.id = pwds.id where users.id = ?", key[0])
		if err != nil {
			log.Fatal(err)
		}
		var (
			id           int
			login        string
			money_amount int
			card_number  string
			status       int
			pwd          string
		)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &login, &money_amount, &card_number, &status, &id, &pwd)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(w, strconv.Itoa(id)+" "+login+" "+strconv.Itoa(money_amount)+" "+
				card_number+" "+strconv.Itoa(status)+" "+pwd+"\n")
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", slashHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/by-login", loginHandler)
	http.HandleFunc("/by-id", idHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
