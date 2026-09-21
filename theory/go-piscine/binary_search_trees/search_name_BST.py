from dataclasses import dataclass


@dataclass
class Student:
    name: str
    ects: int


class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


class BinaryTree:
    def __init__(self):
        self.root = None

    def insert(self, data):
        if self.root is None:
            self.root = Node(data)
            return
        current = self.root
        while True:
            if data.ects < current.data.ects:
                if current.left is None:
                    current.left = Node(data)
                    return
                current = current.left
            else:
                if current.right is None:
                    current.right = Node(data)
                    return
                current = current.right


def search_name(tree: BinaryTree, name: str) -> Student | None:
    return _search_name(tree.root, name)


def _search_name(node: Node, name: str) -> Student | None:
    if node is None:
        return None
    if node.data.name == name:
        return node.data
    found = _search_name(node.left, name)
    if found is not None:
        return found
    return _search_name(node.right, name)


if __name__ == "__main__":
    tree = BinaryTree()
    tree.insert(Student("Alice", 30))
    tree.insert(Student("Bob", 15))
    tree.insert(Student("Carla", 45))
    tree.insert(Student("David", 10))

    result = search_name(tree, "Carla")
    print(result)  # Student(name='Carla', ects=45)

    result = search_name(tree, "Nobody")
    print(result)  # None