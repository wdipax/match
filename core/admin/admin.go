package admin

type Admin struct{}

func New() *Admin {
	return &Admin{}
}

func (a *Admin) NewSession() error {
	return nil
}

func (a *Admin) EndSession() error {
	return nil
}
