// +build !ios,!android

package gomobile

import (
	"errors"
	"io"
	"os"
)

func nativeFileOpen(f *file) (io.ReadCloser, error) {
	if len(f.uri) < 8 || f.uri[:7] != "file://" {
		return nil, errors.New("Mobile simulator mode only supports file:// URIs")
	}

	return os.Open(f.uri[7:])
}

func nativeFileSave(f *file) (io.WriteCloser, error) {
	if len(f.uri) < 8 || f.uri[:7] != "file://" {
		return nil, errors.New("Mobile simulator mode only supports file:// URIs")
	}

	return os.Open(f.uri[7:])
}
