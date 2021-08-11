package file

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"path"
)

func WriteFile(dir, name string, data bytes.Buffer) error {
	return ioutil.WriteFile(path.Join(dir, name+".go"), data.Bytes(), fs.ModePerm)
}
