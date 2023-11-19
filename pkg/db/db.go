package db

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RaniaMidaoui/goMart-product-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{db}
}

func Mock() Handler {
	mockDb, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       mockDb,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mock.ExpectQuery(`SELECT`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "stock", "price"}).
			AddRow(1, "Prod A", 20, 15))

	mock.ExpectQuery(`SELECT`).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "stock", "price"}).
			AddRow(2, "Prod B", 10, 20))

	return Handler{db}
}
