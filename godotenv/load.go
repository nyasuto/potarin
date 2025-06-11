package godotenv

import (
	"os"
	"strings"
)

// Load sets environment variables defined in the given files. If no file is
// specified, it defaults to ".env". This is a small subset of the real
// github.com/joho/godotenv package.
func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}
	for _, name := range filenames {
		data, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if kv := strings.SplitN(line, "=", 2); len(kv) == 2 {
				os.Setenv(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
			}
		}
	}
	return nil
}
