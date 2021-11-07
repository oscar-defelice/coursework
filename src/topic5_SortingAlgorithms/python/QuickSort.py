def quicksort(arr):
    if len(arr) <= 1:
        return arr

    left = []
    right = []
    pivot = arr[0]
    center = [pivot]

    for x in arr:
        if x > pivot:
            right.append(x)
        elif x < pivot:
            left.append(x)

    result = quicksort(left) + center + quicksort(right)
    print(*result)
    return result


n = input()
quicksort([int(x) for x in input().split()])
