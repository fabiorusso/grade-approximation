package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	input := []int{4, 73, 67, 38, 33}
	x, _ := ApproximateGrades(input)
	log.Printf("%d", x)

}

func ApproximateGrades(grades []int) ([]int, error) {
	if grades == nil {
		return nil, errors.New("valor de entrada esta nulo")
	} else if len(grades) < 2 {
		return nil, errors.New("O numero de notas informados não pode ser vazio")
	}

	n := grades[0]
	v := grades[1:]

	if len(v) != n {
		return nil, errors.New("numero de estudantes diferente do numero de notas")
	}

	approximatedGrades := make([]int, n)

	for index := 0; index < n; index++ {
		grade, err := ApproximateGrade(v[index])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("erro ao tentar aproximar a nota para o estudante %d: %s ", index+1, err.Error()))
		}

		approximatedGrades[index] = grade
	}

	return approximatedGrades, nil
}

func CheckGradeBoundaries(grade int) (bool, error) {
	if grade < 0 || grade > 100 {
		return false, errors.New(fmt.Sprintf("a nota deve ser menor que 100 e maior que zero: o valor informado %d não pertence a esse intervalo", grade))
	}
	return true, nil
}

func ApproximateGrade(grade int) (int, error) {
	isInbounds, err := CheckGradeBoundaries(grade)
	if err != nil && !isInbounds {
		return -1, err
	}
	if grade >= 38 {
		gap := grade % 5
		if gap >= 3 {
			return grade + (5 - gap), nil
		} else {
			return grade, nil
		}
	} else {
		return grade, nil
	}
}
