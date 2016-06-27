package formam

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"testing"
	"time"
)

type Text string

func (s *Text) UnmarshalText(text []byte) error {
	var n Text
	n = "the string has changed by UnmarshalText method"
	*s = n
	return nil
}

type UUID [16]byte

func (u *UUID) UnmarshalText(text []byte) error {
	if len(text) != 32 {
		return fmt.Errorf("text must be exactly 16 bytes long, got %d bytes", len(text))
	}
	_, err := hex.Decode(u[:], text)
	if err != nil {
		return err
	}
	return nil
}

func (u UUID) String() string {
	buf := make([]byte, 32)
	hex.Encode(buf[:], u[:])
	return string(buf)
}

const unmarshalTextString = "If you see this text, then it's a bug"

type Anonymous struct {
	AnonymousField string
	FieldOverride  string
}

type TestStruct struct {
	Anonymous
	FieldOverride string

	// traverse
	TraverseStruct struct {
		Field1 [][]struct {
			Field string
		}
		Field2 struct {
			Field int
		}
	}
	TraverseMapByBracket map[string]map[int]map[uint]map[bool]*string
	TraverseMapByPoint   map[string]map[int]map[uint]map[bool]string

	// slices/arrays
	SlicesWithIndex      []string
	SlicesWithoutIndex   []float32
	SlicesMultiDimension [][][][]uintptr
	ArrayWithIndex       [2]interface{}
	ArrayWithoutIndex    [2]bool
	ArrayMultiDimension  [2][2]bool

	// int
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	// uint
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr

	// bool
	Bool bool

	// string
	String string

	// byte
	Byte byte

	// pointer
	Pointer *string
	// pointer to struct
	PointerToStruct *struct{ Field float64 }
	// pointer to map
	PointerToMap *map[string]string

	// map
	Map map[string]string
	// mp with slice
	MapWithSlice map[string][]string
	// map with multi dimension slice
	MapWithMultiDimensionSlice map[string][][]string
	// map with array
	MapWithArray map[string][2]int
	// map with int key
	MapWithIntKey map[int]string
	// map with uint key
	MapWithUintKey map[uint]string
	// map with float key
	MapWithFloatKey map[float32]string
	// map with boolean key
	MapWithBooleanKey map[bool]string
	// map with custom key and decode key by unmarshal key
	MapWithCustomKey map[UUID]string
	// map with custom key pointer and decode key by unmarshal key
	MapWithCustomKeyPointer map[*UUID]string

	// unmarshal text
	UnmarshalTextString Text
	UnmarshalTextUUID   UUID

	// tag
	Tag string `formam:"tag"`

	// time
	Time time.Time

	// url
	URL url.URL

	// interface
	Interface interface{}
	// interface with struct as data
	InterfaceStruct interface{}
}

type InterfaceStruct struct {
	ID   int
	Name string
}

var vals = url.Values{
	// anonymous
	"AnonymousField": []string{"anonymous field"},
	"FieldOverride":  []string{"field not override"},

	// traverse
	"TraverseStruct.Field1[0][2].Field":            []string{"traverse over structs is recursive"},
	"TraverseStruct.Field2.Field":                  []string{"2"},
	"TraverseMapByBracket[by-bracket][1][2][true]": []string{"traverse over map by bracket is recursive too"},
	"TraverseMapByPoint.by-point.1.2.true":         []string{"traverse over map by point is recursive too"},

	// slices/arrays
	"SlicesWithIndex[0]":               []string{"slice index 0"},
	"SlicesWithIndex[2]":               []string{"slice index 2"},
	"SlicesWithIndex[4]":               []string{"slice index 4"},
	"SlicesWithoutIndex":               []string{"1.111", "2.222", "3.333"},
	"SlicesMultiDimension[0][1][2][3]": []string{"8"},
	"ArrayWithIndex[0]":                []string{"array index 0"},
	"ArrayWithIndex[1]":                []string{"array index 1"},
	"ArrayWithoutIndex":                []string{"true", "true"},
	"ArrayMultiDimension[0][0]":        []string{"true"},
	"ArrayMultiDimension[0][1]":        []string{"true"},
	"ArrayMultiDimension[1][0]":        []string{"true"},
	"ArrayMultiDimension[1][1]":        []string{"true"},

	// int
	"Int":   []string{"-1"},
	"Int8":  []string{"-1"},
	"Int16": []string{"-1"},
	"Int32": []string{"-1"},
	"Int64": []string{"-1"},

	// uint
	"Uint":    []string{"1"},
	"Uint8":   []string{"1"},
	"Uint16":  []string{"1"},
	"Uint32":  []string{"1"},
	"Uint64":  []string{"1"},
	"Uintptr": []string{"10"},

	// bool
	"Bool": []string{"true"},

	// string
	"String": []string{"string"},

	// byte
	"Byte": []string{"20"},

	// map
	"Map[by.bracket.with.point]":                                []string{"by bracket"},
	"Map.by_point":                                              []string{"by point"},
	"MapWithSlice[slice][0]":                                    []string{"map with slice"},
	"MapWithMultiDimensionSlice[slice][0][1]":                   []string{"map with multidimension slice"},
	"MapWithArray[array][0]":                                    []string{"0"},
	"MapWithArray[array][1]":                                    []string{"1"},
	"MapWithIntKey[-1]":                                         []string{"int key in map"},
	"MapWithUintKey[1]":                                         []string{"uint key in map"},
	"MapWithFloatKey[3.14]":                                     []string{"float key in map"},
	"MapWithBooleanKey[true]":                                   []string{"bool key in map"},
	"MapWithCustomKey[11e5bf2d3e403a8c86740023dffe5350]":        []string{"UUID key in map"},
	"MapWithCustomKeyPointer[11e5bf2d3e403a8c86740023dffe5350]": []string{"UUID key pointer in map"},

	// unmarshal text
	"UnmarshalTextString": []string{"If you see this text, then it's a bug"},
	"UnmarshalTextUUID":   []string{"11e5bf2d3e403a8c86740023dffe5350"},

	// tag
	"tag": []string{"string placed by tag"},

	// time
	"Time": []string{"2016-06-12"},

	// url
	"URL": []string{"http://www.google.com"},

	// interface
	"Interface":            []string{"Germany"},
	"InterfaceStruct.ID":   []string{"1"},
	"InterfaceStruct.Name": []string{"Germany"},
}

func TestDecodeInStruct(t *testing.T) {
	var m TestStruct
	m.InterfaceStruct = &InterfaceStruct{}
	err := Decode(vals, &m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// anonymous struct
	if m.Anonymous.AnonymousField == "" {
		t.Error("AnonymousField is empty")
	}
	if m.Anonymous.FieldOverride != "" {
		t.Error("FieldOverride is full")
	}
	if m.FieldOverride == "" {
		t.Error("FieldOverride is empty")
	}

	// traverse
	// traverse > struct
	if len(m.TraverseStruct.Field1) == 0 {
		t.Error("TraverseStruct.Field1 is empty")
	} else {
		if len(m.TraverseStruct.Field1[0]) != 3 {
			t.Errorf("TraverseStruct.Field1[0] must has 3 as length but has %v", len(m.TraverseStruct.Field1[0]))
		} else {
			if m.TraverseStruct.Field1[0][2].Field == "" {
				t.Error("TraverseStruct.Field1[0][2].Field is empty")
			}
		}
	}
	// traverse > maps
	n, ok := m.TraverseMapByBracket["by-bracket"]
	if ok {
		m, ok := n[1]
		if ok {
			j, ok := m[2]
			if ok {
				g, ok := j[true]
				if ok {
					if *g == "" {
						t.Error("the value of TraverseMapByBracket in the last field is empty")
					}
				} else {
					t.Error("the key \"true\" in TraverseMapByBracket not exists")
				}
			} else {
				t.Error("the key \"2\" in TraverseMapByBracket not exists")
			}
		} else {
			t.Error("the key \"1\" in TraverseMapByBracket not exists")
		}
	} else {
		t.Error("the key \"bracket\" in TraverseMapByBracket not exists")
	}
	u, ok := m.TraverseMapByPoint["by-point"]
	if ok {
		m, ok := u[1]
		if ok {
			j, ok := m[2]
			if ok {
				g, ok := j[true]
				if ok {
					if g == "" {
						t.Error("the value of TraverseMapByPoint in the last field is empty")
					}
				} else {
					t.Error("the key \"true\" in TraverseMapByPoint not exists")
				}
			} else {
				t.Error("the key \"2\" in TraverseMapByPoint not exists")
			}
		} else {
			t.Error("the key \"1\" in TraverseMapByPoint not exists")
		}
	} else {
		t.Error("the key \"by-point\" in TraverseMapByPoint not exists")
	}

	// slices
	if len(m.SlicesWithIndex) != 5 {
		t.Error("the length SlicesWithIndex is not 5")
	}
	if len(m.SlicesWithoutIndex) != 3 {
		t.Error("the length SlicesWithoutIndex is not 3")
	}
	if len(m.SlicesMultiDimension) != 1 {
		t.Error("the length SlicesMultiDimension is not 1")
	}
	if len(m.SlicesMultiDimension[0]) != 2 {
		t.Error("the length SlicesMultiDimension[0] is not 2")
	}
	if len(m.SlicesMultiDimension[0][1]) != 3 {
		t.Error("the length SlicesMultiDimension[0] is not 3")
	}
	// array
	if len(m.ArrayWithIndex) != 2 {
		t.Error("the length ArrayWithIndex is not 2")
	}
	if len(m.ArrayWithoutIndex) != 2 {
		t.Error("the length ArrayWithoutIndex is not 2")
	}
	if len(m.ArrayMultiDimension) != 2 {
		t.Error("the length ArrayMultiDimension is not 2")
	}
	if len(m.ArrayMultiDimension[0]) != 2 {
		t.Error("the length ArrayMultiDimension[0] is not 2")
	}
	if len(m.ArrayMultiDimension[1]) != 2 {
		t.Error("the length ArrayMultiDimension[1] is not 2")
	}

	// int
	if m.Int != -1 {
		t.Error("the length Int is not -1")
	}
	if m.Int8 != -1 {
		t.Error("the length Int8 is not -1")
	}
	if m.Int16 != -1 {
		t.Error("the length Int16 is not -1")
	}
	if m.Int32 != -1 {
		t.Error("the length Int32 is not -1")
	}
	if m.Int64 != -1 {
		t.Error("the length Int64 is not -1")
	}

	// uint
	if m.Uint != 1 {
		t.Error("the length Uint is not 1")
	}
	if m.Uint8 != 1 {
		t.Error("the length Uint8 is not 1")
	}
	if m.Uint16 != 1 {
		t.Error("the length Uint16 is not 1")
	}
	if m.Uint32 != 1 {
		t.Error("the length Uint32 is not 1")
	}
	if m.Uint64 != 1 {
		t.Error("the length Uint64 is not 1")
	}
	if m.Uintptr != 10 {
		t.Error("the length Uintptr is not 10")
	}

	// bool
	if !m.Bool {
		t.Error("Bool is false")
	}

	// string
	if m.String == "" {
		t.Error("String is empty")
	}

	// byte
	if string(m.Byte) == "" {
		t.Error("Byte is empty")
	}

	// map
	f, ok := m.Map["by.bracket.with.point"]
	if ok {
		if f == "" {
			t.Error("The value in key \"by.bracket.with.point\" of Map is empty")
		}
	} else {
		t.Error("The key \"by.bracket.with.point\" in Map not exists")
	}
	f, ok = m.Map["by_point"]
	if ok {
		if f == "" {
			t.Error("The value in key \"by_point\" of Map is empty")
		}
	} else {
		t.Error("The key \"by_point\" in Map not exists")
	}
	s, ok := m.MapWithSlice["slice"]
	if ok {
		if len(s) == 0 {
			t.Error("The length of key \"slice\" of MapWithSlice is 0")
		} else {
			if s[0] == "" {
				t.Error("The value of key \"slice\" in MapWithSlice is empty")
			}
		}
	} else {
		t.Error("The key \"slice\" in MapWithSlice not exists")
	}
	a, ok := m.MapWithMultiDimensionSlice["slice"]
	if ok {
		if len(a) == 0 {
			t.Error("The length of key \"slice\" of MapWithSlice is 0")
		} else {
			if len(a) == 0 {
				t.Error("The length of MapWithMultiDimensionSlice[slice] is 0")
			} else {
				if len(a[0]) != 2 {
					t.Error("The length of MapWithMultiDimensionSlice[slice][0] is not 2")
				} else {
					if a[0][1] == "" {
						t.Error("The value in MapWithSlice[slice][0][1] is empty")
					}
				}
			}
		}
	} else {
		t.Error("The key \"slice\" in MapWithSlice not exists")
	}
	w, ok := m.MapWithArray["array"]
	if ok {
		if len(w) != 2 {
			t.Error("The length of MapWithArray[array] is not 2")
		}
	} else {
		t.Error("The key \"array\" in MapWithArray not exists")
	}
	q, ok := m.MapWithIntKey[-1]
	if ok {
		if q == "" {
			t.Error("The value of MapWithIntKey[-1] is empty")
		}
	} else {
		t.Error("The key \"-1\" in MapWithIntKey not exists")
	}
	y, ok := m.MapWithUintKey[1]
	if ok {
		if y == "" {
			t.Error("The value of MapWithUintKey[1] is empty")
		}
	} else {
		t.Error("The key \"1\" in MapWithUintKey not exists")
	}
	o, ok := m.MapWithFloatKey[3.14]
	if ok {
		if o == "" {
			t.Error("The value of MapWithFloatKey[3.14] is empty")
		}
	} else {
		t.Error("The key \"3.14\" in MapWithFloatKey not exists")
	}
	b, ok := m.MapWithBooleanKey[true]
	if ok {
		if b == "" {
			t.Error("The value of MapWithFloatKey[true] is empty")
		}
	} else {
		t.Error("The key \"true\" in MapWithFloatKey not exists")
	}
	uuid := UUID{17, 229, 191, 45, 62, 64, 58, 140, 134, 116, 0, 35, 223, 254, 83, 80}
	uu, ok := m.MapWithCustomKey[uuid]
	if ok {
		if uu == "" {
			t.Error("The value of MapWithFloatKey[11e5bf2d3e403a8c86740023dffe5350] is empty")
		}
	} else {
		t.Error("The key \"11e5bf2d3e403a8c86740023dffe5350\" in MapWithCustomKey not exists")
	}
	for k, v := range m.MapWithCustomKeyPointer {
		if k.String() != uuid.String() {
			t.Error("The key in MapWithCustomKeyPointer is not 11e5bf2d3e403a8c86740023dffe5350")
		} else if v == "" {
			t.Error("The value of MapWithCustomKeyPointer[11e5bf2d3e403a8c86740023dffe5350] is empty")
		}
	}

	// unmarshalText
	if m.UnmarshalTextString == unmarshalTextString {
		t.Error("The value of UnmarshalTextString is not correct. It should not to contain the text of the const unmarshalTextString")
	}
	if m.UnmarshalTextUUID.String() != uuid.String() {
		t.Errorf("The value of UnmarshalTextUUID is not 11e5bf2d3e403a8c86740023dffe5350 but %s", m.UnmarshalTextUUID.String())
	}

	// tag
	if m.Tag == "" {
		t.Error("The value of UnmarshalTextString is empty")
	}

	// time
	if m.Time.IsZero() {
		t.Error("The value of Time is zero")
	}

	// interface
	if v, ok := m.Interface.(string); !ok {
		t.Error("The Interface is not string")
	} else if v == "" {
		t.Error("The value of Interface is empty")
	}
	if v, ok := m.InterfaceStruct.(*InterfaceStruct); !ok {
		t.Error("The InterfaceStruct is not InterfaceStruct struct")
	} else {
		if v.ID == 0 {
			t.Error("The value of InterfaceStruct.ID is 0")
		}
		if v.Name == "" {
			t.Error("The value of InterfaceStruct.Name is empty")
		}
	}

	fmt.Println("RESULT: ", m)
}

type TestSlice []string

var sliceValues = url.Values{
	"[0]": []string{"spanish"},
	"[1]": []string{"english"},
}

func TestDecodeInSlice(t *testing.T) {
	var t2 TestSlice
	err := Decode(sliceValues, &t2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("RESULT: ", t2)
}
