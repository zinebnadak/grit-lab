@dataclass
class Student:
    name: str
    ects: int 

@dataclass
class Stack:
    contents: list[Student] = field(default_factory=list)

    def push(self, student: Student) -> None:
        self.contents.append(student)

    def pop(self) -> Student:
        return self.contents.pop()

    def peek(self) -> Student:
        return self.contents[-1]

    def is_empty(self) -> bool:
        return len(self.contents) == 0

    