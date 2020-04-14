package main

import("cmdhub")

func main() {
	cmdChain := cmdhub.CmdChain{}
	cmdChain.Add("ls","-a").Add("pwd","-L")
	cmdChain.Run()
}