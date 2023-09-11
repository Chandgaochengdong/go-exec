package provider

type HelloService struct {
}

func (p *HelloService) SayHello(request string, reply *string) (*string, error) {
	*reply = "hello " + request
	return reply, nil
}
