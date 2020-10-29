package test

type ControllerMock struct {
	Token map[string]interface{}
	Err   error
}

func (c *ControllerMock) Login(name, pass string) (map[string]interface{}, error) {
	return c.Token, c.Err
}
