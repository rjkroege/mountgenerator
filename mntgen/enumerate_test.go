package mntgen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"reflect"
)

func makeTestPaths(td string, fns []string)  []string {
	abspaths := []string{}
	for _, f := 	range fns {
		p := filepath.Join(td, f)
		abspaths = append(abspaths, p)
	}
	return abspaths	
}

// makeFiles creates a list of test files
func makeFiles(fns []string) error {
	for _, f := 	range fns {
		fd, err := os.Create(f)
		if err != nil {
			return err
		}
		fd.Close()
	}
	return nil	
}

var GeneratedTestFileNames = []string{
	"google-brick2",
	"google-persistent-disk-0",
	"scsi-0Google_PersistentDisk_brick2",
}

// Note the inclusion only of non-scsi names and the default
// names (starting with google-persistent) are excluded.
var ReturnedTestFileNames = []string{
	"google-brick2",
}

func Test_EnumerateFail(t *testing.T) {
	// Create a fake directory hierarchy.
	tdir, err := ioutil.TempDir("", "mountgenerator_test")
	if err != nil {
		t.Fatal("can't make temp dir for test", err)
	}
	defer os.RemoveAll(tdir)

	t.Log(tdir)

	// I don't understand how Glob works. Or reflect? Nope. glob.
	nopath := filepath.Join(tdir, "no-such-file")
	disks, err := EnumerateDisks(nopath)
	
	if len(disks) != 0 {
		t.Errorf("returned non-empty list for missing dir")
	}
}

func Test_Enumerate(t *testing.T) {
	// Create a fake directory hierarchy.
	tdir, err := ioutil.TempDir("", "mountgenerator_test")
	if err != nil {
		t.Fatal("can't make temp dir for test", err)
	}
	defer os.RemoveAll(tdir)

	if err := makeFiles(makeTestPaths(tdir, GeneratedTestFileNames)); err != nil {
		t.Fatal("Can't make test files", err)
	}

	disks, err := EnumerateDisks(tdir)
	if err != nil {
		t.Errorf("should not have received error", err)
	}
	if got, want := disks,  makeTestPaths(tdir, ReturnedTestFileNames); !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got: %#v", want, got)
	}
}

