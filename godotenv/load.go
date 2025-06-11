package godotenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Load reads the specified files and sets environment variables found within.
// If no filenames are provided, it tries to load .env.
func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}
	for _, name := range filenames {
		if err := loadFile(name); err != nil {
			return err
		}
	}
	return nil
}

func loadFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return parse(f)
}

func parse(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		os.Setenv(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
	}
	return scanner.Err()
}
