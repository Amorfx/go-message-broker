package broker

import (
    "bufio"
    "encoding/json"
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

    mapChanConnection := map[string] net.Conn{}
    messageChannel := make(chan string)

    // can handle push or send event to a worker
    go handleMessage(messageChannel, mapChanConnection)


    // Client connection
    fmt.Println("Start listening to " + port)
    for {
        client, err := server.Accept()
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        go handleConnection(messageChannel, client, mapChanConnection)
    }
}

func handleMessage(messageChannel chan string, mapChanConnection map[string] net.Conn) {
    for {
        select {
        case message := <- messageChannel:
            // The message is a json with channel and message value
            mapMessage := map[string] string{}
            err := json.Unmarshal([]byte(message), &mapMessage)
            if err != nil {
                fmt.Println("Error decoding json message " + message + " err: " + err.Error())
                break
            }
            chanToSearchWorker := mapMessage["chan"]
            messageValue := mapMessage["value"]

            // TODO get one not only first
            if _, ok := mapChanConnection[chanToSearchWorker]; ok {
                jsonValue, err := json.Marshal(messageValue + "\n")
                if err != nil {
                    fmt.Println("Error when encode json for message value " + messageValue)
                    break
                }
                fmt.Println("Send to client message" + messageValue)
                mapChanConnection[chanToSearchWorker].Write(jsonValue)
            }
        }
    }
}

func handleConnection(messageChan chan string, client net.Conn, mapChanConnection map[string] net.Conn) {
    defer client.Close()

    fmt.Printf("Serving %s\n", client.RemoteAddr().String())

    for {
        netData, err := bufio.NewReader(client).ReadString('\n')
        if err != nil {
            fmt.Println("error reading:", err)
            break
        }

        commandString := strings.TrimSpace(netData)
        if strings.Contains(commandString, "CONNECT:") {
            channelToConnect := strings.Replace(commandString, "CONNECT:", "", 1)
            if (len(channelToConnect) > 0) {
                mapChanConnection[channelToConnect] = client
                fmt.Println("Add client to chan " + channelToConnect)
            }
        } else if strings.Contains(commandString, "SEND:") {
            messageToSend := strings.Replace(commandString, "SEND:", "", 1)
            messageChan <- messageToSend
        }
    }
}