package utils

import "encoding/json"

func DataToConfig(src, dst interface{}) error{
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, dst)
	if err != nil{
		return err
	}

	return nil
}
