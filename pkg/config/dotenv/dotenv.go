package dotenv

import (
	"context"
	"strings"

	"github.com/heetch/confita/backend"

	"github.com/joho/godotenv"
)

func NewBackend(filenames ...string) backend.Backend {
	envMap, err := godotenv.Read(filenames...)
	if err != nil {
		panic(err)
	}

	return backend.Func("dotenv", func(ctx context.Context, key string) ([]byte, error) {
		if envMap[key] != "" {
			return []byte(envMap[key]), nil
		}

		key = strings.Replace(strings.ToUpper(key), "-", "_", -1)

		if v, ok := envMap[key]; ok {
			return []byte(v), nil
		}

		return nil, backend.ErrNotFound
	})
}
