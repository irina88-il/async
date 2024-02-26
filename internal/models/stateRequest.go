package models

type StateRequest struct {
	AccessToken int `json:"access_key"`
	State       int `json:"state"`
}
