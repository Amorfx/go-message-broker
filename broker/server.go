package broker

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)


func InitServer() {
    argument := os.Args
    if len(argument) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }

    // Create the server
    port := ":" + argument[1]
    server, err := net.Listen("tcp4", port)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer server.Close()

    // messageChannel := make(chan Message)

    // can handle push or send event to a worker
    // go handleMessage(messageChannel)


    // Client connection
    fmt.Println("Start listening to " + port)
    for {
        client, err := server.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        go handleConnection(client)
    }
}

func handleConnection(client net.Conn) {
    defer client.Close()

    fmt.Printf("Serving %s\n", client.RemoteAddr().String())

    for {
        netData, err := bufio.NewReader(client).ReadString('\n')
        if err != nil {
            fmt.Println("error reading:", err)
            break
        }

        commandString := strings.TrimSpace(netData)
        parts := strings.Split(commandString, " ")
        command := parts[0]
        switch command {
        case "POST":
            fmt.Println("OK POST")
        }
        client.Write([]byte(command + "\n"))
    }
}