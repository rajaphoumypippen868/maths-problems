def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

def max_product_subarray(nums):
    max_so_far = nums[0]
    max_current = nums[0]

    for i in range(1, len(nums)):
        max_current = max(nums[i], max_current + nums[i])
        if max_current > max_so_far:
            max_so_far = max_current

    return max_so_far
