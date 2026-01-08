package database

import (
	"log"
	"mini-indobat/models"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(gormPostgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error on connecting to the database %v", err.Error())
	}

	log.Println("Successfully connected to the database")
	return db
}

func GetTestDatabaseConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(gormPostgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error on connecting to the test database %v", err.Error())
	}

	db.Exec("TRUNCATE TABLE orders RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE products RESTART IDENTITY CASCADE")

	stok := 1
	db.Create(&models.Products{
		Id:       1,
		NamaObat: "Paracetamol",
		Stok:     &stok,
		Harga:    2000,
	})

	log.Println("Successfully connected to the test database")
	return db
}

func RunMigrations(dsn string) {
	m, err := migrate.New(
		"file://utils/database/migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("Error on migrating %v", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Print(err.Error())
	}

	log.Println("Migration successfull")
}
