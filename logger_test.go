package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoggerWithSingleInput(t *testing.T) {
    os.Remove("logs");

    var log_1 string = "hello";

    Log(log_1);

    log_file, err := os.Open(log_path);
    if err != nil {
        t.Error("No logging file was created.");
    }

    scanner := bufio.NewScanner(log_file);
    for scanner.Scan() {
        log := strings.Split(scanner.Text(), " - ");

        if log[len(log) - 1] != log_1 {
            err := fmt.Sprintf("First log is incorrect! Entry was: %s, logged value was: %s",
                log_1 + "\n", log[len(log) - 1]);
            
            t.Error(err);
        }
    }

    log_file.Close();
    if err := os.Remove(log_path); err != nil {
        os.Stderr.WriteString("Failed to remove test log file!\n");
    }
}

func TestLoggerWithMultipleInputs(t *testing.T) {
    os.Remove("logs");

    var test_cases = [3]string{ "test case 1", "test case 2", "test case 3" };

    i := 0;

    for i < len(test_cases) {
        Log(test_cases[i]);
        i++;
    }

    log_file, err := os.Open(log_path);
    if err != nil {
        t.Error("No logging file was created.");
    }

    i = 0;

    scanner := bufio.NewScanner(log_file);
    for scanner.Scan() {

        log := strings.Split(scanner.Text(), " - ");

        if log[len(log) - 1] != test_cases[2 - i] {
            err := fmt.Sprintf("First log is incorrect! Entry was: %s, logged value was: %s",
                test_cases[2 - i] + "\n", log[len(log) - 1]);
            
            t.Error(err);
        }

        i++;
    }

    log_file.Close();
    if err := os.Remove(log_path); err != nil {
        os.Stderr.WriteString("Failed to remove test log file!\n");
    }
}
