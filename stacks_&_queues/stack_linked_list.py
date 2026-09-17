class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Stack:
    def __init__(self):
        self.top = None

    def push(self, data):
        new_node = Node(data)
        new_node.next = self.top
        self.top = new_node

    def pop(self):
        if self.top is None:
            raise IndexError("pop from empty stack")
        popped = self.top
        self.top = self.top.next
        return popped.data

    def peek(self):
        if self.top is None:
            raise IndexError("peek from empty stack")
        return self.top.data

    def is_empty(self):
        return self.top is None