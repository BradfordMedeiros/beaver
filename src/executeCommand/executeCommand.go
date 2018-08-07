package executeCommand

import "fmt"

type Command struct {
	
}

func GetExecuteCommand() func() {

	command := func(){
		fmt.Println("executecommand placeholder")
	}

	return command
}