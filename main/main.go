package main

import (

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	_"github.com/lib/pq"
	"net/http"

//	"database/sql"
	"log"
	"database/sql"
)


var people  map[int]*person

type person struct{
	Id int
	Name    string
	Vorname string
	Tel     int
	Email   string
	ORT     int
	Region  string
}

func indexHandler(rdr render.Render)  {
	rdr.HTML(200,"index", people)
}

func newHandler(rdr render.Render, r *http.Request)  {

	var info [7]string

	nameM := r.FormValue("name")
	vornameM := r.FormValue("vorname")
	telM := r.FormValue("tel")
	emailM := r.FormValue("email")
	ortM := r.FormValue("ort")
	regionM := r.FormValue("region")

	info[1] = nameM
	info[2] = vornameM
	info[3] = telM
	info[4] = emailM
	info[5] = ortM
	info[6] = regionM

	rdr.HTML(200,"new", info)

}

func endHandler(rdr render.Render, r *http.Request)  {

	rdr.HTML(200, "ende", people)
}


func main()  {

	db,err := sql.Open("postgres", "user=oosy dbname=gorm1 password=oo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from person")
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



	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "html",
		Extensions:[]string{".tmpl",".html"},
		Charset: "UTF-8",
		IndentJSON:true,
	}))

	m.Get("/",indexHandler)
	m.Get("/new", newHandler)
	m.Get("/ende", endHandler)


	m.Run()
}


