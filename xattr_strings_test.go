package xattr

import "fmt"
import "io/ioutil"
import "os"
import "testing"

func TestStrings(t *testing.T) {
    var test_xattrName = "user.xattr.test"
    var test_xattrValue = "1234567890"

    tmpfile, err := ioutil.TempFile("", "xattr_Test")
    if err != nil {
        t.Fatal(err)
    }

    defer os.Remove(tmpfile.Name())

    if err := SetString(tmpfile.Name(), test_xattrName, test_xattrValue); err != nil {
        t.Fatalf("SetString failed: %v\n", err)
    }

    if value, err := GetString(tmpfile.Name(), test_xattrName); value != test_xattrValue {
        t.Fatalf("GetString failed: %v\n\tExpected: '%v'\n\tFound: '%v'\n", err, test_xattrValue, value)
    }

    if attributes, err := ListStrings(tmpfile.Name()); attributes[0] != test_xattrName {
        t.Fatalf("ListStrings failed!\n\tStrings: %v\n\tError: %v\n", attributes, err)
    }

    if err := Remove(tmpfile.Name(), test_xattrName); err != nil {
        t.Fatalf("Remove failed: %v\n", err)
    }
}


func TestListStrings(t *testing.T) {
    var test_xattrName = "user.xattr.test.%v"

    tmpfile, err := ioutil.TempFile("", "xattr_Test")
    if err != nil {
        t.Fatal(err)
    }

    defer os.Remove(tmpfile.Name())

    for i := 0; i < 10; i++ {
        xattrName := fmt.Sprintf(test_xattrName, i)
        xattrValue := fmt.Sprintf("%v", i)
        if err := SetString(tmpfile.Name(), xattrName, xattrValue); err != nil {
            t.Fatalf("SetString failed: %v\n", err)
        }

        if attributes, err := ListStrings(tmpfile.Name()); attributes[len(attributes)-1] != xattrName {
            t.Fatalf("ListStrings failed!\n\tFound: '%v'\n\tExpected: '%v'\n\tError: %v\n", attributes[len(attributes)-1], xattrName, err)
        }
    }
}
