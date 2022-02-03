//go:build integration
// +build integration

package build_tags

import "testing"

func Test_BuildTags(t *testing.T) {
	t.Run("Should only execute when run with tag supplied on the command line", func(t *testing.T) {
		// https://cs.opensource.google/go/go/+/go1.17.6:src/go/build/build.go;l=34

		// place tag (+build integration)  at top of file
		// needs to be followed by an empty line
		// can run on command line using:
		// go test tags=integration,otherExampleTag -v ./...
		// can also use negation eg. +build !integration

		t.Error("integration build tag supplied, failing test intentionally")
	})
}
