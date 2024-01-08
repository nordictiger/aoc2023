package main

import (
	"fmt"
	"math/big"
)

func matrixRowXY(h1, h2 line) []float64 {
	row := make([]float64, 7)
	row[0] = h2.delta.y - h1.delta.y
	row[1] = h1.point.y - h2.point.y
	row[2] = h1.delta.x - h2.delta.x
	row[3] = h2.point.x - h1.point.x
	row[6] = h2.point.x*h2.delta.y -
		h2.point.y*h2.delta.x -
		h1.point.x*h1.delta.y +
		h1.point.y*h1.delta.x
	return row
}

func matrixRowXZ(h1, h2 line) []float64 {
	row := make([]float64, 7)
	row[0] = h2.delta.z - h1.delta.z
	row[1] = h1.point.z - h2.point.z
	row[4] = h1.delta.x - h2.delta.x
	row[5] = h2.point.x - h1.point.x
	row[6] = h2.point.x*h2.delta.z -
		h2.point.z*h2.delta.x -
		h1.point.x*h1.delta.z +
		h1.point.z*h1.delta.x
	return row
}

func matrixRowYZ(h1, h2 line) []float64 {
	row := make([]float64, 7)
	row[2] = h2.delta.z - h1.delta.z
	row[3] = h1.point.z - h2.point.z
	row[4] = h1.delta.y - h2.delta.y
	row[5] = h2.point.y - h1.point.y
	row[6] = h2.point.y*h2.delta.z -
		h2.point.z*h2.delta.y -
		h1.point.y*h1.delta.z +
		h1.point.z*h1.delta.y
	return row
}

func setRat(s string) *big.Rat {
	r := new(big.Rat)
	r.SetString(s)
	return r
}

func ratFromFloatMatrix(matrix [][]float64) [][]*big.Rat {
	ratMatrix := make([][]*big.Rat, 0)
	for _, row := range matrix {
		ratRow := make([]*big.Rat, 0)
		for _, v := range row {
			floatString := fmt.Sprintf("%.0f", v)
			ratRow = append(ratRow, setRat(floatString))
		}
		ratMatrix = append(ratMatrix, ratRow)
	}
	return ratMatrix
}

func gaussianEliminationRat(a [][]*big.Rat) []*big.Rat {
	rows := len(a)
	cols := len(a[0])

	for i := 0; i < rows; i++ {
		maxRow := i
		for k := i + 1; k < rows; k++ {
			absRat := big.NewRat(0, 1).Abs(a[maxRow][i])
			if a[k][i].Cmp(absRat) > 0 {
				maxRow = k
			}
		}
		a[i], a[maxRow] = a[maxRow], a[i]
		for k := i + 1; k < rows; k++ {
			f := new(big.Rat).Quo(a[k][i], a[i][i])
			for j := i; j < cols; j++ {
				if i == j {
					a[k][j] = new(big.Rat)
				} else {
					a[k][j].Sub(a[k][j], new(big.Rat).Mul(f, a[i][j]))
				}
			}
		}
	}
	x := make([]*big.Rat, rows)
	for i := rows - 1; i >= 0; i-- {
		x[i] = new(big.Rat).Set(a[i][cols-1])
		for j := i + 1; j < rows; j++ {
			x[i].Sub(x[i], new(big.Rat).Mul(a[i][j], x[j]))
		}
		x[i].Quo(x[i], a[i][i])
	}
	return x
}
func getSolution(v []*big.Rat) string {
	sum := new(big.Rat)
	sum.Add(sum, v[0])
	sum.Add(sum, v[2])
	sum.Add(sum, v[4])

	num := sum.Num()
	denom := sum.Denom()
	res := new(big.Int).Div(num, denom)
	return res.String()
}

func puzzle2(lines []line) string {
	h1 := lines[0]
	h2 := lines[1]
	h3 := lines[2]

	matrix := make([][]float64, 0)
	matrix = append(matrix, matrixRowXY(h1, h2))
	matrix = append(matrix, matrixRowXZ(h1, h2))
	matrix = append(matrix, matrixRowYZ(h1, h2))
	matrix = append(matrix, matrixRowXY(h1, h3))
	matrix = append(matrix, matrixRowXZ(h1, h3))
	matrix = append(matrix, matrixRowYZ(h1, h3))

	ratMatrix := ratFromFloatMatrix(matrix)
	solutionRat := gaussianEliminationRat(ratMatrix)
	return getSolution(solutionRat)
}
