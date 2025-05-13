package service

type HelloWorldService struct {}

func (s *HelloWorldService) HelloWorld() string {
	return "Hello, World!"
}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}
