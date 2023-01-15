package main

import (
    "clementdecou/messagebroker/broker"
)

type Message struct {
    value string
}

func main() {
    broker.InitServer()
}