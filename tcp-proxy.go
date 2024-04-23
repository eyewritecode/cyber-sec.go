// Typical usage of this would be to bypass restriction on a website X through a network Y.
// Suppose you want to access twitter on your office network but you can't because they've blocked direct access to it
// A proxy server like this stands between you, your work network and the prohibited site.
// You would have to run it and configure your browser proxy settings to this(i.e 127.0.0.1:80)

package main

import (
	"io"
	"log"
	"net"
)

func bypassWorkNetwork(src net.Conn, restricedSite string) {
	dst, err := net.Dial("tcp", restricedSite)

	if err != nil {
		log.Fatalln("Unable to connect to reach website")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//Replace this with your proxy server url(i.e mytwitter.com:80)
	listener, err := net.Listen("tcp", ":80")

	if err != nil {
		log.Fatalln("Failed to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Connection failed")
		}
		go bypassWorkNetwork(conn, "scanme.nmap.org:80")
	}
}
