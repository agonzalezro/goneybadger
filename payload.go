package goneybadger

import (
	"runtime"
	"strconv"
)

const MAX_PCS = 10

type Payload struct {
	Notifier *Notifier `json:"notifier"`
	Error    *Error    `json:"error"`
	Server   *Server   `json:"server"`
}

type Notifier struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

type Error struct {
	Backtrace []*Backtrace `json:"backtrace"`
	Message   string       `json:"message"`
}

type Backtrace struct {
	Number string `json:"number"`
	File   string `json:"file"`
}

type Server struct {
	EnvironmentName string `json:"environment_name"`
	Hostname        string `json:"hostname"`
}

func NewPayload(hostname, env, message string) *Payload {
	payload := Payload{
		Notifier: &Notifier{
			Name:    "goneybadger",
			URL:     "https://github.com/agonzalezro/goneybadger",
			Version: "0.1",
		},
		Error: &Error{
			Message: message,
		},
		Server: &Server{
			EnvironmentName: env,
			Hostname:        hostname,
		},
	}

	pcs := make([]uintptr, MAX_PCS)

	// 3 is the needed offset to get the caller as first position.
	runtime.Callers(3, pcs)
	for i := 0; i <= MAX_PCS && pcs[i] != 0; i++ {
		pc := pcs[i]
		file, line := runtime.FuncForPC(pc).FileLine(pc)
		bt := Backtrace{
			File:   file,
			Number: strconv.Itoa(line),
		}
		payload.Error.Backtrace = append(payload.Error.Backtrace, &bt)
	}

	return &payload
}
