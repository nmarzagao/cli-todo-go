package main 

import "fmt"

func print_list(list *List) {
	tmp := list.head

	for i := 1; tmp != nil; i++{

		var status_str string
		if !tmp.status {
			status_str = "\033[91m" + "Not Done" + "\033[0m"
		} else {
			status_str = "\033[92m" + "Done" + "\033[0m"
		}

		fmt.Printf("\033[33m" + "%d. " + "\033[0m" +  "%s [%s]\n", i, tmp.title, status_str)

		tmp = tmp.next
	}
}

func menu(list *List) {
	return 
}