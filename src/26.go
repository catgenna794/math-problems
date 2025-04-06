def is_prime(n):
    if n < 2:
        return False
    for i in range(2, int(n ** 0.5) + 1):
        if n % i == 0:
            return False
    return True

n = 7
while n < 100:
    if is_prime(n):
        print(n)
    n += 1
