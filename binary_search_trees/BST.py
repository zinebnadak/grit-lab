# Regel för in-order: vänster → nod → höger.


class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


class BST:
    def __init__(self):
        self.root = None

    def insert(self, data):
        if self.root is None:
            self.root = Node(data)
            return
        current = self.root
        while True:
            if data < current.data:
                if current.left is None:
                    current.left = Node(data)
                    return
                current = current.left
            else:
                if current.right is None:
                    current.right = Node(data)
                    return
                current = current.right

    def search(self, data):
        current = self.root
        while current:
            if data == current.data:
                return True
            current = current.left if data < current.data else current.right
        return False

    def in_order(self, node=None, result=None):
    if node is None and result is None:
        node = self.root
        result = []
    if node is None:
        return result
    self.in_order(node.left, result)
    result.append(node.data)
    self.in_order(node.right, result)
    return result