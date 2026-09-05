class Node:
    def __init__(self,data):
        self.data = data 
        self.next = None 

# skapa länkad lista
node_1 = Node(5)
node_2 = Node(10)
node_3 = Node(15)

node_1.next = node_2
node_2.next = node_3

# Insert-at-beginning: lägga till en ny nod före nuvarande head.
# Tanken: den nya noden ska peka på det som var head, och sen ska head uppdateras till att peka på den nya noden.
# head / node = ger alltså det HELA elementet som den pekar på ex node_1
# node.data ger alltså elementets faktiska värde ex 5

def insert_at_beginning(head, data):    # this function takes current ehead and inserts the node (data) you want to be first
    new_node = Node(data)
    new_node.next = head
    return new_node
    
# testa 
head = node_1
new_head = insert_at_beginning(head, 0)

current_node = new_head # start at head 

while current_node != None: # continue until node.next is None
    print(current_node.data) # print its actuall content
    current_node = current_node.next # update the value of "node"