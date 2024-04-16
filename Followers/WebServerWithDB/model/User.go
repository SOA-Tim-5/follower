package model

import (
	"encoding/json"
	"io"
)




type User struct {
	Id    string `json:"Id,omitempty"`
	Username  string  `json:"Username,omitempty"`
	Image string `json:"Image,omitempty"`
}

type Users []*User

func (o *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
