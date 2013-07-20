package nbt

import (
	"compress/gzip"
	"io"
	"os"
	"testing"
)

func bigTestReader() io.Reader {
	f, err := os.Open("../test_files/bigtest.nbt")
	if err != nil {
		panic(err.Error())
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func testMatchesBigTest(t *testing.T, data map[string]interface{}) {
	// Check data matches official test
	// @see http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt
	root := data
	level := root["Level"].(map[string]interface{})
	shortTest := level["shortTest"].(int16)
	if shortTest != 32767 {
		t.Fatal("shortTest not correct, got", shortTest)
	}
	longTest := level["longTest"].(int64)
	if longTest != 9223372036854775807 {
		t.Fatal("longTest not correct, got", longTest)
	}
	floatTest := level["floatTest"].(float32)
	if floatTest != 0.49823147 {
		t.Fatal("floatTest not correct, got", floatTest)
	}
	stringTest := level["stringTest"].(string)
	if stringTest != "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!" {
		t.Fatal("stringTest not correct, got", stringTest)
	}
	intTest := level["intTest"].(int32)
	if intTest != 2147483647 {
		t.Fatal("intTest not correct, got", intTest)
	}
	nestedCompoundTest := level["nested compound test"].(map[string]interface{})
	ham := nestedCompoundTest["ham"].(map[string]interface{})
	hamName := ham["name"].(string)
	if hamName != "Hampus" {
		t.Fatal("nested compound test -> ham -> name not correct, got", hamName)
	}
	hamValue := ham["value"].(float32)
	if hamValue != 0.75 {
		t.Fatal("nested compound test -> ham -> value not correct, got",
			hamValue)
	}
	egg := nestedCompoundTest["egg"].(map[string]interface{})
	eggName := egg["name"].(string)
	if eggName != "Eggbert" {
		t.Fatal("nested compound test -> egg -> name not correct, got", eggName)
	}
	eggValue := egg["value"].(float32)
	if eggValue != 0.5 {
		t.Fatal("nested compound test -> egg -> value not correct, got",
			eggValue)
	}
	listTestLong := level["listTest (long)"].(*List)
	for i := 0; i < 5; i++ {
		listTestLongItem := listTestLong.Items[i].(int64)
		if listTestLongItem != int64(11+i) {
			t.Fatal("listTest (long) value", i, "not correct, got",
				listTestLongItem)
		}
	}
	byteTest := level["byteTest"].(byte)
	if byteTest != 127 {
		t.Fatal("byteTest not correct, got", byteTest)
	}
	listTestCompound := level["listTest (compound)"].(*List)
	listTestCompound0 := listTestCompound.Items[0].(map[string]interface{})
	listTestCompound0Name := listTestCompound0["name"].(string)
	if listTestCompound0Name != "Compound tag #0" {
		t.Fatal("listTest (compound)[0] -> name not correct, got",
			listTestCompound0Name)
	}
	listTestCompound0CreatedOn := listTestCompound0["created-on"].(int64)
	if listTestCompound0CreatedOn != 1264099775885 {
		t.Fatal("listTest (compound)[0] -> created-on not correct, got",
			listTestCompound0CreatedOn)
	}
	listTestCompound1 := listTestCompound.Items[1].(map[string]interface{})
	listTestCompound1Name := listTestCompound1["name"].(string)
	if listTestCompound1Name != "Compound tag #1" {
		t.Fatal("listTest (compound)[1] -> name not correct, got",
			listTestCompound1Name)
	}
	listTestCompound1CreatedOn := listTestCompound1["created-on"].(int64)
	if listTestCompound1CreatedOn != 1264099775885 {
		t.Fatal("listTest (compound)[1] -> created-on not correct, got",
			listTestCompound1CreatedOn)
	}
	byteArrayTest := level["byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"].([]byte)
	for i := 0; i < 1000; i++ {
		calculated := byte((i*i*255 + i*7) % 100)
		if byteArrayTest[i] != calculated {
			t.Fatal("byteArrayTest", i, "not correct, got", calculated)
		}
	}
	doubleTest := level["doubleTest"].(float64)
	if doubleTest != 0.4931287132182315 {
		t.Fatal("doubleTest not correct, got", doubleTest)
	}
}
