package numericalgo

import "fmt"
import "math"

type Matrix []Vector

func (m Matrix) Dim() (int, int) {
	if m.isNil() {
		return 0, 0
	}
	return len(m[0]), len(m)
}

// OTHER METHODS:
// TODO: AddRowAt
// TODO: RemoveRowAt
// TODO: RemoveColumnAt
// TODO: Determinant
// TODO: Check for consistent dimensions
// TODO: IsSingular

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
		pivot := currentRow
		for i := currentRow + 1; i <= rows; i++ {
			if math.Abs(m[i-1][currentRow-1]) > math.Abs(m[pivot-1][currentRow-1]) {
				pivot = i
			}
		}

		// If there exists no element a(k,i) different from zero, matrix is singular and has none or more than one solution
		if math.Abs(m[pivot-1][currentRow-1]) < 1e-10 {
			return nil, fmt.Errorf("Matrix is singular")
		}

		// If we find pivot which is the largest a(i, currentRow), we swap the rows
		if pivot != currentRow {
			tmp := m[currentRow-1]
			m[currentRow-1] = m[pivot-1]
			m[pivot-1] = tmp
		}

		vec[currentRow-1] = float64(pivot)

		mi := m[currentRow-1][currentRow-1]
		m[currentRow-1][currentRow-1] = 1.0

		// Dividing by mi
		divided, err := m[currentRow-1].DivideByScalar(mi)

		if err != nil {
			return nil, err
		}

		m[currentRow-1] = divided

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
		pivot := vec[j-1]
		if pivot != float64(j) {
			for i := 1; i <= rows; i++ {
				tmp := m[i-1][int64(pivot)-1]
				m[i-1][int64(pivot)-1] = m[i-1][j-1]
				m[i-1][j-1] = tmp
			}
		}
	}
	return m, nil
}

func (m Matrix) LeftDivide(m2 Matrix) (Matrix, error) {
	var result Matrix
	mTransposed, err := m.Transpose()

	if err != nil {
		return result, err
	}

	mtm, err := mTransposed.MultiplyBy(m)

	if err != nil {
		return result, err
	}

	pseudoInverse, err := mtm.Invert()

	if err != nil {
		return result, err
	}

	pmt, err := pseudoInverse.MultiplyBy(mTransposed)

	if err != nil {
		return result, err
	}

	result, err = pmt.MultiplyBy(m2)

	if err != nil {
		return result, err
	}

	return result, nil
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

func (m Matrix) MultiplyBy(m2 Matrix) (Matrix, error) {
	var result Matrix

	cols1, _ := m.Dim()
	_, rows2 := m2.Dim()

	if cols1 != rows2 {
		return result, fmt.Errorf("The number of columns of the 1st matrix must equal the number of rows of the 2nd matrix")
	}

	for currentRowIndex := range m {
		result = append(result, Vector{})
		for currentColumnIndex := range m2[0] {
			m2Col, err := m2.GetColumnAt(currentColumnIndex)

			if err != nil {
				return result, err
			}

			dot, err := m[currentRowIndex].Dot(m2Col)
			if err != nil {
				return result, err
			}

			result[currentRowIndex] = append(result[currentRowIndex], dot)
		}
	}

	return result, nil
}

func (m Matrix) AddColumnAt(k int, c Vector) (Matrix, error) {
	var result Matrix

	if k < 0 {
		return result, fmt.Errorf("Index cannot be less than 0")
	} else if width, _ := m.Dim(); k > width {
		return result, fmt.Errorf("Index cannot be greater than number of columns + 1")
	} else if len(c) != len(m) {
		return result, fmt.Errorf("Column dimensions must match")
	}

	for i := 0; i < len(m); i++ {
		row := m[i]
		expandedRow := append(row[:k], append(Vector{c[i]}, row[k:]...)...)
		result = append(result, expandedRow)
	}
	return result, nil
}

func (m Matrix) GetRowAt(i int) (Vector, error) {
	if i < 0 {
		return nil, fmt.Errorf("Index cannot be negative")
	} else if i > len(m) {
		return nil, fmt.Errorf("Index cannot be greater than the length")
	}
	return m[i], nil
}

func (m Matrix) GetColumnAt(i int) (Vector, error) {
	var result Vector

	if i < 0 {
		return nil, fmt.Errorf("Index cannot be negative")
	} else if i > len(m[0]) {
		return nil, fmt.Errorf("Index cannot be greater than the length")
	}

	for row := range m {
		result = append(result, m[row][i])
	}
	return result, nil
}

func (m Matrix) Transpose() (Matrix, error) {
	var transposed Matrix

	for columnIndex := range m[0] {
		column, err := m.GetColumnAt(columnIndex)
		if err != nil {
			return transposed, err
		}
		transposed = append(transposed, column)
	}

	return transposed, nil
}

func (m Matrix) IsSimilar(m2 Matrix, tolerance float64) bool {
	if m.IsEqual(m2) {
		return true
	}

	if !m.areDimsEqual(m2) {
		return false
	}

	for column := range m {
		for row := range m[column] {
			if math.Abs(m[column][row]-m2[column][row]) > tolerance {
				return false
			}
		}
	}

	return true
}

func (m Matrix) IsEqual(m2 Matrix) bool {
	if m == nil && m2 == nil {
		return true
	} else if m == nil || m2 == nil {
		return false
	} else if !m.areDimsEqual(m2) {
		return false
	}

	for column := range m {
		for row := range m[column] {
			if m[column][row] != m2[column][row] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Add(m2 Matrix) (Matrix, error) {
	width, height := m.Dim()
	var result = make(Matrix, width, height)
	if ok, err := m.canPerformOperationsWith(m2); !ok {
		return nil, err
	}

	for row := range m {
		for column := range m[row] {
			result[row] = append(result[row], m[row][column]+m2[row][column])
		}
	}

	return result, nil
}

func (m Matrix) Subtract(m2 Matrix) (Matrix, error) {
	width, height := m.Dim()
	var result = make(Matrix, width, height)
	if ok, err := m.canPerformOperationsWith(m2); !ok {
		return nil, err
	}

	for row := range m {
		for column := range m[row] {
			result[row] = append(result[row], m[row][column]-m2[row][column])
		}
	}

	return result, nil
}

func (m Matrix) areDimsEqual(m2 Matrix) bool {
	mRows, mColumns := m.Dim()
	m2Rows, m2Columns := m2.Dim()

	if mRows != m2Rows || mColumns != m2Columns {
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
