package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func ExtractIDFromPath(w http.ResponseWriter, r *http.Request, moduleName string) (id int, err error) {
	idString := r.URL.Path[len(fmt.Sprintf("/%s/", moduleName)):]
	id, err = strconv.Atoi(idString)

	if err != nil {
		return -1, err
	}

	return id, nil
}
