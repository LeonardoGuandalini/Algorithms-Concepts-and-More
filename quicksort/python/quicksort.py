import random
import time

def quicksort(sequence):
    size = len(sequence)
    if size <= 1:
        return sequence
    else:    

        pivot = sequence.pop()
        greater = []
        smaller = []

        for item in sequence:
            if item > pivot:
                greater.append(item)
            else:
                smaller.append(item)
    
    return quicksort(smaller) + [pivot] + quicksort(greater)

array = []
for i in range(0, 100000):
    array.append(random.randint(0, 100))

beggining = time.time()
quicksort(array)
end = time.time()

print("%s ms" % ((end-beggining)*1000))


