package path

import (
	"fmt"
	"github.com/iesreza/gutil/path"
	"testing"
)

func TestPath(t *testing.T) {
	//create directory pointer
	dir := path.Dir("./test1/test2")

	//check if directory exists
	if !dir.Exist() {
		//create directory if not exist
		dir.Create()
	}

	//create file pointer inside directory
	fp := dir.File("test.txt")

	//create file with given content
	fp.Create("test content here")

	//append to file
	fp.Append("\nNew line")

	//get file content
	content, _ := fp.Content()
	fmt.Println(content)

	//return absolute path of file
	abs, _ := fp.Absolute()
	fmt.Println(abs)

	//create pointer to file
	fp = path.File("./test2.txt")

	//create file with given content
	fp.Create("")

	//write text to file
	fp.Write("Hi")

	//remove file
	fp.Remove()

	//find files inside directory
	files, _ := dir.Find("*,txt")
	fmt.Println(files)

	//go to parent dir
	parent := dir.Parent()
	parent.File("parent.file").Create("test data")

	//get current working directory
	current, _ := path.Current()
	files, _ = current.Find("*")
	fmt.Println(files)

	//remove directory files
	dir.Clean()

	//remove directory
	dir.Remove()
}
