package server

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

//注册服务对象并开启该RPC服务
//arith :=new(Arith)
//rpc.Register(arith)
//rpc.HandleHTTP()
//l, e := net.Listen("tcp", ":1234")
//if e != nil {
//log.Fatal("listen error:", e)
//}
//go http.Serve(l, nil)
