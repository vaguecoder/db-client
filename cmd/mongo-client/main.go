package main

import (
	"fmt"
	"log"

	"github.com/vaguecoder/db-client/pkg/mongowire"
)

func main() {
	// uri := "mongodb://mongoadmin:password123@127.0.0.1:27018/admin"
	conn, err := mongowire.ConnectTCP("127.0.0.1:27018")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", conn)
}
