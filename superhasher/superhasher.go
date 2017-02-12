package superhasher

import (
	log "github.com/Sirupsen/logrus"
)

func Superhash(path string) string {
	log.Printf("superhash called: %s", path)
	ret := "superhasher"
	return ret
}
