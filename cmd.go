package main

import (
	"log"

	"github.com/rjkroege/mountgenerator/mntgen"
)

func main() {
	disks, err := mntgen.EnumerateDisks("/dev/disk/by-id")	
	if err != nil {
		// TODO(rjk): make sure that mountgenerator fails end up in the system
		// log. Maybe they will by default.
		log.Fatalln("mountgenerator had a sad", err)
	}
	mntgen.ForAllDisks(disks, mntgen.ServicedUnitsLocation)
}
