package auto

import (
	"ireul.com/contained"
	"os"
)

func init() {
	if err := contained.LoadSecrets(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}
