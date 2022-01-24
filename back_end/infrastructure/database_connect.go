package infrastructure

import (
	"log"
	"market/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	connectSQL := "host" + dbHost +
	" user=" + dbUser +
	" dbname=" + dbName +
	" password=" + dbPassword +
	" sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{})

	if err != nil {
		ErrLog.Printf("Has problem at connection database: %+v\n", err)
		return nil, err
	}
	return db, nil
}

func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		return err
	}

	if allowMigrate {
		log.Println("Migrating database...")
		db.AutoMigrate(
			&model.User{},
		)
	}

	return nil
}