package primefactors

import (
	"reflect"
	"testing"
)

func list(ints ...int) []int {
	numbers := []int{}
	for _, num := range ints {
		numbers = append(numbers, num)
	}
	return numbers
}

func TestOne(t *testing.T) {
	list := list()
	primes := generate(1)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}

func TestTwo(t *testing.T) {
	list := list(2)
	primes := generate(2)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}

func TestThree(t *testing.T) {
	list := list(3)
	primes := generate(3)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}

func TestFour(t *testing.T) {
	list := list(2, 2)
	primes := generate(4)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}

func TestSix(t *testing.T) {
	list := list(2, 3)
	primes := generate(6)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}

func TestEight(t *testing.T) {
	list := list(2, 2, 2)
	primes := generate(8)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}
func TestNine(t *testing.T) {
	list := list(3, 3)
	primes := generate(9)

	if !reflect.DeepEqual(list, primes) {
		t.Errorf("expected %v but was %v", list, primes)
	}
}
