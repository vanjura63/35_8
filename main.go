package main

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Сетевой адрес.
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

var proverb = []string{

	"Don't communicate by sharing memory, share memory by communicating",
	"Concurrency is not parallelism",
	"Channels orchestrate; mutexes serialize",
	"The bigger the interface, the weaker the abstraction",
	"Make the zero value useful",
	"interface{} says nothing",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite",
	"A little copying is better than a little dependency",
	"Syscall must always be guarded with build tags",
	"Cgo must always be guarded with build tags",
	"Cgo is not Go",
	"With the unsafe package there are no guarantees",
	"Clear is better than clever",
	"Reflection is never clear",
	"Errors are values",
	"Don't just check errors handle them gracefully",
	"Design the architecture, name the components, document the details",
	"Documentation is for users",
	"Don't panic",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	// Запуск сетевой службы
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	var list func() = func() {

		for {
			// Принимаем подключение.
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			// Вызов обработчика подключения.

			handleConn(conn)
		}
	}
	for i := 1; i <= 20; i++ {
		go list()
	}
	wg.Add(1)
	wg.Wait()
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {
	tick := time.NewTicker(time.Second * 3)
	for {
		select {
		case _ = <-tick.C:
			rand := rand.Intn(18)
			conn.Write([]byte(proverb[rand] + "\n"))

		}

	}

}
