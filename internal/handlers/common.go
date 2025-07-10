package handlers

import (
	"log/slog"

	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/google/uuid"
)

func withErr(e error) slog.Attr {
	return applog.UnwrapErrorChain(e)
}

func validateId(id string) bool {
	return uuid.Validate(id) == nil
}
