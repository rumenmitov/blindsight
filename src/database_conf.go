package main

import (
	"database/sql"
	"fmt"
	"os"
    "errors"

	"github.com/joho/godotenv"
)

func configureDatabase() (*sql.DB, error) {
    if err := godotenv.Load(); err != nil {
        return nil, err;
    }

    db_config := DB_Config {
        host: "localhost",
        port: 5432,
        user: os.Getenv("DB_USER"),
        password: os.Getenv("DB_PASS"),
        db_name: "Users",
    };

    psqlconn := 
        fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        db_config.host, db_config.port, db_config.user, db_config.password, db_config.db_name);

    db, err := sql.Open("postgres", psqlconn);
    if err != nil {
        return nil, errors.New("Couldn't open database!\n");
    }

    return db, nil;

}
