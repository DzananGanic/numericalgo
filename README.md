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
  - Newton's method
- [Numerical Differentiation](https://github.com/DzananGanic/numericalgo/tree/master/differentiate)
  - Backward difference formula
  - Forward difference formula
  - Central difference formula
- [Numerical Integration](https://github.com/DzananGanic/numericalgo/tree/master/integrate)
  - Trapezoidal rule integration
  - Simpson’s rule integration

With numericalgo, it is also possible to solve linear equations and work with matrices and vectors, as those types are provided.

## Usage
Below are the examples of usage (one method from each category)

### Interpolation

#### Single value interpolation

``` 
x := []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2}
y := []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07}
valToInterp := 5.1

li := linear.New()
li.Fit(x, y)

estimate, err := interpolate.WithSingle(li, valToInterp)
// (3.1566666666666663, nil)
```

#### Multiple values interpolation

```
x := []float64{1.3, 1.8, 2.5, 3.1, 3.8, 4.4, 4.9, 5.5, 6.2}
y := []float64{3.37, 4.45, 4.81, 3.96, 3.31, 2.72, 3.02, 3.43, 4.07}
valsToInterp := []float64{2.2, 5.1, 1.5}

li := linear.New()
li.Fit(x, y)

estimate, err := interpolate.WithMulti(li, valToInterp)
([]float64{4.655714285714286, 3.1566666666666663, 3.802}, nil)
```

### Fit

#### Single value prediction
``` 
x := numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5}
y := numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89}
valToPred := 1.9

lf := linear.New()
err := lf.Fit(x, y)
result := lf.Predict(valToPred)
```

#### Multiple values prediction
```
x := numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5}
y := numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89}
valsToPred := numericalgo.Vector{6.63, 7.21}

lf := linear.New()
err := lf.Fit(x, y)
result := fit.PredictMulti(lf, valsToPred)
```

### Integrate

```
f := func(x float64) float64 {
  return math.Sin(x)
}
l := 0
r := math.Pi / 2
n := 20

result, err := integrate.Trapezoid(f, l, r, n)
// (0.999, nil)
```

### Differentiate

```
f := func(x float64) float64 {
  return math.Cos(math.Pow(x,2) - 2)
}
val := 1
h := 0.1

result, err := differentiate.Central(f, val, h)
// (1.6609, nil)
```

### Root finding

```
f := func(x float64) float64 {
  return math.Pow(x, 2)
}
eps := 0.01
l := -1
r := 1

result := root.Bisection(f, eps, l, r)
// 0
```
