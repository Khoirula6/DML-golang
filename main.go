package main

import (
	"belajar-database/entity"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "muhammad"
	dbname   = "khairUniversity"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successfully connect!!!")
	}
	return db
}

func main() {
	// student := entity.Student{Id: 7, Name: "Mbak", Email: "mbak@mail.com", Address: "Solo", BirthDate: time.Date(2001, 10, 10, 0, 0, 0, 0, time.Local), Gender: "M"}

	// addStudent(student)
	// updateStudent(student)
	deleteStudent("7")

}

func addStudent(student entity.Student) {

	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("berhasil ditambahkan")
	}
}

func updateStudent(student entity.Student) {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_student SET id= $1, email= $3, address = $4, birth_date = $5, gender = $6 WHERE name = $2;"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("berhasil di update")
	}
}

func deleteStudent(id string) {
	db := connectDB()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_student WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		println("berhasil didelete")
	}
}
