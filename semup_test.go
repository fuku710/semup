package semup_test

import (
	"testing"
	"time"

	"github.com/fuku710/semup"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func TestListVersion(t *testing.T) {
	s := memory.NewStorage()
	f := memfs.New()

	r, err := git.Init(s, f)
	if err != nil {
		t.Error(err)
	}

	w, err := r.Worktree()
	if err != nil {
		t.Error(err)
	}

	h, err := w.Commit("test commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "text@example.com",
			When:  time.Now(),
		}})
	if err != nil {
		t.Error(err)
	}

	_, err = r.CreateTag("1.0.0", h, nil)
	if err != nil {
		t.Error(err)
	}

	vs, err := semup.ListVersions(r)
	if err != nil {
		t.Error(err)
	}

	if len(vs) < 1 {
		t.Errorf("versions are empty")
		return
	}

	if vs[0].String() != "1.0.0" {
		t.Errorf("vs[0] expected 1.0.0 but got %s", vs[0].String())
		return
	}
}
