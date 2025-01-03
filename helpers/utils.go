package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func ExtractIDFromPath(w http.ResponseWriter, r *http.Request, moduleName string) (int, error) {
	idString := r.URL.Path[len(fmt.Sprintf("/%s/", moduleName)):]
	id, err := strconv.Atoi(idString)

	if err != nil {
		return -1, errors.New("Invalid ID")
	}

	return id, nil
}
