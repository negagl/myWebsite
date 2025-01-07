package helpers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ExtractIDFromPath(w http.ResponseWriter, r *http.Request, moduleName string) (int, error) {
	prefix := "/" + moduleName + "/"
	idString := strings.TrimPrefix(r.URL.Path, prefix)

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, errors.New("Invalid ID")
	}

	return id, nil
}
