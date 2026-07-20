package model

type FontInfo struct {
	Prefix   string `json:"prefix"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Encoding string `json:"encoding"`
	Embedded bool   `json:"embedded"`
}
