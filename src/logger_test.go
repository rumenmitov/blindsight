package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestLoggerWithSingleInput(t *testing.T) {
    var log_1 string = "hello";

    Log(log_1);

    log_file, err := os.Open("../logs");
    if err != nil {
        t.Error("No logging file was created.");
    }

    scanner := bufio.NewScanner(log_file);
    for scanner.Scan() {
        if scanner.Text() != log_1 {
            err := fmt.Sprintf("First log is incorrect! Entry was: %s, logged value was: %s",
                log_1 + "\n", scanner.Text());
            
            t.Error(err);
        }
    }

    log_file.Close();
    if err := os.Remove("../logs"); err != nil {
        os.Stderr.WriteString("Failed to remove test log file!\n");
    }
}

func TestLoggerWithMultipleInputs(t *testing.T) {
    var test_cases = [3]string{ "test case 1", "test case 2", "test case 3" };

    i := 0;

    for i < len(test_cases) {
        Log(test_cases[i]);
        i++;
    }

    log_file, err := os.Open("../logs");
    if err != nil {
        t.Error("No logging file was created.");
    }

    i = 0;

    scanner := bufio.NewScanner(log_file);
    for scanner.Scan() {
        if scanner.Text() != test_cases[2 - i] {
            err := fmt.Sprintf("First log is incorrect! Entry was: %s, logged value was: %s",
                test_cases[2 - i] + "\n", scanner.Text());
            
            t.Error(err);
        }

        i++;
    }

    log_file.Close();
    if err := os.Remove("../logs"); err != nil {
        os.Stderr.WriteString("Failed to remove test log file!\n");
    }
}
