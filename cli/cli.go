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
	fmt.Printf("-yaml: Set the yaml file dir \n")
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 8080, "set Port of the server")
	yaml := flag.String("yaml", "properties.yaml", "set Yaml Dir of the server")
	flag.Parse()
	config.Init(*yaml)
	swagger.SwaggerStart(*port)
}