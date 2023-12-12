package main

import (
	"flag"
	"fmt"
	"github.com/talkkonnect/capturecam"
)

func main() {

	config := flag.String("config", "/home/talkkonnect/gocode/src/github.com/talkkonnect/capturecam/capturecam.xml", "full path to capturecam.xml configuration file")

	flag.Usage = capturecamusage
	flag.Parse()

	capturecam.Start(config)

}

func capturecamusage() {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("Usage: talkkonnect [-config=[full path and file to talkkonnect.xml configuration file]]")
	fmt.Println("By Suvir Kumar <suvir@talkkonnect.com>")
	fmt.Println("For more information visit http://www.talkkonnect.com and github.com/talkkonnect")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("-config=/home/talkkonnect/gocode/src/github.com/talkkonnect/talkkonnect/talkkonnect.xml")
	fmt.Println("-serverindex=[n] for the index of the enabled server to connect to in XML file")
	fmt.Println("-version for the version")
	fmt.Println("-help for this screen")
}
