// state maps incoming requests to the inner logic of the application.
package state

type State struct {
	admins []string
	core   CoreHandler

	userSession map[string]string
}

type CoreHandler interface{}

type Settings struct {
	Admins []string
	Core   CoreHandler
}

func New(settings *Settings) *State {
	return &State{
		admins: settings.Admins,
		core:   settings.Core,

		userSession: make(map[string]string),
	}
}
