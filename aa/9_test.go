package aa

import (
	"a/utils"
	"fmt"
	"testing"
)

func TestName111(t *testing.T) {
	fmt.Println(utils.GetHourDiffer("2022-09-01 04:00:00", "2022-09-02 08:00:00"))
}
