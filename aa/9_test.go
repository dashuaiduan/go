package aa

import (
	"a/utils"
	"fmt"
	"sync"
	"testing"
)

func TestName111(t *testing.T) {
	fmt.Println(utils.GetHourDiffer("2022/09/01 04:00:00", "2022/09/02 08:00:00"))

	m := map[string]string{"dsgf": "11", "dd": "555"}
	fmt.Println(m["356"])
}

func Test333(t *testing.T) {
	var m sync.Map
	m.Store(1, 111)
	m.Store(1, 113)
	m.Delete(1)
	fmt.Println(m.Load(1))
	fmt.Println(m)
}
