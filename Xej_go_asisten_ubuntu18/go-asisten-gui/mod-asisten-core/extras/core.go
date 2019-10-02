package extras

import "github.com/otiai10/copy"

// DataPrepare prepare data
func DataPrepare(DataDir string, PreDataDir string) (err error) {
	return copy.Copy(PreDataDir, DataDir);
}

