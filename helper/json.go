package helper

import jsonIterator "github.com/json-iterator/go"

var json = jsonIterator.ConfigCompatibleWithStandardLibrary

func ObjectMapper(in, out any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &out); err != nil {
		return err
	}

	return nil
}
