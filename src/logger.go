package main

import (
	"bufio"
	"errors"
	"os"
)

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
    f_in, err := os.Open("../logs");
    if err != nil {
        os.Stderr.WriteString("Couldn't open logging file.\n");
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

    f_out, err := os.OpenFile("../logs", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm);
    if err != nil {
       os.Stderr.WriteString("Couldn't open logging file.\n");
    }

    defer f_out.Close()

    f_out.WriteString(mes + "\n");

    for {
        log, err := dequeue(&queue);
        if err != nil {
            break;
        }

        f_out.WriteString(log + "\n");
    }


}
