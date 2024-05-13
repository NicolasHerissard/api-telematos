package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	dbuser = flag.String("u", "", "Utilisateur de la base de donnée")
	dbpwd = flag.String("p", "", "Mot de l'utilisateur de la base de donnée")
	dbname = flag.String("d", "", "Nom de la base de donnée")
)

func main() {
	flag.Parse()

	if *dbuser == "" || *dbpwd == "" || *dbname == "" {
		log.Fatal("Please specify a user, password and name of database")
	}

	res := gin.Default()
	res.GET("/users", users)
	res.GET("/products", products)
	res.GET("/report", report)
	res.GET("/productsUsers", productsUsers)

	res.Run()
}

func users(c *gin.Context) {
	db := connectDB(*dbuser, *dbpwd, *dbname)

	user := listUsers(db)

	c.IndentedJSON(200, user)
}

func products(c *gin.Context) {
	db := connectDB(*dbuser, *dbpwd, *dbname)

    product := listProducts(db)

    c.IndentedJSON(200, product)
}

func report(c *gin.Context) {
	db := connectDB(*dbuser, *dbpwd, *dbname)

	report := listReport(db)

	c.IndentedJSON(200, report)
}

func productsUsers(c *gin.Context) {
	db := connectDB(*dbuser, *dbpwd, *dbname)

    productsUsers := listProductsUsers(db)

    c.IndentedJSON(200, productsUsers)
}