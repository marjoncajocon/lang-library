//////////// server 

package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen on local port 8081 using UDP
    addr, err := net.ResolveUDPAddr("udp", ":8081")
    if err != nil {
        fmt.Println("Failed to resolve address:", err)
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Failed to listen on UDP:", err)
        return
    }
    defer conn.Close()
    fmt.Println("UDP server is listening on port 8081...")

    buffer := make([]byte, 1024)
    for {
        // Read data sent by client
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading data:", err)
            continue
        }
        message := string(buffer[:n])
        fmt.Printf("Received message from %s: %s", clientAddr, message)

        // Echo the message back to the client
        _, err = conn.WriteToUDP([]byte("Echo: "+message), clientAddr)
        if err != nil {
            fmt.Println("Failed to send message:", err)
        }
    }
}

/////////// client

package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Resolve server address
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8081")
    if err != nil {
        fmt.Println("Failed to resolve server address:", err)
        return
    }

    // Create a UDP connection (actually connectionless)
    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("Failed to connect to UDP server:", err)
        return
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)
    for {
        // Read user input from standard input
        fmt.Print("Enter message: ")
        message, _ := reader.ReadString('\n')
        // Send message to the server
        _, err := conn.Write([]byte(message))
        if err != nil {
            fmt.Println("Failed to send message:", err)
            return
        }
        // Receive echoed message from the server
        buffer := make([]byte, 1024)
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Failed to receive message:", err)
            return
        }
        fmt.Print("Server echo: " + string(buffer[:n]))
    }
}
