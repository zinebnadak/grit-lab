# 1. Skapa en klass och länka noderna
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None 

    
node_1 = Node(1)
node_2 = Node(2)
node_3 = Node(3)

head = node_1

# skapa linked lista 
node_1.next = node_2
node_2.next = node_3

'''
# skriv ut
print(head.data, node_2.data, node_3.data)
'''

'''
# 2. loopa igenom från head till slut, UTAN att använda en lista
current_node = head
while current_node != None:
    print(current_node.data)
    current_node = current_node.next
'''

# 3. funktion som byter ut/ sätter in element i början 

def insert_at_beginning(current_head, value):
    new_head = Node(value)
    new_head.next = current_head
    return new_head 

new_head = insert_at_beginning(head, 0)
current_node = new_head

while current_node != None:
    print(current_node.data)
    current_node = current_node.next
