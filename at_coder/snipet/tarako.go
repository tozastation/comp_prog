package snipet

type Matrix struct{}

func (m *Matrix) New(nRow, nCol int) interface{} {
	matrix := make([][]int, nCol)
	for i := range matrix {
		matrix[i] = make([]int, nRow)
	}
	return matrix
}
