package monitor

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
)

var c *statsd.Client

func SetUpStatsD() *statsd.Client {
	if c == nil {
		c, err := statsd.New() // Connect to the UDP port 8125 by default.
		if err != nil {
			// If nothing is listening on the target port, an error is returned and
			// the returned client does nothing but is still usable. So we can
			// just log the error and go on.
			log.Error(err)
		}
		defer c.Close()
		log.Info("statsd starts")
	}
	return c
}
