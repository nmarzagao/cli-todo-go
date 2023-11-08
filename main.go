package main

func main() {
	list := create_list()
	list.insert(create_task("Play minecraft"))
	list.insert(create_task("Study Go"))
	list.insert(create_task("Go to the gym"))

	list.set_task_done("Study Go")

	print_list(list)
}