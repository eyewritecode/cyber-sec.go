package main

import (
	"fmt"
	"net"
	"sync"
)

func scanOpenPort(address string, port int) {
	target := fmt.Sprintf("%s:%d", address, port)
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println(err)
	}else{
		conn.Close();
		fmt.Printf("Port: %d is available\n", port)
	}
}

func main(){
	var wg sync.WaitGroup
	for port := 1; port<= 1024; port++ {
		wg.Add(1)
		defer wg.Done()
		go scanOpenPort("127.0.0.1", port)

	}
	wg.Wait()
}
