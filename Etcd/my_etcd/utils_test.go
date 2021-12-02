package main

import "testing"

func TestGenerater(t *testing.T) {
	filepath := generatePath(".", "test.log")
	if filepath != "./test.log" {
		t.Fatalf("TestGenerater failed, Expected %s, Got %s\n", "./test.log", filepath)
	}
}

func TestSave(t *testing.T) {
	err := Save(".", "test.log", []byte("test data"))
	if err != nil {
		t.Fatalf("TestSave failed, Got %s\n", err.Error())
	}
}

func TestMatchString(t *testing.T) {

	match, err := MatchString(`foo.*`, "seafood")
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal()
	}

	match, err = MatchString(`foo.*`, "1234567890")
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal()
	}
}


func TestSet(t *testing.T) {
	arr := []string{"1", "1", "2", "2", "3", "3"}
	arr = DistinctStringList(arr)
	if len(arr) != 3 {
		t.Fatalf("distinct failed. len: %d", len(arr))
	}
}
