package utils

import (
	"fmt"
	"testing"
)

func TestSplitStringSlice(t *testing.T) {
	actualResult := SplitStringSlice("a,b, c")
	expectResult := []string{"a", "b", "c"}

	if len(actualResult) != len(expectResult) {
		t.Errorf("Not match length: actualResult: %v, expectResult: %v", actualResult, expectResult)
	}

	if fmt.Sprintf("%T", actualResult) != fmt.Sprintf("%T", expectResult) {
		t.Errorf("Not match type: actualResult: %v, expectResult: %v", actualResult, expectResult)
	}

	for i := range actualResult {
		if actualResult[i] != expectResult[i] {
			t.Errorf("Not match value: actualResult[%v]: %v, expectResult[%v]: %v", i, actualResult[i], i, expectResult[i])
		}
	}
}
