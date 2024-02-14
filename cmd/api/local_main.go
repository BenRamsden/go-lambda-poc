//go:build !lambda

package main

import (
	"github.com/jugo-io/go-poc/internal/api"
	"github.com/jugo-io/go-poc/internal/api/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/go-poc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repo := sql.NewSQLRepository(db)
	if err := repo.Migrate(); err != nil {
		panic(err)
	}

	options := api.HandlerOptions{
		Repo: repo,
	}

	r := api.Handler(options)
	r.Run(":8080")
}
