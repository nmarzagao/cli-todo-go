package main 

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[91m"
	GREEN  = "\033[92m"
	YELLOW = "\033[33m"
)

func print_list(list *List, cur_pos int) {
	if list.is_empty() {
		return
	}

	tmp := list.head

	for i := 1; tmp != nil; i++{

		var status_str string
		if !tmp.status {
			status_str = RED + "Not Done" + RESET
		} else {
			status_str = GREEN + "Done" + RESET
		}

		if i == cur_pos {
			fmt.Printf(" > " + YELLOW + "%d. " + RESET +  "%s [%s]\n", i, tmp.title, status_str)
		} else {
			fmt.Printf("   " + YELLOW + "%d. " + RESET +  "%s [%s]\n", i, tmp.title, status_str)
		}

		tmp = tmp.next
	}
}

func menu(list *List) {
	cur_pos := 1
	var input []byte = make([]byte, 1)

	// disable terminal for input
	// fix ctr C because it breaks the program
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
    
    for {
		cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()

		fmt.Println(GREEN + "\t\tTODO LIST" + RESET + "\n")
		print_list(list, cur_pos)

		fmt.Println(input)
		fmt.Println(cur_pos)

		os.Stdin.Read(input)
		if input[0] == 119 || input[0] == 65 { // up
			if cur_pos > 1 {
				cur_pos -= 1
			}
		} else if input[0] == 115 || input[0] == 66 { // down
			if cur_pos < list.size {
				cur_pos += 1
			}
		} else if input[0] == 32 || input[0] == 10 {
			list.set_task(cur_pos)
		} else if input[0] == 113 {
			break
		}

    }

	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}