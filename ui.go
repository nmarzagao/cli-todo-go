package main 

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const (
	// COLORS
	RESET  = "\033[0m"
	RED    = "\033[91m"
	GREEN  = "\033[92m"
	YELLOW = "\033[33m"

	// CONTROLLS
	W_KEY     = 119
	UP_KEY    = 65
	J_KEY	  = 106
	S_KEY     = 115
	DOWN_KEY  = 66
	K_KEY     = 107
	ENTER_KEY = 10
	SPACE_KEY = 32
	Q_KEY 	  = 113
	A_KEY	  = 97
	D_KEY     = 100
)

func disable_CtrC() {
	// Handle interrupt signal (Ctrl+C)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		// Handle interrupt (Ctrl+C) here if necessary
		// For example, you can print a message instead of breaking the loop
		fmt.Println("Ctrl+C is disabled in this program.")
	}()
}

func clear_screen() {
	fmt.Print("\033[H\033[2J")
}

func disable_terminal() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func enable_terminal() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func add_item_screen() {

}

func print_list(list *List, cur_pos int) {
	fmt.Println(GREEN + "\t\tTODO LIST" + RESET + "\n")

	if list.is_empty() {
		fmt.Println("List is Empty!\n")
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
	exit := false

	disable_terminal()
	disable_CtrC()
    
    for !exit {
		clear_screen()
		print_list(list, cur_pos)

		fmt.Println("\n\n[ENTER] Mark 'Done/Not Done' [A] Add Item [D] Remove Item\n\n")

		os.Stdin.Read(input)

		switch input[0] {
		case W_KEY, UP_KEY, J_KEY:
			if cur_pos > 1 {
				cur_pos -= 1
			}
		
		case S_KEY, DOWN_KEY, K_KEY:
			if cur_pos < list.size {
				cur_pos += 1
			}
		
		case ENTER_KEY, SPACE_KEY:
			list.set_task(cur_pos)

		case A_KEY:
			// Add

		case D_KEY:
			list.remove(cur_pos)

		case Q_KEY:
			exit = true
		
		default:
			continue
		}
    }
	enable_terminal()
}