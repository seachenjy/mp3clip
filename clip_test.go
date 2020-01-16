package mp3clip

import "testing"

func TestClip(t *testing.T) {
	tag, err := OpenMp3("./test.mp3")
	if err != nil {
		t.Error(err)
	} else {
		t.Errorf("%+v", tag)
	}

	defer tag.Close()

}
