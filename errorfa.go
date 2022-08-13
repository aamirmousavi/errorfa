package errorfa

import "encoding/json"

type (
	ErrorFarsi interface {
		Error() string
		ErrorFarsi() string
		AppName() string
		JSON() string
	}
	MyError struct {
		Appname   string `json:"appName"`
		Text      string `json:"error"`
		TextFarsi string `json:"errorFarsi"`
	}
)

func New(appName, text, farsiText string) ErrorFarsi {
	return &MyError{
		Appname:   appName,
		Text:      text,
		TextFarsi: farsiText,
	}
}
func (e *MyError) AppName() string {
	return e.Appname
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
