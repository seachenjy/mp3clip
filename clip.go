package mp3clip

import (
	"os"

	"github.com/tcolgate/mp3"
)

//Cliper cliper
type Cliper struct {
	File *os.File
}

//DvInfo the mp3 headers message
type DvInfo struct {
}

//OpenMp3 open one mp3 file
func OpenMp3(path string) {
	f,err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	d := NewDecoder(r)
	var f Frame
	for {
	
		if err := d.Decode(&f, &skipped); err != nil {
			fmt.Println(err)
			return
		}
	
		fmt.Println(&f)
	}
}
