package errorfa

import "encoding/json"

type (
	ErrorFarsi interface {
		Error() string
		ErrorFarsi() string
		JSON() string
	}
	MyError struct {
		Text      string `json:"error"`
		TextFarsi string `json:"errorFarsi"`
	}
)

func New(text, farsiText string) {

}

func (e *MyError) Error() string {
	return e.Text
}

func (e *MyError) ErrorFarsi() string {
	return e.TextFarsi
}

func (e *MyError) JSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
