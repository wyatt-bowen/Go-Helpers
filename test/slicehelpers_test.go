package test

import (
	"fmt"
	"testing"

	"github.com/wyatt-bowen/go-helpers/filters"
	"github.com/wyatt-bowen/go-helpers/slicehelpers"
)

func TestSliceFilter(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	testSlice = slicehelpers.FilterSlice(testSlice, filters.Equals(5))
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{5}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}

func TestSliceFilterOrEquals(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	filter := filters.Or(filters.Equals(4), filters.Equals(5), filters.Equals(3))
	testSlice = slicehelpers.FilterSlice(testSlice, filter)
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{3, 4, 5}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}

func TestSliceFilterOrEvenEquals(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	filter := filters.Or(filters.Even(), filters.Equals(3))
	testSlice = slicehelpers.FilterSlice(testSlice, filter)
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{2, 3, 4}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}

func TestSliceFilterNotOrEvenEquals(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	filter := filters.Not(filters.Or(filters.Even(), filters.Equals(3)))
	testSlice = slicehelpers.FilterSlice(testSlice, filter)
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{1, 5}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}

func TestReverseSlice(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	slicehelpers.ReverseSlice(testSlice)
	fmt.Println(slicehelpers.ToString(testSlice))
}

func TestMapSquare(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	testSlice = slicehelpers.Map(testSlice, slicehelpers.Square())
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{1, 4, 9, 16, 25}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}

func TestMapInplaceSquare(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(slicehelpers.ToString(testSlice))
	slicehelpers.MapInplace(testSlice, slicehelpers.Square())
	fmt.Println(slicehelpers.ToString(testSlice))
	expected := []int{1, 4, 9, 16, 25}
	if !slicehelpers.SliceEquals(testSlice, expected) {
		t.Fail()
	}
}
