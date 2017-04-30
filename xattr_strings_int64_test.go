package xattr

import "io/ioutil"
import "os"
import "testing"

func TestStringsInt64(t *testing.T) {
    var test_xattrName = "user.xattr.test"
    var test_xattrValue = int64(1234567890)

    tmpfile, err := ioutil.TempFile("", "xattr_Test")
    if err != nil {
        t.Fatal(err)
    }

    defer os.Remove(tmpfile.Name())

    if err := SetStringInt64(tmpfile.Name(), test_xattrName, test_xattrValue); err != nil {
        t.Fatalf("SetStringInt64 failed: %v\n", err)
    }

    if value, err := GetStringInt64(tmpfile.Name(), test_xattrName); value != test_xattrValue {
        t.Fatalf("GetStringInt64 failed: %v\n\tExpected: '%v'\n\tFound: '%v'\n", err, test_xattrValue, value)
    }

    if err := Remove(tmpfile.Name(), test_xattrName); err != nil {
        t.Fatalf("Remove failed: %v\n", err)
    }
}
