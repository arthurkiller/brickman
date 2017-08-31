package main

import (
	"io"
	"net"
	"os"
	"sync"
)

var (
	ServerHost string
	ClientHost string
	BLOCK      int = 128 * 1024 * 1024
)

func openFile(path string) *os.File {
	fd, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return fd
}

func serverSend() {}

func server() net.PacketConn {
	net.ResolveUDPAddr("udp", ServerHost)
	packconn, err := net.ListenPacket("udp", ClientHost)
	if err != nil {
		panic(err)
	}
	return packconn
}

func main() {
	var err error
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024*1024*1024+BLOCK, 1024*1024*1024+BLOCK)
	var i, n int
	var wg sync.WaitGroup
	//var handler index

	handler := func(buffer []byte) {
		defer wg.Done()
		// TODO split the buffer with \n
		for k, _ := range buffer {
			_ = k
		}
	}

	processor := func(buffer []byte) {
		// TODO process every block
		buf = make([]byte, 1472, 1472)
		buf = append(buffer[:len(buffer)], buffer[len(buffer)+len(buffer):]...)
		_ = buf
	}

	_ = processor

	for {
		n, err = file.Read(buf[i : i+BLOCK])
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					wg.Add(1)
					go handler(buf[i : i+n])
					i += n
				}
				break
			} else {
				panic(err)
			}
		}
		if n > 0 {
			wg.Add(1)
			go handler(buf[i : i+n])
			i += n
		}
	}
}
