int factorial(int x) {
    if (x == 1) return x;
    return factorial(x-1) * x;
}