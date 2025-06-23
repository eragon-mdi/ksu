package handlers

import (
	"fmt"
	"log/slog"
	"strings"
)

// длинную заврапенную ошибку в json структуру {...},
// каждый ключ который - слой на котором была ошибка
func withErr(e error) slog.Attr {
	usedKeys := make(map[string]int)
	errs := make([]slog.Attr, 0)

	lines := strings.Split(e.Error(), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		key := "in-err"
		val := line

		if idx := strings.Index(line, ":"); idx != -1 {
			prefix := strings.TrimSpace(line[:idx])
			if prefix != "" {
				key = prefix
				val = strings.TrimSpace(line[idx+1:])
			}
		}

		// проблема уникальных ключей json
		if count, exists := usedKeys[key]; exists {
			usedKeys[key] = count + 1
			key = fmt.Sprintf("%s_%d", key, count+1)
		} else {
			usedKeys[key] = 0
		}

		errs = append(errs, slog.String(key, val))
	}

	return slog.Attr{
		Key:   "cause",
		Value: slog.GroupValue(errs...),
	}
}
