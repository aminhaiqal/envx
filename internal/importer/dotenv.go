package importer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/axelyn/envx/pkg/envx"
)

type Importer struct{}

func New() *Importer {
	return &Importer{}
}

func (i *Importer) ImportFromDotenv(filePath string) (map[string]envx.Variable, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	variables := make(map[string]envx.Variable)
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format at line %d: %s", lineNumber, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		value = strings.Trim(value, "\"'")
		
		value = strings.ReplaceAll(value, "\\\"", "\"")

		variables[key] = envx.Variable{
			Key:       key,
			Value:     value,
			IsSecret:  false,
			Base: envx.Base{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return variables, nil
}

func (i *Importer) PreviewImport(filePath string, existing map[string]envx.Variable) (new, updated, unchanged []string, err error) {
	toImport, err := i.ImportFromDotenv(filePath)
	if err != nil {
		return nil, nil, nil, err
	}

	for key, importVar := range toImport {
		if existingVar, exists := existing[key]; exists {
			if existingVar.Value != importVar.Value {
				updated = append(updated, key)
			} else {
				unchanged = append(unchanged, key)
			}
		} else {
			new = append(new, key)
		}
	}

	return new, updated, unchanged, nil
}
