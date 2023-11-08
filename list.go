package main

type Task struct {
	title  string
	status bool

	next   *Task
}

type List struct {
	head *Task
	size int
}


func create_list() *List {
	return &List {head: nil, size: 0}
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


func (l *List) remove(value string) bool {
	if l.is_empty() {
        return false
    } else if l.head.title == value {
		l.size -= 1
        l.head = l.head.next
        return true
    }

    curr := l.head
    for curr.next != nil && curr.next.title != value {
        curr = curr.next
    }

    if curr.next != nil {
        curr.next = curr.next.next
		l.size -= 1

		return true
    } else {
		return false
	}
}

func (l *List) set_task_done(value string) bool {
	if l.is_empty() {
        return false
    } else if l.head.title == value {
		l.head.status = true
        return true
    }

    curr := l.head
    for curr.next != nil && curr.next.title != value {
        curr = curr.next
    }

    if curr.next != nil {
		curr.next.status = true
		return true
    } else {
		return false
	}
}
