package array

import (
	"fmt"
	"math/rand"
)

type ArrayItem struct {
	Data interface{}
}

type Array struct {
	Items []ArrayItem
}

const LastElement = -9223372036854775808

// NewArray return new Array
func NewArray() *Array {
	return &Array{}
}

// MakeArray return new Array with give data
func MakeArray(data ...interface{}) *Array {
	r := NewArray()
	for _, v := range data {
		r.Push(v)
	}
	return r
}

// MakeNArray return new n element Array with nil filled data
func MakeNArray(n int) *Array {
	r := NewArray()
	for i := 0; i < n; i++ {
		r.Push(nil)
	}
	return r
}

// Push just push to end items
func (a *Array) Push(items ...interface{}) *Array {
	for _, v := range items {
		a.Items = append(a.Items, ArrayItem{Data: v})
	}
	return a
}

// Unshift just push to start items
func (a *Array) Unshift(items ...interface{}) *Array {
	for _, v := range items {
		a.Items = append([]ArrayItem{{Data: v}}, a.Items...)
	}
	return a
}

func slice(a *Array, start, end int) []ArrayItem {
	e := end
	s := start
	l := len(a.Items)

	if start == LastElement {
		s = l - 1
	} else if start < 0 {
		s = l + start
	}

	if end == LastElement {
		e = l - 1
	} else if end < 0 {
		e = l + end
	}

	if e < s {
		e = s
	}

	return a.Items[s:e]
}

// NewSlice return new slice between start & end
// 	if start/end < 0, then count from end
// 	if start/end = -0, then equal to array length
func (a *Array) NewSlice(start, end int) *Array {
	items := slice(a, start, end)
	return &Array{Items: items}
}

// Slice make current array to slice between start & end
// 	if start/end < 0, then count from end
// 	if start/end = -0, then equal to array length
func (a *Array) Slice(start, end int) *Array {
	a.Items = slice(a, start, end)
	return a
}

// Pop return&remove last element
func (a *Array) Pop() ArrayItem {
	i := a.Items[len(a.Items)-1]
	a.Slice(0, -1)
	return i
}

// Shift return&remove first element
func (a *Array) Shift() ArrayItem {
	i := a.Items[0]
	a.Slice(1, len(a.Items))
	return i
}

// Every check is every element equal to data
func (a *Array) Every(callback func(value ArrayItem, index int, array *Array) bool) bool {
	for i, v := range a.Items {
		if !callback(v, i, a) {
			return false
		}
	}
	return true
}

// Some check is have at least one element equal to data
func (a *Array) Some(callback func(value ArrayItem, index int, array *Array) bool) bool {
	for i, v := range a.Items {
		if callback(v, i, a) {
			return true
		}
	}
	return false
}

// Find return finded element searched by callback or empty ArrayItem
func (a *Array) Find(callback func(value ArrayItem, index int, array *Array) bool) ArrayItem {
	for i, v := range a.Items {
		if callback(v, i, a) {
			return v
		}
	}
	return ArrayItem{}
}

// FindIndex return finded element index searched by callback or -1
func (a *Array) FindIndex(callback func(value ArrayItem, index int, array *Array) bool, fromIndex int) int {
	if fromIndex < 0 {
		return -1
	}
	for i := fromIndex; i < len(a.Items); i++ {
		fmt.Println(callback(a.Items[i], i, a))
		if callback(a.Items[i], i, a) {
			return i - fromIndex
		}
	}
	return -1
}

// Includes check is have at least one element equal to data
func (a *Array) Includes(data interface{}, fromIndex int) bool {
	if fromIndex < 0 {
		return false
	}
	for i := fromIndex; i < len(a.Items); i++ {
		if a.Items[i].Data == data {
			return true
		}
	}
	return false
}

// Fill fill all element equal to data
func (a *Array) Fill(data interface{}) *Array {
	for i := range a.Items {
		a.Items[i].Data = data
	}
	return a
}

// Join return joined by separator string
func (a *Array) Join(separator string) string {
	r := ""
	if separator == "" {
		separator = ","
	}
	for _, v := range a.Items {
		r += fmt.Sprint(v.Data) + separator
	}
	r = r[0 : len(r)-1]
	return r
}

// IndexOf return finding element index or -1
func (a *Array) IndexOf(data interface{}, fromIndex int) int {
	if fromIndex < 0 {
		return -1
	}
	for i := fromIndex; i < len(a.Items); i++ {
		if a.Items[i].Data == data {
			return i - fromIndex
		}
	}
	return -1
}

// LastIndexOf return last finding element index or -1
func (a *Array) LastIndexOf(data interface{}, fromIndex int) int {
	if fromIndex < 0 {
		return -1
	}
	for i := len(a.Items) - 1 - fromIndex; i >= 0; i-- {
		if a.Items[i].Data == data {
			return i
		}
	}
	return -1
}

// Reverse return reversed array
func (a *Array) Reverse() *Array {
	for i, j := 0, len(a.Items)-1; i < j; i, j = i+1, j-1 {
		a.Items[i], a.Items[j] = a.Items[j], a.Items[i]
	}
	return a
}

// Filter return new filtered array; remove elements not equal in callback
func (a *Array) Filter(callback func(value ArrayItem, index int, array *Array) bool) *Array {
	arr := NewArray()
	for i, v := range a.Items {
		if callback(v, i, a) {
			arr.Push(v.Data)
		}
	}
	return arr
}

// Map return new array; elements maked in callback
func (a *Array) Map(callback func(value ArrayItem, index int, array *Array) ArrayItem) *Array {
	arr := NewArray()
	for i, v := range a.Items {
		arr.Push(callback(v, i, a).Data)
	}
	return arr
}

// Reduce return common data for all array; data maked in callback in ascending order
func (a *Array) Reduce(callback func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{}, initValue interface{}) interface{} {
	for i, v := range a.Items {
		initValue = callback(initValue, v, i, a)
	}
	return initValue
}

// ReduceRight return common data for all array; data maked in callback in descending order
func (a *Array) ReduceRight(callback func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{}, initValue interface{}) interface{} {
	for i := len(a.Items) - 1; i >= 0; i-- {
		initValue = callback(initValue, a.Items[i], i, a)
	}
	return initValue
}

func qsort(arr []ArrayItem, compareFunction func(a, b ArrayItem) int) []ArrayItem {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(arr)

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if compareFunction(arr[i], arr[right]) < 0 {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	qsort(arr[:left], compareFunction)
	qsort(arr[left+1:], compareFunction)

	return arr
}

// Sort return sorted array
// 	if compareFunction(a, b) < 0, sort will place a before b.
// 	if compareFunction(a, b) == 0, sort will not change order between a and b, but change order among other element.
// 	if compareFunction(a, b) > 0, sort will place b before a.
func (a *Array) Sort(compareFunction func(a, b ArrayItem) int) *Array {
	a.Items = qsort(a.Items, compareFunction)
	return a
}
