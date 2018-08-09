package ioLoop

import "bufio"
import "os"

func StartRepl (commandChannel chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0{		// happens if user just hits enter, so filter it out
			continue
		}
		commandChannel <- text 
	}
}