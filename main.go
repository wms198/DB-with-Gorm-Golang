package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/initD"
	"net/http"

	"gorm.io/gorm"
)

type Book struct {
	Id    int
	Title string
}

type GoTestModel struct {
	gorm.Model
	UserId uint
	Name   string
	Year   string
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	// GoTestModelID int
	// Constrains
	//oneToone
	GoTestModel GoTestModel `gorm:"foreignKey:UserId"`

	// ManyToMany
	Book []Book `gorm:"many2many:user_book"`
}

func init() {
	initD.ConnectDB()
}
func main() {

	// http.HandleFunc("/createstuff", GoDatabaseCreate)
	http.HandleFunc("/createuser", CreateUser)
	initD.DB.AutoMigrate(&User{}, &GoTestModel{}, Book{})
	fmt.Println("Tables created")

	fmt.Println("Server running on localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))

}

// Create single items
// func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
// 	GoTestModel := GoTestModel{
// 		Name: "Mike",
// 		Year: "2021",
// 	}
// 	result := initD.DB.Create(&GoTestModel)
// 	if result.Error != nil {
// 		log.Fatalln(result.Error)
// 	}
// 	json.NewEncoder(w).Encode(GoTestModel)
// 	fmt.Println("Fields Added", GoTestModel)
// }

// Create mutiole items
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//create multiple records with Create():
	//You cannot pass a struct to ‘create’, so you should pass a pointer to the data.
	users := []*User{
		{FirstName: "Mike", LastName: "White"},
		{FirstName: "Lili", LastName: "Moll"},
		{FirstName: "Mila", LastName: "Woo"},
	}
	/* 	User := User{
		FirstName: "Mike",
		LastName:  "white",
	} */
	result := initD.DB.Create(users)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	json.NewEncoder(w).Encode(users)
	fmt.Println("User Added", users)
}

// Get the Last record ordered by primary key
func showLastRecird() {
	user := User{}
	initD.DB.Last(&user)

	fmt.Println(user)
}

// Fetch record using where condition
func fetchUseWhere() {
	// user := User{}
	var users []User

	initD.DB.Where("last_name =?", "Woo").Find(&users)
	fmt.Println(users)
}

// Fetch Record using custume SQL statement
func custumeSQl() {
	user := User{}
	initD.DB.Raw("SELECT id, first_name, last_name from users where id = ?", 7).Scan(&user)
	fmt.Println(user)
}
