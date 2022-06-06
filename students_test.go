package coverage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

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

func TestPersonLen(t *testing.T) {
	t.Parallel()
	testData := map[string]People{
		"nil":           nil,
		"empty len":     make(People, 0),
		"not empty len": make(People, 200)}

	for name, value := range testData {
		tests := value
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, len(tests), tests.Len())
		})
	}
}

func TestPersonLess(t *testing.T) {
	t.Parallel()
	tests := People{
		{
			"Andrey",
			"Aewsa",
			time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Alake",
			"Asena",
			time.Date(2000, time.January, 2, 1, 1, 1, 0, time.UTC),
		},
		{
			"Alake",
			"Akilart",
			time.Date(2000, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Vrat",
			"Bekas",
			time.Date(2000, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Alake",
			"Vasert",
			time.Date(2000, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Bndrey",
			"Aewsa",
			time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
	}

	testData := map[string]struct {
		tests    People
		i        int
		j        int
		expected bool
	}{
		"older":                           {tests, 0, 1, true},
		"younger":                         {tests, 1, 0, false},
		"first name win, birthday equal":  {tests, 0, 5, true},
		"first name lose, birthday equal": {tests, 5, 0, false},
		"last name win, birthday equal and firstname equal":  {tests, 2, 4, true},
		"last name lose, birthday equal and firstname equal": {tests, 4, 2, false},
	}
	for nameTest, value := range testData {
		tests := value.tests
		i := value.i
		j := value.j
		expected := value.expected
		t.Run(nameTest, func(t *testing.T) {
			assert.Equal(t, expected, tests.Less(i, j))

		})

	}
}

func TestPersonSwap(t *testing.T) {
	t.Parallel()
	tests1 := People{
		{
			"Andrey",
			"Aewsa",
			time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Alake",
			"Asena",
			time.Date(2000, time.January, 2, 1, 1, 1, 0, time.UTC),
		},
	}
	tests2 := People{
		{
			"Andrey",
			"Aewsa",
			time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			"Alake",
			"Asena",
			time.Date(2000, time.January, 2, 1, 1, 1, 0, time.UTC),
		},
	}
	t.Run("SwapTest", func(t *testing.T) {
		tests1.Swap(0, 1)
		assert.Equal(t, tests2[0], tests1[1])
	})

}

func TestMatrixNew(t *testing.T) {
	t.Parallel()
	testData1notNil := map[string]struct {
		input     string
		valueRows int
		valueCols int
	}{
		"input so simple 1":                     {"1\n1", 2, 1},
		"input so simple 2":                     {"1 1 1 1", 1, 4},
		"input is 2x2":                          {"1 2\n1 3", 2, 2},
		"input with first and last whitespaces": {" 1 2 3\n1 2 3 ", 2, 3},
	}
	testData2isNil := map[string]struct {
		input     string
		valueRows int
		valueCols int
	}{
		"input is nil 1": {"", 0, 0},
		"input is nil 2": {"1 2 3 4 5 6\n 1", 2, 6},
	}
	for nameTest, value := range testData1notNil {
		tests := value.input
		valueRows := value.valueRows
		valueCols := value.valueCols
		t.Run(nameTest, func(t *testing.T) {
			testValue, err := New(tests)
			assert.Equal(t, valueRows, testValue.rows)
			assert.Equal(t, valueCols, testValue.cols)
			assert.Equal(t, nil, err)

		})

	}
	for nameTest, value := range testData2isNil {
		tests := value.input
		t.Run(nameTest, func(t *testing.T) {
			testValue, err := New(tests)
			assert.Nil(t, testValue)
			assert.NotNil(t, err)
		})
	}

}
func TestMatrixRows(t *testing.T) {
	t.Parallel()
	testData := map[string]struct {
		input    string
		expected [][]int
	}{
		"2x2 matrix": {"1 2\n1 2", [][]int{{1, 2}, {1, 2}}},
		"3x3 matrix": {"1 2 3\n1 2 3\n1 2 3", [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}},
		"2x1 matrix": {"1\n1", [][]int{{1}, {1}}},
		"1x2 matrix": {"1 2", [][]int{{1, 2}}},
	}
	for nameTest, value := range testData {
		tests := value.input
		expected := value.expected
		t.Run(nameTest, func(t *testing.T) {
			testValue, err := New(tests)
			assert.Equal(t, expected, testValue.Rows())
			assert.Equal(t, nil, err)
		})
	}
}
func TestMatrixCols(t *testing.T) {
	t.Parallel()
	testData := map[string]struct {
		input    string
		expected [][]int
	}{
		"2x2 matrix": {"1 2\n1 2", [][]int{{1, 1}, {2, 2}}},
		"3x3 matrix": {"1 2 3\n1 2 3\n1 2 3", [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}},
		"2x1 matrix": {"1\n1", [][]int{{1, 1}}},
		"1x2 matrix": {"1 2", [][]int{{1}, {2}}},
	}
	for nameTest, value := range testData {
		tests := value.input
		expected := value.expected
		t.Run(nameTest, func(t *testing.T) {
			testValue, err := New(tests)
			assert.Equal(t, expected, testValue.Cols())
			assert.Equal(t, nil, err)
		})
	}
}
func TestMatrixSet(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		input           string
		col, row, value int
		expectMatrix    []int
		expectBool      bool
	}{
		"simple string":        {"1\n1", 0, 0, 5, []int{5, 1}, true},
		"not simple string":    {"1 2 3 4\n1 2 3 4\n1 2 3 4", 1, 1, 13, []int{1, 2, 3, 4, 1, 13, 3, 4, 1, 2, 3, 4}, true},
		"col,row not positive": {"1 2\n1 2", -1, -1, 5, []int{1, 2, 1, 2}, false},
		"col,row bigger":       {"1 2\n1 2", 4, 4, 5, []int{1, 2, 1, 2}, false},
	}
	for nameTest, tvalue := range testData {
		input := tvalue.input
		col, row, value := tvalue.col, tvalue.row, tvalue.value
		expectMatrix := tvalue.expectMatrix
		expectBool := tvalue.expectBool
		t.Run(nameTest, func(t *testing.T) {
			testValueBeforeSet, _ := New(input)
			valueBool := testValueBeforeSet.Set(row, col, value)

			assert.Equal(t, expectBool, valueBool)
			assert.Equal(t, expectMatrix, testValueBeforeSet.data)

		})
	}
}
