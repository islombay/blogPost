package objects

import "encoding/json"

func ReObject(a, b interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, b)
	return err
}
