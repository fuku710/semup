package semup

import (
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func ListVersions(r *git.Repository) ([]*semver.Version, error) {
	tIter, err := r.Tags()
	if err != nil {
		return nil, err
	}

	vs := []*semver.Version{}
	err = tIter.ForEach(func(t *plumbing.Reference) error {
		v, err := semver.NewVersion(t.Name().Short())
		if err == nil {
			vs = append(vs, v)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(semver.Collection(vs)))
	return vs, nil
}
