package model

import (
	"encoding/json"
	"io"
)

type UserFollowing struct {
	UserId                string `json:"userId,omitempty"`
	Username              string `json:"username,omitempty"`
	Image          string `json:"image,omitempty"`
	FollowingUserId       string `json:"followingUserId,omitempty"`
	FollowingUsername     string `json:"followingUsername,omitempty"`
	FollowingImage string `json:"followingImage,omitempty"`
}

func (o *UserFollowing) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *UserFollowing) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
