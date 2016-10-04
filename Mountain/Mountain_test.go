package Mountain

import (
	"testing"
	"fmt"
)

func Test_Mountains(t *testing.T) {
	for _, mount := range mountain_List{
		fmt.Printf("%s %f \n",
			mount.name, mountain_name.Dist(mount))
	}
}

