package array

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

type test struct {
	a string
	b int
}

var TestLog = func(testName string, t *testing.T, incoming, got, except, descr interface{}) bool {
	if !reflect.DeepEqual(got, except) {
		t.Errorf("%v:(%v) = %v, want %v. Test: %v\n", testName, incoming, got, except, descr)
		return false
	} else {
		t.Logf("%v:(%v) PASS", testName, descr)
		return true
	}
}

func TestArrayMakeArray(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming: []interface{}{1, 2, 3},
			want: &Array{
				Items: []ArrayItem{
					{Data: 1}, {Data: 2}, {Data: 3},
				},
			},
			description: "int array",
		},
		{
			incoming: []interface{}{"str1", "str2", "str3"},
			want: &Array{
				Items: []ArrayItem{
					{Data: "str1"}, {Data: "str2"}, {Data: "str3"},
				},
			},
			description: "string array",
		},
		{
			incoming: []interface{}{"str1", 2, true},
			want: &Array{
				Items: []ArrayItem{
					{Data: "str1"}, {Data: 2}, {Data: true},
				},
			},
			description: "mixed array",
		},
		{
			incoming: []interface{}{test{"a", 1}, test{"b", -1}},
			want: &Array{
				Items: []ArrayItem{
					{Data: test{"a", 1}}, {Data: test{"b", -1}},
				},
			},
			description: "struct array",
		},
		{
			incoming:    []interface{}{},
			want:        &Array{},
			description: "empty array",
		},
	}

	for _, tt := range tests {
		got := MakeArray(tt.incoming...)

		if !TestLog("MakeArray", t, tt.incoming, got, tt.want, tt.description) {
			continue
		}
	}
}

func TestArrayMakeNArray(t *testing.T) {
	tests := []struct {
		incoming    int
		want        *Array
		description string
	}{
		{
			incoming: 5,
			want: &Array{
				Items: []ArrayItem{
					{Data: nil}, {Data: nil}, {Data: nil}, {Data: nil}, {Data: nil},
				},
			},
			description: "n = 5 array",
		},
		{
			incoming:    -5,
			want:        &Array{},
			description: "n = -5 array",
		},
	}

	for _, tt := range tests {
		got := MakeNArray(tt.incoming)

		if !TestLog("MakeNArray", t, tt.incoming, got, tt.want, tt.description) {
			continue
		}
	}
}

func TestArrayPush(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming: []interface{}{1, 2, 3},
			want: &Array{
				Items: []ArrayItem{
					{Data: 1}, {Data: 2}, {Data: 3},
				},
			},
			description: "int array",
		},
		{
			incoming: []interface{}{"str1", "str2", "str3"},
			want: &Array{
				Items: []ArrayItem{
					{Data: "str1"}, {Data: "str2"}, {Data: "str3"},
				},
			},
			description: "string array",
		},
		{
			incoming: []interface{}{"str1", 2, true},
			want: &Array{
				Items: []ArrayItem{
					{Data: "str1"}, {Data: 2}, {Data: true},
				},
			},
			description: "mixed array",
		},
		{
			incoming: []interface{}{test{"a", 1}, test{"b", -1}},
			want: &Array{
				Items: []ArrayItem{
					{Data: test{"a", 1}}, {Data: test{"b", -1}},
				},
			},
			description: "struct array",
		},
		{
			incoming:    []interface{}{},
			want:        &Array{},
			description: "empty array",
		},
	}

	for _, tt := range tests {
		arr := MakeArray()
		got := arr.Push(tt.incoming...)

		if !TestLog("Push", t, tt.incoming, got, tt.want, tt.description) {
			continue
		}
	}
}

func TestArrayUnshift(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming: []interface{}{1, 2, 3},
			want: &Array{
				Items: []ArrayItem{
					{Data: 3}, {Data: 2}, {Data: 1},
				},
			},
			description: "int array",
		},
		{
			incoming: []interface{}{"str1", "str2", "str3"},
			want: &Array{
				Items: []ArrayItem{
					{Data: "str3"}, {Data: "str2"}, {Data: "str1"},
				},
			},
			description: "string array",
		},
		{
			incoming: []interface{}{"str1", 2, true},
			want: &Array{
				Items: []ArrayItem{
					{Data: true}, {Data: 2}, {Data: "str1"},
				},
			},
			description: "mixed array",
		},
		{
			incoming: []interface{}{test{"a", 1}, test{"b", -1}},
			want: &Array{
				Items: []ArrayItem{
					{Data: test{"b", -1}}, {Data: test{"a", 1}},
				},
			},
			description: "struct array",
		},
		{
			incoming:    []interface{}{},
			want:        &Array{},
			description: "empty array",
		},
	}

	for _, tt := range tests {
		arr := MakeArray()
		got := arr.Unshift(tt.incoming...)

		if !TestLog("Unshift", t, tt.incoming, got, tt.want, tt.description) {
			continue
		}
	}
}

func TestArraySlice(t *testing.T) {
	tests := []struct {
		incoming    []int
		want        *Array
		description string
	}{
		{
			incoming: []int{0, 2},
			want: &Array{
				Items: []ArrayItem{
					{Data: 1}, {Data: 2},
				},
			},
			description: "[0:2] slice",
		},
		{
			incoming: []int{2, 4},
			want: &Array{
				Items: []ArrayItem{
					{Data: 3}, {Data: 4},
				},
			},
			description: "[2:4] slice",
		},
		{
			incoming: []int{0, -2},
			want: &Array{
				Items: []ArrayItem{
					{Data: 1}, {Data: 2}, {Data: 3}, {Data: 4}, {Data: 5},
				},
			},
			description: "[0:-2] slice",
		},
		{
			incoming: []int{-4, -2},
			want: &Array{
				Items: []ArrayItem{
					{Data: 4}, {Data: 5},
				},
			},
			description: "[-4:-2] slice",
		},
		{
			incoming: []int{0, LastElement},
			want: &Array{
				Items: []ArrayItem{
					{Data: 1}, {Data: 2}, {Data: 3}, {Data: 4}, {Data: 5}, {Data: 6},
				},
			},
			description: "[0:LastElement] slice",
		},
	}

	for _, tt := range tests {
		arr := MakeArray([]interface{}{1, 2, 3, 4, 5, 6, 7}...)
		got := arr.Slice(tt.incoming[0], tt.incoming[1])

		if !TestLog("Slice", t, tt.incoming, got, tt.want, tt.description) {
			continue
		}
	}
}

func TestArrayPop(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			description: "int array",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			description: "string array",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			description: "mixed array",
		},
		{
			incoming:    []interface{}{test{"a", 1}, test{"b", -1}},
			description: "struct array",
		},
		{
			incoming:    []interface{}{},
			description: "empty array",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		for i := len(tt.incoming) - 1; i >= 0; i-- {
			got := arr.Pop()
			if !TestLog("Pop", t, tt.incoming[i], got.Data, tt.incoming[i], tt.description) {
				break
			}
		}
	}
}

func TestArrayShift(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			description: "int array",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			description: "string array",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			description: "mixed array",
		},
		{
			incoming:    []interface{}{test{"a", 1}, test{"b", -1}},
			description: "struct array",
		},
		{
			incoming:    []interface{}{},
			description: "empty array",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		for i := 0; i < len(tt.incoming); i++ {
			got := arr.Shift()
			if !TestLog("Shift", t, tt.incoming[i], got.Data, tt.incoming[i], tt.description) {
				break
			}
		}
	}
}

func TestArrayEvery(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        bool
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        true,
			description: ">0",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        false,
			description: "=str1",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			want:        true,
			description: "!=&Array{}",
		},
	}

	// test1
	test := tests[0]
	arr := MakeArray(test.incoming...)
	got := arr.Every(func(value ArrayItem, index int, array *Array) bool {
		return value.Data.(int) > 0
	})
	TestLog("Every", t, test.incoming, got, test.want, test.description)

	// test2
	test = tests[1]
	arr = MakeArray(test.incoming...)
	got = arr.Every(func(value ArrayItem, index int, array *Array) bool {
		return value.Data.(string) == "str1"
	})
	TestLog("Every", t, test.incoming, got, test.want, test.description)

	// test3
	test = tests[2]
	arr = MakeArray(test.incoming...)
	got = arr.Every(func(value ArrayItem, index int, array *Array) bool {
		return value.Data != &Array{}
	})
	TestLog("Every", t, test.incoming, got, test.want, test.description)
}

func TestArraySome(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        bool
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        true,
			description: ">2",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        true,
			description: "=str1",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			want:        false,
			description: "=&Array{}",
		},
	}

	// test1
	test := tests[0]
	arr := MakeArray(test.incoming...)
	got := arr.Some(func(value ArrayItem, index int, array *Array) bool {
		return value.Data.(int) > 2
	})
	TestLog("Some", t, test.incoming, got, test.want, test.description)

	// test2
	test = tests[1]
	arr = MakeArray(test.incoming...)
	got = arr.Some(func(value ArrayItem, index int, array *Array) bool {
		return value.Data.(string) == "str1"
	})
	TestLog("Some", t, test.incoming, got, test.want, test.description)

	// test3
	test = tests[2]
	arr = MakeArray(test.incoming...)
	got = arr.Some(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == &Array{}
	})
	TestLog("Some", t, test.incoming, got, test.want, test.description)
}

func TestArrayFind(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        ArrayItem
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        ArrayItem{2},
			description: "find int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        ArrayItem{"str1"},
			description: "find string",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			want:        ArrayItem{},
			description: "find ArrayItem",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.Find(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == 2
	})
	TestLog("Find", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.Find(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == "str1"
	})
	TestLog("Find", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.Find(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == ArrayItem{}
	})
	TestLog("Find", t, tt.incoming, got, tt.want, tt.description)
}

func TestArrayFindIndex(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		fromIndex   int
		want        int
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        1,
			description: "findIndex int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        0,
			description: "findIndex string",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			want:        -1,
			fromIndex:   -1,
			description: "findIndex ArrayItem",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.FindIndex(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == 2
	}, tt.fromIndex)
	TestLog("FindIndex", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.FindIndex(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == "str1"
	}, tt.fromIndex)
	TestLog("FindIndex", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.FindIndex(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == ArrayItem{}
	}, tt.fromIndex)
	TestLog("FindIndex", t, tt.incoming, got, tt.want, tt.description)
}

func TestArrayIncludes(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		desired     interface{}
		fromIndex   int
		want        bool
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			want:        true,
			description: "includes int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			desired:     "str1",
			want:        true,
			description: "includes string",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			desired:     ArrayItem{Data: ArrayItem{}},
			want:        false,
			description: "includes ArrayItem",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   1,
			want:        true,
			description: "includes from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   2,
			want:        false,
			description: "unincludes from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   5,
			want:        false,
			description: "unincludes from index out range",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		got := arr.Includes(tt.desired, tt.fromIndex)
		TestLog("Includes", t, tt.incoming, got, tt.want, tt.description)
	}
}

func TestArrayFill(t *testing.T) {
	tests := []struct {
		filler      interface{}
		length      int
		want        *Array
		description string
	}{
		{
			filler:      2,
			length:      5,
			want:        &Array{[]ArrayItem{{2}, {2}, {2}, {2}, {2}}},
			description: "fill int",
		},
		{
			filler:      "str",
			length:      3,
			want:        &Array{[]ArrayItem{{"str"}, {"str"}, {"str"}}},
			description: "fill string",
		},
		{
			filler:      ArrayItem{2},
			length:      3,
			want:        &Array{[]ArrayItem{{ArrayItem{2}}, {ArrayItem{2}}, {ArrayItem{2}}}},
			description: "fill struct",
		},
	}

	for _, tt := range tests {
		arr := MakeNArray(tt.length)
		got := arr.Fill(tt.filler)
		TestLog("Fill", t, fmt.Sprintf("filler:%v, length:%v", tt.filler, tt.length), got, tt.want, tt.description)
	}
}

func TestArrayJoin(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        string
		separator   string
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        "1,2,3",
			description: "join int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        "str1,str2,str3",
			description: "join string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        "str1,1,true",
			description: "join struct",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			separator:   "|",
			want:        "1|2|3",
			description: "join '|' seperator",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			separator:   `"`,
			want:        `1"2"3`,
			description: `join '"' seperator`,
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		got := arr.Join(tt.separator)
		TestLog("Join", t, fmt.Sprintf("incoming:%v, separator:%v", tt.incoming, tt.separator), got, tt.want, tt.description)
	}
}

func TestArrayIndexOf(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		desired     interface{}
		fromIndex   int
		want        int
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			want:        1,
			description: "indexOf int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			desired:     "str1",
			want:        0,
			description: "indexOf string",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			desired:     ArrayItem{},
			want:        -1,
			description: "indexOf ArrayItem",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   1,
			want:        0,
			description: "indexOf from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   2,
			want:        -1,
			description: "indexOf from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   5,
			want:        -1,
			description: "indexOf from index out range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   -5,
			want:        -1,
			description: "indexOf from index out range",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		got := arr.IndexOf(tt.desired, tt.fromIndex)
		TestLog("IndexOf", t, fmt.Sprintf("incoming:%v, desired:%v, fromIndex:%v", tt.incoming, tt.desired, tt.fromIndex), got, tt.want, tt.description)
	}
}

func TestArrayLastIndexOf(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		desired     interface{}
		fromIndex   int
		want        int
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			want:        1,
			description: "lastIndexOf int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			desired:     "str1",
			want:        0,
			description: "lastIndexOf string",
		},
		{
			incoming:    []interface{}{"str1", 2, true},
			desired:     ArrayItem{},
			want:        -1,
			description: "lastIndexOf ArrayItem",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   1,
			want:        1,
			description: "lastIndexOf from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   2,
			want:        -1,
			description: "lastIndexOf from index in range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   5,
			want:        -1,
			description: "lastIndexOf from index out range",
		},
		{
			incoming:    []interface{}{1, 2, 3},
			desired:     2,
			fromIndex:   -5,
			want:        -1,
			description: "lastIndexOf from index out range",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		got := arr.LastIndexOf(tt.desired, tt.fromIndex)
		TestLog("LastIndexOf", t, fmt.Sprintf("incoming:%v, desired:%v, fromIndex:%v", tt.incoming, tt.desired, tt.fromIndex), got, tt.want, tt.description)
	}
}

func TestArrayReverse(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        &Array{[]ArrayItem{{3}, {2}, {1}}},
			description: "reverse int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        &Array{[]ArrayItem{{"str3"}, {"str2"}, {"str1"}}},
			description: "reverse string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        &Array{[]ArrayItem{{true}, {1}, {"str1"}}},
			description: "reverse mix",
		},
	}

	for _, tt := range tests {
		arr := MakeArray(tt.incoming...)
		got := arr.Reverse()
		TestLog("Reverse", t, tt.incoming, got, tt.want, tt.description)
	}
}

func TestArrayFilter(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        &Array{[]ArrayItem{{3}}},
			description: "filter int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        &Array{[]ArrayItem{{"str2"}, {"str3"}}},
			description: "filter string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        &Array{[]ArrayItem{{true}}},
			description: "filter mix",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.Filter(func(value ArrayItem, index int, array *Array) bool {
		return value.Data.(int) > 2
	})
	TestLog("Filter", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.Filter(func(value ArrayItem, index int, array *Array) bool {
		return value.Data != "str1"
	})
	TestLog("Filter", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.Filter(func(value ArrayItem, index int, array *Array) bool {
		return value.Data == true
	})
	TestLog("Filter", t, tt.incoming, got, tt.want, tt.description)
}

func TestArrayMap(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 2, 3},
			want:        &Array{[]ArrayItem{{3}, {4}, {5}}},
			description: "map int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        &Array{[]ArrayItem{{"sstr1"}, {"sstr2"}, {"sstr3"}}},
			description: "map string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        &Array{[]ArrayItem{{}, {}, {true}}},
			description: "map mix",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.Map(func(value ArrayItem, index int, array *Array) ArrayItem {
		value.Data = value.Data.(int) + 2
		return value
	})
	TestLog("Map", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.Map(func(value ArrayItem, index int, array *Array) ArrayItem {
		value.Data = "s" + value.Data.(string)
		return value
	})
	TestLog("Map", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.Map(func(value ArrayItem, index int, array *Array) ArrayItem {
		if value.Data == true {
			return value
		}
		return ArrayItem{}
	})
	TestLog("Map", t, tt.incoming, got, tt.want, tt.description)
}

func TestArrayReduce(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        interface{}
		description string
	}{
		{
			incoming:    []interface{}{2, 3},
			want:        64,
			description: "reduce int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        "str1str2str3",
			description: "reduce string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        1,
			description: "reduce mix",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.Reduce(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		p := reflect.ValueOf(prevValue).Int()
		c := reflect.ValueOf(currValue.Data).Int()
		return int(math.Pow(float64(p), float64(c)))
	}, 2)
	TestLog("Reduce", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.Reduce(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		return prevValue.(string) + currValue.Data.(string)
	}, "")
	TestLog("Reduce", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.Reduce(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		if reflect.ValueOf(currValue.Data).CanConvert(reflect.TypeOf(1)) {
			return prevValue.(int) + currValue.Data.(int)
		}
		return prevValue
	}, 0)
	TestLog("Reduce", t, tt.incoming, got, tt.want, tt.description)
}

func TestArrayReduceRight(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        interface{}
		description string
	}{
		{
			incoming:    []interface{}{2, 3},
			want:        64,
			description: "reduceRight int",
		},
		{
			incoming:    []interface{}{"str1", "str2", "str3"},
			want:        "str3str2str1",
			description: "reduceRight string",
		},
		{
			incoming:    []interface{}{"str1", 1, true},
			want:        1,
			description: "reduceRight mix",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.ReduceRight(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		p := reflect.ValueOf(prevValue).Int()
		c := reflect.ValueOf(currValue.Data).Int()
		return int(math.Pow(float64(p), float64(c)))
	}, 2)
	TestLog("ReduceRight", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.ReduceRight(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		return prevValue.(string) + currValue.Data.(string)
	}, "")
	TestLog("ReduceRight", t, tt.incoming, got, tt.want, tt.description)

	// test3
	tt = tests[2]
	arr = MakeArray(tt.incoming...)
	got = arr.ReduceRight(func(prevValue interface{}, currValue ArrayItem, index int, array *Array) interface{} {
		if reflect.ValueOf(currValue.Data).CanConvert(reflect.TypeOf(1)) {
			return prevValue.(int) + currValue.Data.(int)
		}
		return prevValue
	}, 0)
	TestLog("ReduceRight", t, tt.incoming, got, tt.want, tt.description)
}

func TestArraySort(t *testing.T) {
	tests := []struct {
		incoming    []interface{}
		want        *Array
		description string
	}{
		{
			incoming:    []interface{}{1, 5, 4, 8},
			want:        &Array{[]ArrayItem{{1}, {4}, {5}, {8}}},
			description: "sort int",
		},
		{
			incoming:    []interface{}{"str3", "str1", "str2"},
			want:        &Array{[]ArrayItem{{"str1"}, {"str2"}, {"str3"}}},
			description: "sort string",
		},
	}

	// test1
	tt := tests[0]
	arr := MakeArray(tt.incoming...)
	got := arr.Sort(func(a, b ArrayItem) int {
		if a.Data.(int) > b.Data.(int) {
			return 1
		} else if a.Data.(int) == b.Data.(int) {
			return 0
		}
		return -1
	})
	TestLog("Sort", t, tt.incoming, got, tt.want, tt.description)

	// test2
	tt = tests[1]
	arr = MakeArray(tt.incoming...)
	got = arr.Sort(func(a, b ArrayItem) int {
		if a.Data.(string) > b.Data.(string) {
			return 1
		} else if a.Data == b.Data {
			return 0
		}
		return -1
	})
	TestLog("Sort", t, tt.incoming, got, tt.want, tt.description)
}
