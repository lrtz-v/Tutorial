package find

import "testing"

func TestFindFiles(t *testing.T) {
	res := Files("./", "README.md")
	if len(res) == 0 {
		t.Fatal("[*] FindFiles Error!")
	} 
}
