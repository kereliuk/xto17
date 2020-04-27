package storage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Copy stolen from stackoverflow.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// PrintLineByLine takes an io.ReadCloser, rc, and reads it line by line
func PrintLineByLine(rc io.ReadCloser) {
	r := bufio.NewReader(rc)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Still need to figure out error handling here. Maybe ask Ivan?
			fmt.Printf("ERR: %v\n", err)
			break
		}
		fmt.Printf("%s\n", string(line))
	}
}

// GetPathReader is a helper function that opens a file. Honestly I am not sure why I made this function instead of just writing it out.
func GetPathReader(input string) (io.Reader, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// PathIterate takes a glob and then matches that to a list of files. Then it peferforms a callback cb on them iteratively.
func PathIterate(glob string, cb func(string) error) error {
	fileNames, err := filepath.Glob(glob)
	if err != nil {
		return errors.Wrap(err, "failed to read dir")
	}

	for _, f := range fileNames {
		fileInfo, err := os.Stat(f)
		if err != nil {
			return errors.Wrap(err, "failed to stat path")
		}

		if fileInfo.IsDir() {
			continue
		}

		if err := cb(f); err != nil {
			return errors.Wrap(err, "failed callback")
		}
	}

	return nil
}
