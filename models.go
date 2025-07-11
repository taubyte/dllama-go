package dllama

import (
	"fmt"

	"github.com/taubyte/dllama-go/symbols"
)

func ListModels() ([]string, error) {
	listId := symbols.ListModels()

	if int64(listId) < 0 {
		return nil, fmt.Errorf("ListModels returned %d", listId)
	}

	buf := make([]byte, 1024)
	var models []string

	for {
		n := symbols.ListModelsNext(listId, &buf[0])
		if n < 0 {
			return nil, fmt.Errorf("ListModelsNext returned %d", n)
		}
		if n == 0 {
			break
		}
		models = append(models, string(buf[:n]))
	}

	return models, nil
}

func FetchHuggingFaceModel(name string, repo string, file string) error {
	nameBytes := []byte(name)
	repoBytes := []byte(repo)
	fileBytes := []byte(file)

	errBuf := make([]byte, 1024)
	var errLen uint32

	ret := symbols.FetchHFModel(
		&nameBytes[0], uint32(len(nameBytes)),
		&repoBytes[0], uint32(len(repoBytes)),
		&fileBytes[0], uint32(len(fileBytes)),
		&errBuf[0], &errLen,
	)

	if ret != 0 {
		return fmt.Errorf("%s", string(errBuf[:errLen]))
	}

	return nil
}
