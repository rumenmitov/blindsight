package main

import (
	"database/sql"
	"fmt"
	"os"
    "errors"
    "strconv"

	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func configureDatabase() (*sql.DB, error) {
    if err := godotenv.Load(); err != nil {
        return nil, err;
    }

    db_port, err := strconv.Atoi(os.Getenv("DB_PORT"));
    if err != nil {
        return nil, err;
    }

    db_config := DB_Config {
        host: "psql",
        port: uint16(db_port),
        user: os.Getenv("DB_USER"),
        password: os.Getenv("DB_PASS"),
        db_name: os.Getenv("DB_NAME"),
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
