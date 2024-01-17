package main

type DB_Config struct {
    host string;
    port uint16;
    user string;
    password string;
    db_name string;
};

type User struct {
    Fname string    `json:"fname"`;
    Lname string    `json:"lname"`;
    Email string    `json:"email"`;
    Username string `json:"username"`;
    Password string `json:"password"`;
};

type ImageFile struct {
    name string;
    bytes string;
};

type AuthError uint
const (
    Ok AuthError = iota + 201;
    NotANumberError;
    WrongCredentialsError;
    UnkownError;
)
