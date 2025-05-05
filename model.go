package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID string
	Name string
	Email string
	Password string
	Role int
	Created_at time.Time
}

type Product struct {
	ID string
	Name_Product string
	stock int
}

type Report struct {
	ID string
	UserID string
	ProductID string
	Description string
	Created_at time.Time
}

type ProductsUsers struct {
	ID string
	UserID string
	ProductID string
	Take_Product int
}

func connectDB(user string, pwd string, dbname string) (*gorm.DB) {
	stringConnection := user+":"+pwd+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(stringConnection))
	if err != nil {
		fmt.Println("Error connecting to database : ", err.Error())
	}
	return db
}

func listUsers(db *gorm.DB) []User {
	var users []User

	db.Find(&users)

	var user []User
	user = append(user, users...) 

	return user
}

func listProducts(db *gorm.DB) []Product {
	var p []Product
	db.Find(&p)

	var product []Product
	product = append(product, p...)

	return product
}

func listReport(db *gorm.DB) []Report {
	var r []Report
	db.Find(&r)

	var report []Report
	report = append(report, r...)

	return report
}

func listProductsUsers(db *gorm.DB) []ProductsUsers {
	var pu []ProductsUsers
    db.Find(&pu)

    var productsUsers []ProductsUsers
    productsUsers = append(productsUsers, pu...)

    return productsUsers
}

func createUsers(db *gorm.DB, name string, email string, pwd string, role int) int64 {

	users := []User {
		{Name: name, Email: email, Password: pwd, Role: role, Created_at: time.Now()},
	}

	result := db.Create(&users)

	if result.Error != nil {
		fmt.Println(result.Error)
        return 0
	}

	return result.RowsAffected
}