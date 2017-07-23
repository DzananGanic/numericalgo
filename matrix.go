package numericalgo

import "fmt"

type Matrix []Vector

func (m Matrix) Dim() (int, int) {
	if m.isNil() {
		return 0, 0
	}
	return len(m[0]), len(m)
}

// PRIORITY:
// TODO: Multiplication
// TODO: Inverse

// OTHER METHODS:
// TODO: AddRowAt
// TODO: RemoveRowAt
// TODO: RemoveColumnAt
// TODO: Determinant
// TODO: Check for consistent dimensions
// TODO: IsSingular

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
