package cli

import (
	"flag"
	"fmt"
	"os"
	"test/config"
	"test/swagger"
)

func usage() {
	fmt.Printf("================================= [START] WKMS Alert =================================\n\n")
	fmt.Printf("\t해당 플래그로 Port 와 환경설정 파일의 경로를 지정할 수 있습니다.\n")
	fmt.Printf("\t\t-port: 스웨거 기동 port\n")
	fmt.Printf("\t\t-yaml: 환경설정 파일 \n\n")
	fmt.Printf("=============================== Licenced by WMP CERT =================================\n")
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 8080, "set Port of the server")
	yaml := flag.String("yaml", "properties.yaml", "set Yaml Dir of the server")
	flag.Parse()
	config.Init(*yaml)
	swagger.Start(*port)
}
