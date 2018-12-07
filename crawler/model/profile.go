package model

import "encoding/json"

type Profile struct {
	Title  string
	Price  int
	Type   string
	Door   string
	Area   int
	Photos []string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
