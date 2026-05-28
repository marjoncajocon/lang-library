///////////////// server
package main

import (
    "bufio"
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        // Read data sent by client
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading data:", err)
            break
        }
        fmt.Printf("Received message: %s", message)
        // Echo the message back to the client
        conn.Write([]byte("Echo: " + message))
    }
}

func main() {
    // Listen on local port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Failed to listen on port:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening on port 8080...")

    for {
        // Accept client connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Failed to accept connection:", err)
            continue
        }
        // Handle the connection (can handle multiple concurrently)
        go handleConnection(conn)
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
    // Connect to the server (localhost:8080)
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to connect to server:", err)
        return
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)
    for {
        // Read user input from standard input
        fmt.Print("Enter message: ")
        message, _ := reader.ReadString('\n')
        // Send the message to the server
        _, err := conn.Write([]byte(message))
        if err != nil {
            fmt.Println("Failed to send message:", err)
            return
        }
        // Receive echoed message from the server
        response, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Failed to receive message:", err)
            return
        }
        fmt.Print("Server echo: " + response)
    }
}
