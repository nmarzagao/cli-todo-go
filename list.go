package main

type Task struct {
	id     int
	title  string
	status bool

	next   *Task
}

type List struct {
	head *Task
	size int
}


func create_list() *List {
	return &List {head: nil}
}

func create_task(t string) *Task {
	return &Task {
		title:  t,
		status: false,
		next:   nil,
	}
}


func (l *List) is_empty() bool {
    return (l.head == nil)
}

func (l *List) insert(t *Task) {

    if l.is_empty() {
        l.head = t
    } else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}

		curr.next = t
	}

	l.size += 1
}

func (l *List) remove(id int) bool {
	if l.is_empty() {
        return false
    } else if id == 1 {
		l.size -= 1
        l.head = l.head.next
        return true
    }

    curr := l.head
	count := 1
    for curr.next != nil && count != id  {
        curr = curr.next
		count += 1
    }

    if curr.next != nil {
        curr.next = curr.next.next
		l.size -= 1

		return true
    } else {
		return false
	}
}

func (l *List) set_task(id int) bool {
	if l.is_empty() {
		return false
	} else if id == 1 {
		if l.head.status {
			l.head.status = false
		} else {
			l.head.status = true
		}
		return true
	}

	curr := l.head
	count := 1
	for curr != nil && count != id {
		curr = curr.next
		count += 1
	}

	if curr != nil {
		if curr.status {
			curr.status = false
		} else {
			curr.status = true
		}
		return true
	} else {
		return false
	}
}

