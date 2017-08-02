# numericalgo

numericalgo is a set of numerical methods implemented in Golang. The idea was to implement everything from scratch - not just the methods, but the custom types as well (matrices, vectors etc.)

## Currently implemented methods:
- [Interpolation:](https://github.com/DzananGanic/numericalgo/tree/master/interpolate)
  - [Linear](https://github.com/DzananGanic/numericalgo/tree/master/interpolate/linear)
  - [Lagrange](https://github.com/DzananGanic/numericalgo/tree/master/interpolate/lagrange)
- [Regressions (fits)](https://github.com/DzananGanic/numericalgo/tree/master/fit)
  - [Linear](https://github.com/DzananGanic/numericalgo/tree/master/fit/linear)
  - [Polynomial](https://github.com/DzananGanic/numericalgo/tree/master/fit/poly)
  - [Exponential](https://github.com/DzananGanic/numericalgo/tree/master/fit/exponential)
- [Root finding:](https://github.com/DzananGanic/numericalgo/tree/master/root)
  - Bisection
- [Numerical Differentiation](https://github.com/DzananGanic/numericalgo/tree/master/differentiate)
  - Backward difference formula
  - Forward difference formula
  - Central difference formula
- [Numerical Integration](https://github.com/DzananGanic/numericalgo/tree/master/integrate)
  - Trapezoidal rule integration
  - Simpson’s rule integration

With numericalgo, it is also possible to solve linear equations and work with matrices and vectors, as those types are provided.
