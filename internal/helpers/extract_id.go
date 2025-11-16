package helpers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ExtractIDFromPath(r *http.Request) (uint, error) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return 0, errors.New("invalid path format")
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID format")
	}

	return uint(id), nil
}
