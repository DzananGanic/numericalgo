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
Below are the examples of usage (at least one method from each category). Methods that are not explained are used in the same way as explained methods.

### Interpolation

#### Single value interpolation
For the single value interpolation, we need to define slices for x and y, which will be converted to coordinate pairs and sorted. We fit the slices into our interpolation type, and call interpolate.WithSingle function which receives the two parameters - first parameter is any type which conforms to validateInterpolator interface (thus implementing Interpolate and Validate methods), and the second parameter is the float value we want to interpolate. It returns the interpolated value, and the error (if it exists).

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
With multiple value interpolation, instead of declaring a single value, we define a slice of values we want to interpolate. Instead of calling interpolate.WithSingle, we call interpolate.WithMulti. Other details are the same.

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
For the single value prediction, we define two vectors - x and y. We fit the vectors, and call Predict method on defined type while passing the value to predict parameter.

``` 
x := numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5}
y := numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89}
valToPred := 1.9

lf := linear.New()
err := lf.Fit(x, y)
result := lf.Predict(valToPred)
```

#### Multiple values prediction
If we want to predict multiple values, instead of calling Predict method on defined type, we call fit.PredictMulti helper method, while passing our fit type (which conforms to predictor interface) as the first parameter, and the vector of values we want to predict.
```
x := numericalgo.Vector{0.3, 0.8, 1.2, 1.7, 2.4, 3.1, 3.8, 4.5, 5.1, 5.8, 6.5}
y := numericalgo.Vector{8.61, 7.94, 7.55, 6.85, 6.11, 5.17, 4.19, 3.41, 2.63, 1.77, 0.89}
valsToPred := numericalgo.Vector{6.63, 7.21}

lf := linear.New()
err := lf.Fit(x, y)
result := fit.PredictMulti(lf, valsToPred)
```

### Integrate
We need to define the f(x) function with the signature func (float64) float64. 'l' is our left bound and 'r' is our right bound for the integration. 'n' is the number of subdivisions (the higher the number, the more precise our result will be). Then we just call integrate.Trapezoid function and pass the defined values. It is the same thing with Simpsons rule (we just call integrate.Simpson)
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
When differentiating, we need to define f(x) function with the signature func (float64) float64, value and the step size (the smaller the step size - the better the precision). Then we call the Central, Backward or Forward function from the differentiate package
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
#### Bisection
When performing root finding, we define f(x) function with the signature func (float64) float64, eps as the tolerance, and left and right bounds of the function. Then we simply call the Bisection function from the root package while passing those parameters.
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
#### Newton's method
Newton's method usage differs from the bisection. Newton function from differentiate package receives the function with the same signature as with bisection, but here second parameter is the initial guess (which is reasonably close to the true root) and the third parameter is the number of iterations for the Newton method.
```
f := func(x float64) float64 {
  return math.Pow(x, 3) - 2*math.Pow(x, 2) + 5
}
initialGuess := -1
iter := 3

result := root.Newton(f, initialGuess, iter)
// -1.2419
```
