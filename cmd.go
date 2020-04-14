package cmdhub

import (
    "bufio"
    "fmt"
    "io"
    "os/exec"
)

type CmdConfig struct{
	Cmd string
	Args []string
} 


type CmdChain struct{
	cmds []CmdConfig
}


func (self *CmdChain) Add(cmd string,args ...string) *CmdChain{
	cmdConfig := CmdConfig{
		Cmd : cmd,
		Args : args,
	}
    self.cmds = append(self.cmds,cmdConfig)
    return self
}


func (self *CmdChain) Run(){
	for _,v := range self.cmds{
		execOne(v.Cmd,v.Args)
	}
}

func execOne(sh string, args []string) (int, error) {
    cmd := exec.Command(sh, args...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println(err)
        return 0, nil
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        fmt.Println(err)
        return 0, nil
    }
    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        return 0, nil
    }
    s := bufio.NewScanner(io.MultiReader(stdout, stderr))
    for s.Scan() {
        text := s.Text()
        fmt.Println(text)
    }
    if err := cmd.Wait(); err != nil {
        fmt.Println(err)
    }
    return 0, nil
}