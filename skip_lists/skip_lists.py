class Node:
    def __init__(self, data):
        self.data = data
        self.next_bottom = None   # next node on the detailed level
        self.next_top = None      # next node on the "express lane"


class SkipList:
    def __init__(self):
        self.head = Node(None)  # sentinel, holds no real data

    def insert(self, data):
        # find where it goes on the bottom level
        current = self.head
        while current.next_bottom and current.next_bottom.data < data:
            current = current.next_bottom

        new_node = Node(data)
        new_node.next_bottom = current.next_bottom
        current.next_bottom = new_node

    def search(self, data):
        current = self.head
        while current.next_bottom and current.next_bottom.data < data:
            current = current.next_bottom
        current = current.next_bottom
        return current is not None and current.data == data