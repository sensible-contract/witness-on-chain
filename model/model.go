package model

import "encoding/json"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (t *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

////////////////
type PublicKey struct {
	Rabin string `json:"rabin"`
}

type Welcome struct {
	PublicKey PublicKey `json:"public_key"`
	Contact   string    `json:"contact"`
	Message   string    `json:"message"`
}

////////////////
type Signature struct {
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Padding   string `json:"padding"`
}

type Signatures struct {
	Rabin Signature `json:"rabin"`
}

type TimestampResponse struct {
	Timestamp  int        `json:"timestamp"`
	UTC        string     `json:"utc"`
	Digest     string     `json:"digest"`
	Signatures Signatures `json:"signatures"`
}

func (t *TimestampResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
