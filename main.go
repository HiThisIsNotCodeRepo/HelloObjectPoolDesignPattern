package main

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	aPool, err := initPool(connections)
	if err != nil {
		log.Fatalf("init pool error: %s", err)
	}
	conn1, err := aPool.loan()
	if err != nil {
		log.Fatalf("pool loan error: %s", err)
	}
	conn2, err := aPool.loan()
	if err != nil {
		log.Fatalf("pool load error: %s", err)
	}
	aPool.receive(conn1)
	aPool.receive(conn2)
}
