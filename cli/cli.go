package cli

import (
	"flag"
	"fmt"
	"os"
	"test/config"
	"test/swagger"
)

func usage() {
	fmt.Printf("###### Welcome WKMS Alert ######\n")	
	fmt.Printf("Please use the floowing flags:\n")
	fmt.Printf("-port: Set the port of the swagger Server\n")
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 8080, "set Port of the server")
	flag.Parse()
	config.Properties()
	swagger.SwaggerStart(*port)
	fmt.Println(*port)
}