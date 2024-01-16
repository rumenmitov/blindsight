package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

const log_path = "logs";

type Node struct {
    log string;
    next *Node;
}

type Queue struct {
    len int;
    head *Node;
    tail *Node;
}

func enqueue(queue *Queue, mes string) {
    var newLog *Node = &Node {
        log: mes,
        next: nil,
    }

    if queue.len == 0 {
        queue.head = newLog;
    } else {
        queue.tail.next = newLog;
    }

    queue.tail = newLog;
    queue.len++;
}

func dequeue(queue *Queue) (string, error) {
    if queue.len < 1 {
        return "", errors.New("Queue is empty!\n");
    }

    var log string = queue.head.log;

    if queue.len == 1 {
        queue.head = nil;
        queue.tail = nil;
    } else {
        queue.head = queue.head.next;
    }

    queue.len--;

    return log, nil;

}

func Log(mes string) {
    f_in, err := os.Open(log_path);
    if err != nil {
        os.Stderr.WriteString(err.Error());
    }

    defer f_in.Close();

    var queue Queue = Queue {
        len: 0,
        head: nil,
        tail: nil,
    };

    scanner := bufio.NewScanner(f_in);
    for scanner.Scan() {
        enqueue(&queue, scanner.Text());
    }

    f_out, err := os.OpenFile(log_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm);
    if err != nil {
       os.Stderr.WriteString(err.Error());
    }

    defer f_out.Close()

    message := fmt.Sprintf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), mes);

    f_out.WriteString(message);

    for {
        log, err := dequeue(&queue);
        if err != nil {
            break;
        }

        f_out.WriteString(log + "\n");
    }


}
