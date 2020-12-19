package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/route"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
	"os"
)

func main() {
	log.Info("webapp starts...")

	//set up logrus
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	filename := "/var/log/webapp.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(f)

	//log.SetLevel(log.WarnLevel)

	//set up statsd
	client, err := statsd.New() // Connect to the UDP port 8125 by default.
	if err != nil {
		log.Error(err)
	}
	defer client.Close()

	//set up db
	if err = config.ConfigDB(); err != nil {
		log.Fatal(err)
	}

	log.Info("waiting for request...")
	r := route.SetupRouter(client)

	//running
	log.Info("webapp is running...")
	r.Run()
	log.Info("webapp ends...")
}
