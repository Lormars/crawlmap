package internal

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/lormars/crawlmap/common"
)

func SaveToFile(dirname string) error {
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}

	for key, node := range common.Nodemap {
		data, err := json.MarshalIndent(node, "", "  ")
		if err != nil {
			return err
		}
		filepath := filepath.Join(dirname, key+".json")
		err = os.WriteFile(filepath, data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
