package Hobbies

import (
	"reflect"
	"testing"
)

func TestPractice(t *testing.T) {
	got := Hobbies()
	want := []string{"Get a job", "Buy a new laptop", "Teach others"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed: got %v, want %v", got, want)
	}
}
