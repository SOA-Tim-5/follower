package model

import (
	"encoding/json"
	"io"
)

type Following struct {
	UserId                string `json:"userId,omitempty"`
	Username              string `json:"username,omitempty"`
	ProfileImage          string `json:"profileImage,omitempty"`
	FollowingUserId       string `json:"followingUserId,omitempty"`
	FollowingUsername     string `json:"followingUsername,omitempty"`
	FollowingProfileImage string `json:"followingProfileImage,omitempty"`
}

func (o *Following) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Following) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
