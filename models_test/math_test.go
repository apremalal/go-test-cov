package models_test

import (
	"testing"
  "github.com/apremalal/go-test-cov/models"
)

func TestAdd(t *testing.T) {
    total := models.Add(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
