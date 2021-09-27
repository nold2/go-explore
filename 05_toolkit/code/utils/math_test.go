package utils

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 4
	actual := utils.Add(2, 2)

	if actual != expected {
		t.Errorf("Addition is incorrect! Expected %d, Actual %d", expected, actual)
	}

}
