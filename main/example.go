package main

import(
	_"github.com/lib/pq"
	"database/sql"
	"log"
	)

type person struct{
	Id      int
	Name    string
	Vorname string
	Tel     int
	Email   string
	ORT     int
	Region  string
}

func main(){
	db,err := sql.Open("postgres", "user=oosy dbname=gorm1 password=oo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * fron person")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	people := make(map[int]*person,0)

	for rows.Next()  {

		pers := new(person)

		err := rows.Scan(&pers.Id, &pers.Name, &pers.Vorname, &pers.Tel, &pers.Email, &pers.ORT, &pers.Region)

		if err != nil{
			log.Fatal(err)
		}
		people[pers.Id] = pers
	}


}
