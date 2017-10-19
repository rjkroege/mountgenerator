package mntgen

import (
	"io/ioutil"
	"os"
	"testing"
)

var RelativeDeviceNames = []string{
	"google-brick2",
}

const CorrectUnit = `[Unit]
    Description=Mount brick2

[Mount]
What=/dev/disk/google-brick2
Where=/mnt/disks/brick2
Type=ext4
Options=defaults

[Install]
WantedBy=multi-user.target
`

func Test_ForAllDisks(t *testing.T) {
	testdevices :=  makeTestPaths("/dev/disk", RelativeDeviceNames)

	// Create a fake directory hierarchy.
	tdir, err := ioutil.TempDir("", "mountgenerator_test")
	if err != nil {
		t.Fatal("can't make temp dir for test", err)
	}
	defer os.RemoveAll(tdir)

	ForAllDisks(testdevices, tdir)

	
	if got, want := Diskname(testdevices[0]), "brick2"; got != want {
		t.Fatal("Something wrong with Diskname, got", got, "want", want)
	}

	bslice, err := ioutil.ReadFile(ServicedName(tdir, Diskname(testdevices[0])))
	if err != nil {
		t.Fatal("couldn't open the file that I was building", err)
	}

	if got, want := string(bslice), CorrectUnit; got != want {
		t.Fatalf("got %#v but want %#v", got, want)
	}
}