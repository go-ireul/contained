package contained

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const secretsDir = "/var/run/secrets"

// LoadSecrets load docker secrets from /var/run/secrets to os.Environ
func LoadSecrets() (err error) {
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(secretsDir); err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var data []byte
		if data, err = ioutil.ReadFile(filepath.Join(secretsDir, file.Name())); err != nil {
			return
		}
		os.Setenv(file.Name(), string(data))
	}
	return
}
