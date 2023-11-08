package main

func main() {
	
	
	list := create_list()
	list.insert(create_task("Play minecraft"))
	list.insert(create_task("Study Go"))
	list.insert(create_task("Go to the gym"))
	list.insert(create_task("Option 4"))
	list.insert(create_task("Option 5"))
	list.insert(create_task("Option 6"))
	list.insert(create_task("Option 7"))
	list.insert(create_task("Option 8"))

	menu(list)
	
}