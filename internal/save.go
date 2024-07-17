package internal

import (
	"encoding/json"
	"os"

	"github.com/lormars/crawlmap/common"
)

func SaveToFile(filename string) error {
	data, err := json.MarshalIndent(common.Nodemap, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
