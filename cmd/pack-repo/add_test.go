package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Azure/draft/pkg/draft/draftpath"
	"github.com/Azure/draft/pkg/draft/pack/repo"
)

func TestAddReturnsPackRepoExistsErr(t *testing.T) {
	tDir, teardown := tempDir(t)
	defer teardown()

	draftHome := draftpath.Home(tDir)
	if err := os.Mkdir(draftHome.Packs(), 0755); err != nil {
		t.Fatal(err)
	}

	create := &addCmd{
		source: "testdata/packrepo",
		out:    ioutil.Discard,
		err:    ioutil.Discard,
		home:   draftHome,
	}

	if err := create.run(); err != nil {
		t.Errorf("pack-repo add testdata/packrepo should not have errored the first time. Got error '%v'", err)
	}

	// run it a second time, expecting there to be an error
	if err := create.run(); err != repo.ErrExists {
		t.Errorf("pack-repo add testdata/packrepo != repo.ErrExists; got '%v'", err)
	}
}

// tempDir creates and returns the path as well as a function to clean the temporary directory
func tempDir(t *testing.T) (string, func()) {
	t.Helper()
	path, err := ioutil.TempDir("", "pack-repo-test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	return path, func() {
		if err := os.RemoveAll(path); err != nil {
			t.Fatalf("err: %s", err)
		}
	}
}
