

```go
package main

import (
	"bufio"
	"io"
	"net"
	"os/exec"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:12345")
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
