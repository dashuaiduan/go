package aa

import (
	"encoding/json"
	"fmt"
	"testing"
)

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := `{"Name":"zs","Age":16,"Sex":"nan"}`
		user := new(user)
		json.Unmarshal([]byte(str), user)
		fmt.Println(user)
		v, err := json.Marshal(user)
		fmt.Println(string(v), err)
	}

}

func BenchmarkEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := `{"Name":"zs","Age":16,"Sex":"nan"}`
		obj := new(user)
		obj.UnmarshalJSON([]byte(str))
		fmt.Println(obj)
		v, err := obj.MarshalJSON()
		fmt.Println(string(v), err)
	}

}
