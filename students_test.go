package coverage

import (
	"os"
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
			expectedResult: false,
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
