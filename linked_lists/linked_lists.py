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

# Hela poängen med en länkad lista är att du inte har en extern lista med alla noder 
# ,du har bara head, och måste ta dig vidare via .next för att hitta nästa nod.

# Kolla None-villkoret på node-variabeln själv, inte på node.next, och gör kollen innan du printar

current_node = head # start at head 

while current_node != None: # continue until node.next is None
    print(current_node.data) # print its actuall content
    current_node = current_node.next # update the value of "node"
