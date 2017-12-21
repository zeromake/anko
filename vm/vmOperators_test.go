package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBasicOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	testInput1 := map[string]interface{}{"b": func() {}}
	tests := []testStruct{
		{script: "]", parseError: fmt.Errorf("syntax error"), runOutput: nil},
		{script: "1 = 2", runError: fmt.Errorf("Invalid operation"), runOutput: nil},

		{script: "2 + 1", runOutput: int64(3)},
		{script: "2 - 1", runOutput: int64(1)},
		{script: "2 * 1", runOutput: int64(2)},
		{script: "2 / 1", runOutput: float64(2)},
		{script: "2.1 + 1.1", runOutput: float64(3.2)},
		{script: "2.1 - 1.1", runOutput: float64(1)},
		{script: "2.1 * 2.0", runOutput: float64(4.2)},
		{script: "6.5 / 2.0", runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3)},
		{script: "a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1)},
		{script: "a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2)},
		{script: "a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2)},
		{script: "a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2)},
		{script: "a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1)},
		{script: "a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2)},
		{script: "a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": "a", "b": "b"}, runOutput: "ab"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": int64(1)}, runOutput: "a1"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": float64(1.1)}, runOutput: "a1.1"},
		{script: "a + z", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z + b", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = nil; a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = b", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = b; b = nil", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = true", runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "a = false", input: map[string]interface{}{"a": true}, runOutput: false, ouput: map[string]interface{}{"a": false}},

		{script: "a = b", input: map[string]interface{}{"b": reflect.Value{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{script: "a = b", input: map[string]interface{}{"b": struct{}{}}, runOutput: struct{}{}, ouput: map[string]interface{}{"a": struct{}{}, "b": struct{}{}}},
		{script: "a = b", input: testInput1, runOutput: testInput1["b"], ouput: map[string]interface{}{"a": testInput1["b"], "b": testInput1["b"]}},

		{script: "a = b", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = b", input: map[string]interface{}{"b": true}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": float64(2.1)}, runOutput: float64(2.1), ouput: map[string]interface{}{"a": float64(2.1), "b": float64(2.1)}},
		{script: "a = b", input: map[string]interface{}{"b": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a", "b": "a"}},
		{script: "a = b", input: map[string]interface{}{"b": 'a'}, runOutput: 'a', ouput: map[string]interface{}{"a": 'a', "b": 'a'}},

		{script: "a = b", input: map[string]interface{}{"b": testVarValue}, runOutput: testVarValue, ouput: map[string]interface{}{"a": testVarValue, "b": testVarValue}},
		{script: "a = b", input: map[string]interface{}{"b": testVarBoolP}, runOutput: testVarBoolP, ouput: map[string]interface{}{"a": testVarBoolP, "b": testVarBoolP}},
		{script: "a = b", input: map[string]interface{}{"b": testVarInt32P}, runOutput: testVarInt32P, ouput: map[string]interface{}{"a": testVarInt32P, "b": testVarInt32P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarInt64P}, runOutput: testVarInt64P, ouput: map[string]interface{}{"a": testVarInt64P, "b": testVarInt64P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarFloat32P}, runOutput: testVarFloat32P, ouput: map[string]interface{}{"a": testVarFloat32P, "b": testVarFloat32P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarFloat64P}, runOutput: testVarFloat64P, ouput: map[string]interface{}{"a": testVarFloat64P, "b": testVarFloat64P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarStringP}, runOutput: testVarStringP, ouput: map[string]interface{}{"a": testVarStringP, "b": testVarStringP}},
		{script: "a = b", input: map[string]interface{}{"b": testVarFuncP}, runOutput: testVarFuncP, ouput: map[string]interface{}{"a": testVarFuncP, "b": testVarFuncP}},

		{script: "a = b", input: map[string]interface{}{"b": testVarValueBool}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueInt32}, runOutput: int32(1), ouput: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueInt64}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueFloat32}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueFloat64}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueString}, runOutput: "a", ouput: map[string]interface{}{"a": "a", "b": "a"}},

		{script: "a, b = 1, 2", input: map[string]interface{}{"a": int64(3)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "a, b, c = 1, 2, 3", input: map[string]interface{}{"a": int64(3)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{script: "a, b = [1, 2], [3, 4]", runOutput: []interface{}{int64(3), int64(4)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": []interface{}{int64(3), int64(4)}}},

		{script: "y = z", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z.y.x = 1", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "c = a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3), ouput: map[string]interface{}{"c": int64(3)}},
		{script: "c = a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"c": int64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"c": int64(2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2), ouput: map[string]interface{}{"c": float64(2)}},
		{script: "c = a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2), ouput: map[string]interface{}{"c": float64(3.2)}},
		{script: "c = a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1), ouput: map[string]interface{}{"c": float64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2), ouput: map[string]interface{}{"c": float64(4.2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25), ouput: map[string]interface{}{"c": float64(3.25)}},

		{script: "a++", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a--", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a++", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},

		{script: "1++", runError: fmt.Errorf("Invalid operation"), runOutput: nil},
		{script: "1--", runError: fmt.Errorf("Invalid operation"), runOutput: nil},
		{script: "z++", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z--", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "a += 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a -= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), ouput: map[string]interface{}{"a": int64(4)}},
		{script: "a /= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1), ouput: map[string]interface{}{"a": float64(1)}},
		{script: "a += 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a -= 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": 2.1}, runOutput: float64(4.2), ouput: map[string]interface{}{"a": float64(4.2)}},
		{script: "a /= 2", input: map[string]interface{}{"a": 6.5}, runOutput: float64(3.25), ouput: map[string]interface{}{"a": float64(3.25)}},

		{script: "a ** 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a ** 3", input: map[string]interface{}{"a": float64(2)}, runOutput: float64(8), ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a &= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a &= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a |= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a |= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a << 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": int64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(8)}},
		{script: "a << 2", input: map[string]interface{}{"a": float64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": float64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(8)}},

		{script: "a % 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": float64(2.1)}},
		{script: "a % 3", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(2.1)}},

		{script: "-a", input: map[string]interface{}{"a": nil}, runOutput: float64(-0), ouput: map[string]interface{}{"a": nil}},
		{script: "-a", input: map[string]interface{}{"a": "a"}, runOutput: float64(-0), ouput: map[string]interface{}{"a": "a"}},
		{script: "-a", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(-2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "-a", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(-2.1), ouput: map[string]interface{}{"a": float64(2.1)}},
		{script: "^a", input: map[string]interface{}{"a": nil}, runOutput: int64(-1), ouput: map[string]interface{}{"a": nil}},
		{script: "^a", input: map[string]interface{}{"a": "a"}, runOutput: int64(-1), ouput: map[string]interface{}{"a": "a"}},
		{script: "^a", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(-3), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "^a", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(-3), ouput: map[string]interface{}{"a": float64(2.1)}},

		{script: "a * 4", input: map[string]interface{}{"a": "a"}, runOutput: "aaaa", ouput: map[string]interface{}{"a": "a"}},
		{script: "a * 4.0", input: map[string]interface{}{"a": "a"}, runOutput: float64(0), ouput: map[string]interface{}{"a": "a"}},
		{script: "!true", runOutput: false},
		{script: "!1", runOutput: false},
	}
	runTests(t, tests)
}

func TestComparisonOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a == 1", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a == 1", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == nil", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},
		{script: "a == 2.0", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},

		{script: "1 == 1.0", runOutput: true},
		{script: "1 != 1.0", runOutput: false},
		{script: "\"a\" != \"a\"", runOutput: false},

		{script: "a > 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a > 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a >= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 3", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 3", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a >= 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 3", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 3", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "1 == 1 && 1  == 1", runOutput: true},
		{script: "1 == 1 && 1  == 2", runOutput: false},
		{script: "1 == 2 && 1  == 1", runOutput: false},
		{script: "1 == 2 && 1  == 2", runOutput: false},
		{script: "1 == 1 || 1  == 1", runOutput: true},
		{script: "1 == 1 || 1  == 2", runOutput: true},
		{script: "1 == 2 || 1  == 1", runOutput: true},
		{script: "1 == 2 || 1  == 2", runOutput: false},

		{script: `true == "true"`, runOutput: true},
		{script: `true == "TRUE"`, runOutput: true},
		{script: `true == "True"`, runOutput: true},
		{script: `true == "false"`, runOutput: false},
		{script: `true == "foo"`, runOutput: false},
		{script: `false == "false"`, runOutput: true},
		{script: `false == "FALSE"`, runOutput: true},
		{script: `false == "False"`, runOutput: true},
		{script: `false == "true"`, runOutput: false},
		{script: `false == "foo"`, runOutput: false},

		{script: `0 == "0"`, runOutput: true},
		{script: `"1.0" == 1`, runOutput: true},
		{script: `1 == "1"`, runOutput: true},
		{script: `0.0 == "0"`, runOutput: true},
		{script: `0.0 == "0.0"`, runOutput: true},
		{script: `1.0 == "1.0"`, runOutput: true},
		{script: `1.2 == "1.2"`, runOutput: true},
		{script: `"7" == 7.2`, runOutput: true},
		{script: `1.2 == "1"`, runOutput: false},
		{script: `"1.1" == 1`, runOutput: false},
		{script: `0 == "1"`, runOutput: false},

		{script: "a == b", input: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}, runOutput: true, ouput: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{script: "a == b", input: map[string]interface{}{"a": reflect.Value{}, "b": true}, runOutput: false, ouput: map[string]interface{}{"a": reflect.Value{}, "b": true}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": reflect.Value{}}, runOutput: false, ouput: map[string]interface{}{"a": true, "b": reflect.Value{}}},

		{script: "a == b", input: map[string]interface{}{"a": nil, "b": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a == b", input: map[string]interface{}{"a": nil, "b": true}, runOutput: false, ouput: map[string]interface{}{"a": nil, "b": true}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": nil}, runOutput: false, ouput: map[string]interface{}{"a": true, "b": nil}},

		{script: "a == b", input: map[string]interface{}{"a": false, "b": false}, runOutput: true, ouput: map[string]interface{}{"a": false, "b": false}},
		{script: "a == b", input: map[string]interface{}{"a": false, "b": true}, runOutput: false, ouput: map[string]interface{}{"a": false, "b": true}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": false}, runOutput: false, ouput: map[string]interface{}{"a": true, "b": false}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": true}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},

		{script: "a == b", input: map[string]interface{}{"a": int32(1), "b": int32(1)}, runOutput: true, ouput: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(1), "b": int32(2)}, runOutput: false, ouput: map[string]interface{}{"a": int32(1), "b": int32(2)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(2), "b": int32(1)}, runOutput: false, ouput: map[string]interface{}{"a": int32(2), "b": int32(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(2), "b": int32(2)}, runOutput: true, ouput: map[string]interface{}{"a": int32(2), "b": int32(2)}},

		{script: "a == b", input: map[string]interface{}{"a": int64(1), "b": int64(1)}, runOutput: true, ouput: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(1), "b": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2), "b": int64(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(2), "b": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},

		{script: "a == b", input: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}, runOutput: true, ouput: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}, runOutput: false, ouput: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}, runOutput: false, ouput: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}, runOutput: true, ouput: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}},

		{script: "a == b", input: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}, runOutput: true, ouput: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}},

		{script: "a == b", input: map[string]interface{}{"a": 'a', "b": 'a'}, runOutput: true, ouput: map[string]interface{}{"a": 'a', "b": 'a'}},
		{script: "a == b", input: map[string]interface{}{"a": 'a', "b": 'b'}, runOutput: false, ouput: map[string]interface{}{"a": 'a', "b": 'b'}},
		{script: "a == b", input: map[string]interface{}{"a": 'b', "b": 'a'}, runOutput: false, ouput: map[string]interface{}{"a": 'b', "b": 'a'}},
		{script: "a == b", input: map[string]interface{}{"a": 'b', "b": 'b'}, runOutput: true, ouput: map[string]interface{}{"a": 'b', "b": 'b'}},

		{script: "a == b", input: map[string]interface{}{"a": "a", "b": "a"}, runOutput: true, ouput: map[string]interface{}{"a": "a", "b": "a"}},
		{script: "a == b", input: map[string]interface{}{"a": "a", "b": "b"}, runOutput: false, ouput: map[string]interface{}{"a": "a", "b": "b"}},
		{script: "a == b", input: map[string]interface{}{"a": "b", "b": "a"}, runOutput: false, ouput: map[string]interface{}{"a": "b", "b": "a"}},
		{script: "a == b", input: map[string]interface{}{"a": "b", "b": "b"}, runOutput: true, ouput: map[string]interface{}{"a": "b", "b": "b"}},

		{script: "a = \"test\"; a == \"test\"", runOutput: true},
		{script: "a = \"test\"; a[0:1] == \"t\"", runOutput: true},
		{script: "a = \"test\"; a[0:2] == \"te\"", runOutput: true},
		{script: "a = \"test\"; a[1:3] == \"es\"", runOutput: true},
		{script: "a = \"test\"; a[0:4] == \"test\"", runOutput: true},

		// TOFIX:
		//		{script: "a = \"a b\"; a[1] == ' '", runOutput: true},
		//		{script: "a = \"test\"; a[0] == 't'", runOutput: true},
		//		{script: "a = \"test\"; a[1] == 'e'", runOutput: true},
		//		{script: "a = \"test\"; a[3] == 't'", runOutput: true},

		// TOFIX:
		//		{script: "a = \"a b\"; a[1] != ' '", runOutput: false},
		//		{script: "a = \"test\"; a[0] != 't'", runOutput: false},
		//		{script: "a = \"test\"; a[1] != 'e'", runOutput: false},
		//		{script: "a = \"test\"; a[3] != 't'", runOutput: false},
	}
	runTests(t, tests)
}