# python class vd @Dataclass
# __init__ is a "constructure", when you run the program python immidiatley looks for anmethod called __init__ on the class
# .data and -next is both attributes, and to assign a value to them we use node_x.attribute = value
# when printing we need to print with .data to accec the actual values of the node and not the memory adress

class Node:
    def __init__(self,data):
        self.data = data 
        self.next = None 

node_1 = Node(5)

head = node_1
node_2 = Node(10)
node_3 = Node(15)

node_1.next = node_2
node_2.next = node_3

print(head.data, node_2.data, node_3.data)