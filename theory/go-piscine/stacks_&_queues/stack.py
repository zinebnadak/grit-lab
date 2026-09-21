@dataclass
class Student:
    name: str
    ects: int 

@dataclass
class Stack:
    contents: list[Student] = field(default_factory=list) 
    '''
    field(default_factory=list) säger: skapa en ny tom lista varje gång ett Stack-objekt skapas.
    De skulle dela samma lista — pusha till en `Stack` och den andra `Stack`-instansen skulle också få det elementet, trots att de är separata objekt.
    `default_factory=list` löser det genom att anropa `list()` på nytt för varje instans, så var och en får sin egen tomma lista.
    '''

    def push(self, student: Student) -> None:
        self.contents.append(student)

    def pop(self) -> Student:
        return self.contents.pop()

    def peek(self) -> Student:
        return self.contents[-1]

    def is_empty(self) -> bool:
        return len(self.contents) == 0

