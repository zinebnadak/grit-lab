from __future__ import annotations
from dataclasses import dataclass, field

@dataclass
class Student:
  name: str
  ects: int

@dataclass
class Heap:
  # lambda = "anonymous function", preallocate a list with 5 elements
  array: list[Student|None] = field(default_factory=lambda: [None for _ in range(5)])
  last_index: int = 0
  ALLOCATION_STRATEGY: int = 5

def insert(heap: Heap, student: Student):
  # Increase size of list in blocks of five
  if heap.last_index == len(heap.array)-1:
    heap.array = heap.array + [None]*heap.ALLOCATION_STRATEGY

  # New element is inserted in end of array / bottom of tree
  heap.last_index += 1
  heap.array[heap.last_index] = student

  # Bubble the new element up by repeated swapping with parent
  current = heap.last_index
  parent = current // 2

  # Note: Hardcoded comparison of ects property
  while parent >= 1 and heap.array[parent].ects < heap.array[current].ects:
    heap.array[parent], heap.array[current] = heap.array[current], heap.array[parent]
    current = parent
    parent = current // 2

def remove_max(heap: Heap) -> Student|None:
  if heap.last_index == 0:
    return None

  # Highest-priority element is always the first (ignoring 0-indexed element)
  result = heap.array[1]
  # Move last element to first, shrink array
  heap.array[1] = heap.array[heap.last_index]
  heap.array[heap.last_index] = None
  heap.last_index -= 1

  # Bubble new root down
  i = 1
  while True: # Note: relying on break to end the loop
    largest = i
    if 2*i <= heap.last_index and (heap.array[largest].ects < heap.array[2*i].ects):
      largest = 2*i
    if 2*i+1 <= heap.last_index and (heap.array[largest].ects < heap.array[2*i+1].ects):
      largest = 2*i+1

    if i!=largest:
      # Found a child which is larger than root; do the swap
      heap.array[i], heap.array[largest] = heap.array[largest], heap.array[i]
      i = largest # and move down one level in the tree
    else:
      # We are done, end the loop
      break

  return result

heap = Heap()
insert(heap, Student("Anna", 20))
print(heap.array)
insert(heap, Student("Kalle", 40))
print(heap.array)
insert(heap, Student("Ville", 10))
print(heap.array)
insert(heap, Student("Justus", 25))
print(heap.array)

top = remove_max(heap)
print(top)
print(heap.array)
