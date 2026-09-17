class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Queue:
    def __init__(self):
        self.head = None
        self.tail = None

    def enqueue(self, data):
        new_node = Node(data)
        if self.tail is None:
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node

    def dequeue(self):
        if self.head is None:
            raise IndexError("dequeue from empty queue")
        popped = self.head
        self.head = self.head.next
        if self.head is None:
            self.tail = None
        return popped.data

    def front(self):
        if self.head is None:
            raise IndexError("front from empty queue")
        return self.head.data

    def is_empty(self):
        return self.head is None