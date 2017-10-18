package main

import (
	"log"

	"github.com/rjkroege/mountgenerator"
)

func main() {
	log.Println("hello")

	disks, err := mountgenerator.EnumerateDisks("/dev/disk/by-id")	
	if err != nil {
		// TODO(rjk): make sure that mountgenerator fails end up in the system
		// log. Maybe they will by default.
		log.Fatalln("mountgenerator had a sad", err)
	}



}