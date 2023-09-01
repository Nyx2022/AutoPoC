package AutoPoC

import (
	"fmt"
	"testing"
)

func TestSearchByQuery(t *testing.T) {
	Ips := SearchByQuery("body=\"/general/login/index.php\"")
	fmt.Println(Ips)
}
