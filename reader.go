package board

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

type impossibleOpenFile struct {
	error
}

// ReadInput read the file with the inputs
func ReadInput(path string) ([]string, error) {
	var messages []string

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return messages, &impossibleOpenFile{errors.Wrapf(err, "Impossible open file")}
	}
	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = ';'
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		messages = append(messages, line[0])
	}

	return messages, nil
}

func IsImpossibleOpenFile(error error) bool {
	err := errors.Cause(error)

	_, ok := err.(*impossibleOpenFile)

	return ok
}
