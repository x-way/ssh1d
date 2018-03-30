package main

import (
	"github.com/coreos/go-systemd/activation"
	"github.com/coreos/go-systemd/daemon"
	"log"
	"net"
)

func main() {
	listeners, err := activation.Listeners(true)
	if err != nil {
		log.Fatalf("cannot retrieve listeners: %s", err)
	}
	if len(listeners) < 1 {
		listener, err := net.Listen("tcp", ":22")
		if err != nil {
			log.Fatalf("cannot listen: %s", err)
		}
		listeners = append(listeners, listener)
	} else if len(listeners) > 1 {
		log.Fatalf("unexpected number of socket activation (%d != 1)", len(listeners))
	}
	daemon.SdNotify(false, "READY=1")
	for {
		conn, err := listeners[0].Accept()
		if err != nil {
			log.Fatalf("cannot accept: %s", err)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			conn.Write([]byte("SSH-1.99-ssh1d\n"))
			log.Printf("Connection from %s\n", conn.RemoteAddr())
		}(conn)
	}
}
