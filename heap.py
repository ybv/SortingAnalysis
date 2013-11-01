import heapq
@profile
def main():
  heap =[]
  f = open('10MB', 'r+b')
  fw = open('bighpy', 'r+b')
  for line in f:
    heap.append(line)

  heapq.heapify(heap)

  while heap:
    fw.write(heapq.heappop(heap))


if __name__ == '__main__':
    main()