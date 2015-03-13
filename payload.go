package goneybadger

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
	Backtrace []*struct{} `json:"backtrace"` // The key needs to be on the payload
	Message   string      `json:"message"`
}

type Server struct {
	EnvironmentName string `json:"environment_name"`
	Hostname        string `json:"hostname"`
}

func NewPayload(hostname, env, message string) *Payload {
	return &Payload{
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
}
