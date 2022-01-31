import random
import time
def binary_search(sequence, item):
    beggining = 0
    end = len(sequence) - 1

    while beggining <= end:
        mid = beggining + (end-beggining) // 2
        mid_value = sequence[mid]

        if mid_value == item:
            return mid
        
        elif item < mid_value:
            end = mid-1
        else:
            beggining = mid + 1
    return None


array = []
i=0
while i < 10000000:
    array.append(random.randint(0,100))
    i += 1

print(len(array))

array.sort()
beggining = time.time()

print(binary_search(array, 16))
end = time.time()
print(end-beggining)