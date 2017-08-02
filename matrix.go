package numericalgo

import "fmt"
import "math"

// Matrix type is the slice of Vectors, with custom methods needed for matrix operations.
type Matrix []Vector

// Dim returns the dimensions of the matrix in the form (rows, columns).
func (m Matrix) Dim() (int, int) {
	if m.isNil() {
		return 0, 0
	}
	return len(m), len(m[0])
}

// Invert returns the inverted matrix by using Gauss-Jordan elimination
func (m Matrix) Invert() (Matrix, error) {

	if !m.isSquare() {
		return nil, fmt.Errorf("Cannot invert non-square Matrix")
	}

	var rows, _ = m.Dim()

	vec := make(Vector, rows)

	// 1. Reduction to identity form
	for currentRow := 1; currentRow <= rows; currentRow++ {

		// Pivoting
		p := currentRow
		for i := currentRow + 1; i <= rows; i++ {
			if math.Abs(m[i-1][currentRow-1]) > math.Abs(m[p-1][currentRow-1]) {
				p = i
			}
		}

		// If there exists no element a(k,i) different from zero, matrix is singular and has none or more than one solution
		if math.Abs(m[p-1][currentRow-1]) < 1e-10 {
			return nil, fmt.Errorf("Matrix is singular")
		}

		// If we find pivot which is the largest a(i, currentRow), we swap the rows
		if p != currentRow {
			tmp := m[currentRow-1]
			m[currentRow-1] = m[p-1]
			m[p-1] = tmp
		}

		vec[currentRow-1] = float64(p)

		mi := m[currentRow-1][currentRow-1]
		m[currentRow-1][currentRow-1] = 1.0

		// Dividing by mi
		div, err := m[currentRow-1].DivideByScalar(mi)

		if err != nil {
			return nil, err
		}

		m[currentRow-1] = div

		for i := 1; i <= rows; i++ {
			if i != currentRow {
				mi = m[i-1][currentRow-1]
				m[i-1][currentRow-1] = 0.0
				for j := 1; j <= rows; j++ {
					m[i-1][j-1] -= mi * m[currentRow-1][j-1]
				}
			}
		}
	}

	// Reverse swapping
	for j := rows; j >= 1; j-- {
		p := vec[j-1]
		if p != float64(j) {
			for i := 1; i <= rows; i++ {
				tmp := m[i-1][int64(p)-1]
				m[i-1][int64(p)-1] = m[i-1][j-1]
				m[i-1][j-1] = tmp
			}
		}
	}
	return m, nil
}

// Log applies natural logarithm to all the elements of the matrix, and returns the resulting matrix.
func (m Matrix) Log() Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] = math.Log(m[i][j])
		}
	}
	return m
}

// Log applies e^x to all the elements of the matrix, and returns the resulting matrix.
func (m Matrix) Exp() Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] = math.Exp(m[i][j])
		}
	}
	return m
}

// LeftDivide receives another matrix as a parameter. The method solves the symbolic system of linear equations in matrix form, A*X = B for X. It returns the results in matrix form and error (if there is any).
func (m Matrix) LeftDivide(m2 Matrix) (Matrix, error) {
	var r Matrix
	mT, err := m.Transpose()

	if err != nil {
		return r, err
	}

	mtm, err := mT.MultiplyBy(m)

	if err != nil {
		return r, err
	}

	pInv, err := mtm.Invert()
	if err != nil {
		return r, err
	}

	pmt, err := pInv.MultiplyBy(mT)

	if err != nil {
		return r, err
	}

	r, err = pmt.MultiplyBy(m2)

	if err != nil {
		return r, err
	}

	return r, nil
}

func (m Matrix) sumAbs() float64 {
	var sum float64
	rows, cols := m.Dim()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			sum += math.Abs(m[i][j])
		}
	}
	return sum
}

func (m Matrix) isSquare() bool {
	rows, cols := m.Dim()
	return rows == cols
}

// MultiplyBy receives another matrix as a parameter. It multiplies the matrices and returns the resulting matrix and error.
func (m Matrix) MultiplyBy(m2 Matrix) (Matrix, error) {
	var r Matrix

	_, cols1 := m.Dim()
	rows2, _ := m2.Dim()

	if cols1 != rows2 {
		return r, fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix")
	}

	for currentRowIndex := range m {
		r = append(r, Vector{})
		for currentColumnIndex := range m2[0] {
			m2Col, err := m2.GetColumnAt(currentColumnIndex)

			if err != nil {
				return r, err
			}

			dot, err := m[currentRowIndex].Dot(m2Col)
			if err != nil {
				return r, err
			}

			r[currentRowIndex] = append(r[currentRowIndex], dot)
		}
	}

	return r, nil
}

// AddColumnAt receives the index and the vector. It adds the provided vector as a column at index k, and returns the resulting matrix and the error (if there is any).
func (m Matrix) AddColumnAt(k int, c Vector) (Matrix, error) {
	var r Matrix

	if k < 0 {
		return r, fmt.Errorf("Index cannot be less than 0")
	} else if _, width := m.Dim(); k > width {
		return r, fmt.Errorf("Index cannot be greater than number of columns + 1")
	} else if len(c) != len(m) {
		return r, fmt.Errorf("Column dimensions must match")
	}

	for i := 0; i < len(m); i++ {
		row := m[i]
		expRow := append(row[:k], append(Vector{c[i]}, row[k:]...)...)
		r = append(r, expRow)
	}

	return r, nil
}

// GetRowAt receives the index as a parameter. It returns the vector row at provided index and the error (if there is any).
func (m Matrix) GetRowAt(i int) (Vector, error) {
	if i < 0 {
		return nil, fmt.Errorf("Index cannot be negative")
	} else if i > len(m) {
		return nil, fmt.Errorf("Index cannot be greater than the length")
	}
	return m[i], nil
}

// GetColumnAt receives the index as a parameter. It returns the vector column at provided index and the error (if there is any).
func (m Matrix) GetColumnAt(i int) (Vector, error) {
	var r Vector

	if i < 0 {
		return nil, fmt.Errorf("Index cannot be negative")
	} else if i > len(m[0]) {
		return nil, fmt.Errorf("Index cannot be greater than the length")
	}

	for row := range m {
		r = append(r, m[row][i])
	}

	return r, nil
}

// Transpose returns the transposed matrix and the error.
func (m Matrix) Transpose() (Matrix, error) {
	var t Matrix

	for columnIndex := range m[0] {
		column, err := m.GetColumnAt(columnIndex)
		if err != nil {
			return t, err
		}
		t = append(t, column)
	}

	return t, nil
}

// IsSimilar receives another matrix and tolerance as the parameters. It checks whether the two matrices are similar within the provided tolerance.
func (m Matrix) IsSimilar(m2 Matrix, tol float64) bool {

	if m.IsEqual(m2) {
		return true
	}

	if !m.areDimsEqual(m2) {
		return false
	}

	for col := range m {
		for row := range m[col] {
			if math.Abs(m[col][row]-m2[col][row]) > tol {
				return false
			}
		}
	}

	return true
}

// IsEqual receives another matrix as a parameter. It returns true if the values of the two matrices are equal, and false otherwise.
func (m Matrix) IsEqual(m2 Matrix) bool {
	if m == nil && m2 == nil {
		return true
	} else if m == nil || m2 == nil {
		return false
	} else if !m.areDimsEqual(m2) {
		return false
	}

	for row := range m {
		for col := range m[row] {
			if m[row][col] != m2[row][col] {
				return false
			}
		}
	}
	return true
}

// Add receives another matrix as a parameter. It adds the two matrices and returns the result matrix and the error (if there is any).
func (m Matrix) Add(m2 Matrix) (Matrix, error) {
	rows, cols := m.Dim()
	var r = make(Matrix, rows, cols)
	if ok, err := m.canPerformOperationsWith(m2); !ok {
		return nil, err
	}

	for row := range m {
		for col := range m[row] {
			r[row] = append(r[row], m[row][col]+m2[row][col])
		}
	}

	return r, nil
}

// Subtract receives another matrix as a parameter. It subtracts the two matrices and returns the result matrix and the error (if there is any).
func (m Matrix) Subtract(m2 Matrix) (Matrix, error) {
	rows, cols := m.Dim()
	var r = make(Matrix, rows, cols)
	if ok, err := m.canPerformOperationsWith(m2); !ok {
		return nil, err
	}

	for row := range m {
		for col := range m[row] {
			r[row] = append(r[row], m[row][col]-m2[row][col])
		}
	}

	return r, nil
}

func (m Matrix) areDimsEqual(m2 Matrix) bool {
	mRows, mCols := m.Dim()
	m2Rows, m2Cols := m2.Dim()

	if mRows != m2Rows || mCols != m2Cols {
		return false
	}
	return true
}

func (m Matrix) isNil() bool {
	if m == nil {
		return true
	}
	return false
}

func (m Matrix) canPerformOperationsWith(m2 Matrix) (bool, error) {
	if m == nil || m2 == nil {
		return false, fmt.Errorf("Matrices cannot be nil")
	} else if !m.areDimsEqual(m2) {
		return false, fmt.Errorf("Matrix dimensions must match")
	}
	return true, nil
}

// OTHER CONVENIENT METHODS THAT CAN BE IMPLEMENTED:
// TODO: AddRowAt
// TODO: RemoveRowAt
// TODO: RemoveColumnAt
// TODO: Determinant
// TODO: Check for consistent dimensions
// TODO: IsSingular
