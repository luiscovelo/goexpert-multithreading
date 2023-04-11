package formatter

import "encoding/json"

func JSON(in interface{}) string {
	b, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}
