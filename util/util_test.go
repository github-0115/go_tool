package util

import (
	"testing"
)

func TestUtil(t *testing.T) {
	// test int to string
	i := 2
	if "2" == IS(i) {
		t.Log("Test IS:int to string")
	}

	// test string to int
	v, err := SI("e2")
	if err == nil && v == i {
		t.Log("Test SI:string to int")
	} else {
		t.Log("Test SI:" + err.Error())
	}

	// caller dir
	t.Log("Test CurDir:"+CurDir())

	// create dir
	err = MakeDir("../data")
	if err != nil {
		t.Log("Test MakeDir:" + err.Error())
	} else {
		t.Log("Test MakeDir:dir already exist")
	}

	// create dir by filename
	filename := "../data/testutil.txt"
	err = MakeDirByFile(filename)
	if err != nil {
		t.Log("Test MakeDirByFile:" + err.Error())
	} else {
		t.Log("Test MakeDirByFile: dir already exist")
	}

	// save bytes into file
	err = SaveToFile(filename, []byte("testutil"))
	if err != nil {
		t.Log("Test SaveToFile" + err.Error())
	}

	// read bytes from file
	filebytes, err := ReadfromFile(filename)
	if err != nil {
		t.Error("Test ReadfromFile:" + err.Error())
	} else {
		t.Log("Test ReadfromFile:" + string(filebytes))
	}

	t.Log(TodayString(3))

	t.Logf("%v",FileExist("../r.txt"))

	filenames,err:=ListDir(`G:\smartdogo\src\github.com\hunterhug\go_tool`,"go")
	t.Logf("%v:%v",filenames,err)
}
