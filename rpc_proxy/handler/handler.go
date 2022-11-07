package handler

const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

func (s NewHelloService) Hello(res string, reply *string) error {
	*reply = "hello, " + res
	return nil
}
