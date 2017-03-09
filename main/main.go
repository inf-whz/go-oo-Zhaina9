package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/codegangsta/martini"
//	"database/sql"
	"database/sql"
)

type person2 struct {
	/*Id  		int	`sql:"AUTO_INCREAMENT" gorm:"primary_key"`
	Name 		string	`sql:"size:255;unique;index"`
	Vorname 	string	`sql:"size:255;unique;index"`
	Tel	 	string	`sql:"size:255;unique;index"`
	Email 		string	`sql:"size:255;unique;index"`
	ORT	 	string	`sql:"size:255;unique;index"`
	Region 		string	`sql:"size:255;unique;index"`
*/
	Id  		int		`sql:"AUTO_INCREAMENT" gorm:"primary_key"`
	Name 		string
	Vorname 	string
	Tel	 	string
	Email 		string
	ORT	 	string
	Region 		string

}

type session struct {
	Inf string
}

//var database []*person2 = make([]*person2,0)
var db *gorm.DB
var err error
//var sessionData person2 = person2{Id:-1}

var people  map[int]*person2

func index(rdr render.Render){
	rdr.HTML(200,"index", people)
}

var data [7]string

func newHandler(rdr render.Render, r *http.Request)  {

	//infM := r.FormValue("inf")
	nameM := r.FormValue("name")
	vornameM := r.FormValue("vorname")
	telM := r.FormValue("tel")
	emailM := r.FormValue("email")
	ortM := r.FormValue("ort")
	regionM := r.FormValue("region")

	//info[0] = infM
	data[1] = nameM
	data[2] = vornameM
	data[3] = telM
	data[4] = emailM
	data[5] = ortM
	data[6] = regionM


	rdr.HTML(200,"new", data)

}

func checkInputHandler(rdr render.Render, r *http.Request)  {

	db, err = gorm.Open("postgres","user=oosy password=oo dbname=gorm1 host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}


	
	db.CreateTable(person2{})
	jayna := person2{Name:data[1], Vorname:data[2], Tel:data[3], Email:data[4], ORT:data[5], Region:data[6]}
	db.Create(&jayna)
	/*u := []person2{}
	db.Find(&u)

	for _,value := range u{
		fmt.Println(value)
	}

	db.Get(&jayna)*/

	rdr.HTML(200, "ende", people)

}

func admin(rdr render.Render)  {
	rdr.HTML(200,"admin", people)
}

func admin_show(rdr render.Render){

	db,err := sql.Open("postgres", "user=oosy dbname=gorm1 password=oo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	people := make(map[int]*person2,0)

	for rows.Next()  {

		pers := new(person2)

		err := rows.Scan(&pers.Id, &pers.Name, &pers.Vorname, &pers.Tel, &pers.Email, &pers.ORT, &pers.Region)

		if err != nil{
			log.Fatal(err)
		}
		people[pers.Id] = pers
	}


	rdr.HTML(200,"admin_show", people)
}


func main()  {
	fmt.Println("Application starts!")

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "html",
		Extensions:[]string{".tmpl",".html"},
		Charset: "UTF-8",
		IndentJSON:true,
	}))

	m.Get("/", index)
	m.Get("/new", newHandler)
	m.Get("/ende", checkInputHandler)
	m.Get("/admin", admin)
	m.Get("/admin_show", admin_show)


	m.Run()
}