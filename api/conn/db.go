package conn

import (
	"api/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBPgsql *gorm.DB

func ConnectionPgsqlDB(db_host string, db_user string, db_password string, db_name string, db_port string, ssl_mode string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	DBPgsql = db

	err = db.AutoMigrate(&entity.User{}, &entity.Auth{}, &entity.Role{}, &entity.TShirt{}, entity.Order{}, entity.Cart{}, entity.Store{}, entity.Customer{}, entity.CustomerToken{})
	if err != nil {
		db.DisableForeignKeyConstraintWhenMigrating = true
		db.AutoMigrate(&entity.User{}, &entity.Auth{}, &entity.Role{}, &entity.TShirt{}, entity.Order{}, entity.Cart{}, entity.Store{}, entity.Customer{}, entity.CustomerToken{})
		db.DisableForeignKeyConstraintWhenMigrating = false
		db.AutoMigrate(&entity.User{}, &entity.Auth{}, &entity.Role{}, &entity.TShirt{}, entity.Order{}, entity.Cart{}, entity.Store{}, entity.Customer{}, entity.CustomerToken{})
	}
}
