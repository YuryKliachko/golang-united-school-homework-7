package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestPeopleLen(t *testing.T) {
	people := People{{firstName: "Yury", lastName: "Kliachko", birthDay: time.Now()}}
	actualLen := people.Len()
	expectedLen := 1
	if actualLen != expectedLen {
		t.Errorf("Actual length %d does not match expected %d", actualLen, expectedLen)
	}
}

type LessTestInput struct {
	people         People
	expectedResult bool
}

func TestPeopleLess(t *testing.T) {
	testData := map[string]LessTestInput{
		"Different people": {
			people: People{
				{
					firstName: "Gary",
					lastName:  "Weber",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Michael",
					lastName:  "Tompson",
					birthDay:  time.Date(2008, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: true,
		},
		"Birthday equal": {
			people: People{
				{
					firstName: "Gary",
					lastName:  "Weber",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Michael",
					lastName:  "Tompson",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: true,
		},
		"Birthday and first name equal": {
			people: People{
				{
					firstName: "Michael",
					lastName:  "Weber",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Michael",
					lastName:  "Tompson",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: false,
		},
		"Birthday, first and last names equal": {
			people: People{
				{
					firstName: "Michael",
					lastName:  "Tompson",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Michael",
					lastName:  "Tompson",
					birthDay:  time.Date(2009, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: false,
		},
	}
	for testName, testInput := range testData {
		acutalResult := testInput.people.Less(0, 1)
		if acutalResult != testInput.expectedResult {
			t.Errorf("%s: Actual result %t does not match expected %t", testName, acutalResult, testInput.expectedResult)
		}
	}

}

func TestPeopleSwap(t *testing.T) {
	firstPerson := Person{firstName: "Yury", lastName: "Kliachko", birthDay: time.Now()}
	secondPerson := Person{firstName: "Test", lastName: "Tester", birthDay: time.Now()}
	people := People{firstPerson, secondPerson}
	people.Swap(0, 1)
	if people[0].firstName != "Test" || people[1].firstName != "Yury" {
		t.Errorf("Wrong people list after swapping %v", people)
	}
}

type MatrixTestInput struct {
	matrixString   string
	expectedMatrix Matrix
	expectedErr    string
}

func TestMatrix(t *testing.T) {
	testData := map[string]MatrixTestInput{
		"Valid matrix": {
			matrixString:   "1 2 3\n4 5 6",
			expectedMatrix: Matrix{rows: 2, cols: 3, data: []int{1, 2, 3, 4, 5, 6}},
			expectedErr:    "",
		},
		"Invalid matrix": {
			matrixString: "1 2 3\n1",
			expectedErr:  "Rows need to be the same length",
		},
		"Invalid elements in matrix": {
			matrixString: "1 2 3\n3 2 b",
			expectedErr:  "strconv.Atoi: parsing \"b\": invalid syntax",
		},
	}
	for testName, testInput := range testData {
		actualMatrix, actualErr := New(testInput.matrixString)
		if testInput.expectedErr != "" {
			if actualErr.Error() != testInput.expectedErr {
				t.Errorf("%s: actual error '%s' does not match expected '%s'", testName, actualErr, testInput.expectedErr)
			}
		} else {
			if actualMatrix.rows != testInput.expectedMatrix.rows || actualMatrix.cols != testInput.expectedMatrix.cols {
				t.Errorf("%s: actual matrix %v does not match expected matrix %v", testName, actualMatrix, testInput.expectedMatrix)
			}
		}
	}
}

func TestMatrixRows(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6")
	expectedRows := [][]int{{1, 2, 3}, {4, 5, 6}}
	acutalRows := matrix.Rows()
	if reflect.DeepEqual(acutalRows, expectedRows) == false {
		t.Errorf("Actual rows %v do not match expected rows %v", acutalRows, expectedRows)
	}
}

func TestMatrixCols(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6")
	expectedCols := [][]int{{1, 4}, {2, 5}, {3, 6}}
	acutalCols := matrix.Cols()
	if reflect.DeepEqual(acutalCols, expectedCols) == false {
		t.Errorf("Actual rows %v do not match expected rows %v", acutalCols, expectedCols)
	}
}

func TestMatrixSet(t *testing.T) {
	matrix, _ := New("1 2\n3 4")
	matrix.Set(1, 1, 666)
	expectedData := []int{1, 2, 3, 666}
	if reflect.DeepEqual(matrix.data, expectedData) == false {
		t.Errorf("Actual data %v does not match expected data %v", matrix.data, expectedData)
	}
}

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
