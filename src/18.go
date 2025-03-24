import numpy as np

class QuadraticEquationSolver:
    def __init__(self):
        self.a = 1.0
        self.b = -2.0
        self.c = 3.0

    def solve(self):
        delta = self.b**2 - 4 * self.a * self.c
        if delta < 0:
            return "No real roots"
        elif delta == 0:
            x1 = -self.b / (2 * self.a)
            return f"Real root: {x1}"
        else:
            x1 = (-self.b + np.sqrt(delta)) / (2 * self.a)
            x2 = (-self.b - np.sqrt(delta)) / (2 * self.a)
            return f"Two real roots: {x1}, {x2}"

# Example usage
equation_solver = QuadraticEquationSolver()
print(equation_solver.solve())
