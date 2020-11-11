package monitor

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SetUpLog() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	//set up logrus
	filename := "/var/log/webapp.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(f)

	//log.SetLevel(log.WarnLevel)
}
