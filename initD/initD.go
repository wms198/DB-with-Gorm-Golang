package initD

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "root:@tcp(localhost)/gotest?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(" Could not connect mysql DB ")
	}
	//fmt.Println("DB is connected")

}

// package initD

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB
// var dsn = "root:@tcp(localhost)/gotest?charset=utf8mb4&parseTime=True&loc=Local"

// func ConnectDB() {
// 	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err)
// 		// panic("Could not connect mysql DB")
// 	}
// 	fmt.Print(DB)

// }
