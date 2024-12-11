package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/gen2brain/beeep"
)

const (
	markname="cli_reminder"
	markValue="1"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s Enter as <hh:mm> <text/message>\n",os.Args[0])
		os.Exit(1)
	}
	now:=time.Now();
	w:=when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t,err :=w.Parse(os.Args[1],now)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if t==nil {
		fmt.Print("mention time pls")
		os.Exit(2) 
	}
	if now.After(t.Time){
		fmt.Printf("time travel aint possible child :/")
		os.Exit(3)
	}
	diff:=t.Time.Sub(now)
	if os.Getenv(markname)==markValue{
		time.Sleep(diff)
		err = beeep.Alert("Reminder:",strings.Join(os.Args[2:]," "), "assets/information.png")
		if err!=nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd:= exec.Command(os.Args[0],os.Args[1:]...)
		cmd.Env = append(os.Environ(),fmt.Sprintf("%s=%s",markname,markValue))
		if err = cmd.Start(); err !=nil{
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed after ",diff.Round(time.Second))
		os.Exit(0)
	}
	fmt.Printf("Message to be displayed in %s",diff)
}
