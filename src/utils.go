package main

import (
	"encoding/base64"
	"errors"
	"os"
)

func saveImage(file_name string, file_bytes string) error {
    dec, err := base64.StdEncoding.DecodeString(file_bytes);
    if err != nil {
        return errors.New("Couldn't decode file!\n");
    }

    _, err = os.Stat("images/");
    if err != nil {
        if os.IsNotExist(err) {
            os.Mkdir("images", os.ModePerm);
        } else {
            return err;
        }
    }

    file, err := os.Create("images/" + file_name + ".png");
    if err != nil {
        return errors.New("Couldn't create file!\n");
    }

    defer file.Close();

    _, err = file.Write(dec);
    if err != nil {
        return errors.New("Couldn't write to file!\n");
    }

    err = file.Sync();
    if err != nil {
        return errors.New("Couldn't sync file to disk!\n");
    }

    return nil;
}
