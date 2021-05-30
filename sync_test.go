package bssync

import "testing"

func TestSyncFilesWithCheck(t *testing.T) {
	if err := SyncFiles(SyncCheck, "a.sqlite", "b.sqlite"); err != nil {
		t.Error(err)
	}
}
