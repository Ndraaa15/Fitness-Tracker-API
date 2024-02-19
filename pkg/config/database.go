package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func InitializeDatabaseConn(v *viper.Viper) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		v.GetString("DB.HOST"),
		v.GetString("DB.PORT"),
		v.GetString("DB.USER"),
		v.GetString("DB.PASSWORD"),
		v.GetString("DB.NAME"))

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
