package hello

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (hs HelloService) SayHello() string {
	return "Hello !"
}

func (hs HelloService) SayGoodBye() string {
	return "Bybye :/"
}
