package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/debarshibasak/kubestrike/v1alpha1/config"
)

func main() {

	configuration := flag.String("config", "", "location of configuration")
	run := flag.Bool("run", false, "install operation")
	strictInstalltion := flag.Bool("use-strict", false, "uninstall operation")
	verbose := flag.Bool("verbose", false, "uninstall operation")

	flag.Parse()

	log.Println("[kubestrike] started")

	if *run {
		configRaw, err := ioutil.ReadFile(*configuration)
		if err != nil {
			log.Fatal(err)
		}

		clusterOperation, err := config.NewParser(*strictInstalltion).Parse(configRaw)
		if err != nil {
			log.Fatal(err)
		}

		if err := clusterOperation.Validate(); err != nil {
			log.Fatal(err)
		}

		if err := clusterOperation.Run(*verbose); err != nil {
			log.Fatal(err)
		}
	}
}
