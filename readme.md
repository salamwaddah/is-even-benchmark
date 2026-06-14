# is-even-benchmark

I benchmarked different ways to check if a number is-even so you don't have to.

### Negative numbers

They are not supported, but feel free to run the `recursive()` on negative numbers if you feel like waiting forever.

I would love to see each benchmarked. Compare how long it takes each function to chew through a list of 1,000,000 random
integers and spit out a list of 1,000,000 booleans

## Time complexities

### Constant time: O(1)

Takes the same amount of time regardless of the size of input.

### Logarithmic time: O(log(n))

Reduces the problem size by a factor in each step.

### Linear time: O(n)

Time grows proportionally with the size of the input.

### Quadratic Time: O(n^2)

Time grows proportional to the square of the input size, often due to nested loops.

### Exponential Time: O(2^n)

The time doubles with each additional input, making it extremely inefficient for large inputs.

### Factorial Time: O(n!)

The time grows as the factorial of the input size, often seen in problems involving permutations.