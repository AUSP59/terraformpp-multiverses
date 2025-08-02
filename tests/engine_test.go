
package engine

import "testing"

func TestProcess(t *testing.T) {
    input := "test input"
    expected := "processed: test input"
    got := Process(input)
    if got != expected {
        t.Errorf("Expected %s but got %s", expected, got)
    }
}
