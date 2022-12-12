
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

## Tcp Client
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
