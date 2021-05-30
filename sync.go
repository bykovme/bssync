package bssync

const (
	SyncCheck   = "check"
	SyncMerge   = "merge"
	SyncReplace = "replace"
)

// SyncFiles - syncs the files, no password required
func SyncFiles(syncType string, src string, dest ...string) error {
	switch syncType {
	case SyncCheck:
		return checkFiles(src, dest)
	case SyncMerge:
		if err := checkFiles(src, dest); err != nil {
			return err
		}
		return mergeFiles(src, dest)
	case SyncReplace:
		if err := checkFiles(src, dest); err != nil {
			return err
		}
		return replaceFiles(src, dest)
	}
	return nil
}
