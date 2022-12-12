
## TCP Server
```go
package main

import (
	"bufio"
	"io"
	"net"
	"os/exec"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:8080")
	for {
		c, _ := l.Accept()
		go func() {
			handle(&c)
		}()
	}

}

func handle(c *net.Conn) {
	cmd := exec.Command("top", "-b")
	cstdout, _ := cmd.StdoutPipe()
	go io.Copy(bufio.NewWriter(*c), bufio.NewReader(cstdout))
	cmd.Run()
}
```

## TCP Client
```go
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096) // big buffer
	var c int
	for {
		c++
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		fmt.Println(c, "Received message:", n, string(buf))
	}
	fmt.Println(string(buf))
}
```
## TCP Server | Publish to web
```go
func tcpToweb() {

	l, _ := net.Listen("tcp", "127.0.0.1:12345")
	defer l.Close()

	for {
		// Wait for a connection.
		conn, _ := l.Accept()

		go func(c net.Conn) {
			//for {

			// msg, err := bufio.NewReader(conn).ReadString('\n')
			// if err != nil {
			// 	break
			// }
			// program := strings.Trim(msg, "\r\n")

			cmd := exec.Command("top", "-bn1")
			var b bytes.Buffer
			cmd.Stdout = &b
			cmd.Run()
			//conn.Write(b.Bytes())
			body := b.String()
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)

			//}
			//connection closed
			c.Close()
		}(conn)
	}
}
```
