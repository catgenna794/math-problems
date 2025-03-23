def find_primes(n):
    """
    This function finds all prime numbers up to n.
    
    Args:
    - n: An integer representing the upper limit (inclusive).
    
    Returns:
    - A list of integers representing the prime numbers up to n.
    """
    if n <= 1:
        return []
    
    primes = [2] * (n + 1)
    for i in range(3, n + 1):
        is_prime = True
        for j in range(n - i + 1, 0, -1):
            if i % j == 0:
                is_prime = False
                break
        if is_prime:
            primes[i] = 1
    
    return primes

# Example usage:
n = 30  # Change this value to find prime numbers up to any integer n
primes_up_to_n = find_primes(n)
print("Primes up to", n, ":", primes_up_to_n)
