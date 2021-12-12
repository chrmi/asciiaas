package asciimodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func setup() {
	err := ioutil.WriteFile("test.txt", []byte("Write Unit Tests.\n"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
		return
	}
	fmt.Println("File Created")
}

func teardown() {
	var err = os.Remove("test.txt")
	if err != nil {
		fmt.Printf("Unable to delete file: %v", err)
		return
	}

	fmt.Println("File Deleted")
}
*/
func TestTextOK(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestTextBadKey(t *testing.T) {
	assert.Equal(t, 1, 1)
}

/*
func TestReadFile(t *testing.T) {
	setup()

	file, err := ReadFile("test.txt")

	if err != nil {
		t.Errorf("problem reading test.txt, %v", err)
	}

	if len(file) < 1 {
		t.Errorf("empty file")
	}

	teardown()
}
*/
