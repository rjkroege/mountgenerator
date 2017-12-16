package mntgen

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const mount_unit_template = `[Unit]
    Description=Mount {{.Diskname}}

[Mount]
What={{.Device}}
Where={{.MntPoint}}
Type=ext4
Options=defaults

[Install]
WantedBy=multi-user.target
`

type Mount struct {
	Diskname string
	MntPoint string
	Device   string
}

// I presume that the template is parseable.
var tpl = template.Must(template.New("unit").Parse(mount_unit_template))

// Unit generates the contents of the serviced to the provided writer for disk
func Unit(writer io.Writer, disk, device string) error {
	mnt := &Mount{
		Diskname: disk,
		Device:   device,
		MntPoint: MountPoint(disk),
	}

	if err := tpl.Execute(writer, mnt); err != nil {
		return err
	}
	return nil
}

// MountPoint generates the name of the mountpoint
func MountPoint(disk string) string {
	if disk == "home" {
		return "/home"
	}
	return filepath.Join("/mnt/disks", disk)
}

// Diskname generates the name of the given device path.
// The disk named homedir or home is treated specially and mounted on
// /home on the host system.
func Diskname(devpath string) string {
	bp := filepath.Base(devpath)
	tbp := strings.TrimPrefix(bp, "google-")
	if tbp == "home" || tbp == "homedir" {
		return "home"
	}
	return tbp
}

const ServicedUnitsLocation = "/etc/systemd/system"

// ServicedName generates the name for this serviced unit from the
// provided diskname.
func ServicedName(location, diskname string) string {
	if diskname == "home" {
		return filepath.Join(location, diskname+".mount")
	}
	return filepath.Join(location, "mnt-disks-"+diskname+".mount")
}

// ForAllDisks writes a serviced to the appropriate location for each device path
// in devices.
func ForAllDisks(devices []string, location string) {
	for _, d := range devices {
		fname := ServicedName(location, Diskname(d))
		fd, err := os.Create(fname)
		if err != nil {
			// TODO(rjk): make sure that these get logged correctly.
			log.Println("Couldn't create mount unit:", fname, "because", err)
			continue
		}
		if err := Unit(fd, Diskname(d), d); err != nil {
			log.Println("Couldn't write the unit", fname, "because", err)
		}
		fd.Close()
	}
}
