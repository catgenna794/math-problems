def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)

numbers = [1, 4, 2, 5, 7, 6, 3, 8, 9]
sorted_numbers = quicksort(numbers)
print(sorted_numbers)
