package main

import (
    "github.com/TarsCloud/TarsGo/tars"

    "github.com/tarsprotocol/test/helloworld" 
)

type GreeterImp  struct {
}

func (imp *GreeterImp) SayHello(input helloworld.HelloRequest)(output helloworld.HelloReply, err error) {
    output.Message = "hello" +  input.GetName() 
    return output, nil 
}

func main() { //Init servant

    imp := new(GreeterImp)                             //New Imp
    app := new(helloworld.Greeter)                            //New init the A JCE
    cfg := tars.GetServerConfig()                      //Get Config File Object
    app.AddServant(imp, cfg.App+"."+cfg.Server+".GreeterTestObj") //Register Servant
    tars.Run()
}
