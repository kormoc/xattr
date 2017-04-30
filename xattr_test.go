package xattr

import "io/ioutil"
import "os"
import "reflect"
import "testing"

func TestHas(t *testing.T) {
    var test_xattrName = "user.xattr.test"
    var test_xattrValue = []byte{11,22,33,44,55,66,77,88,99}

    tmpfile, err := ioutil.TempFile("", "xattr_Test")
    if err != nil {
        t.Fatal(err)
    }

    defer os.Remove(tmpfile.Name())

    if has, err := Has(tmpfile.Name(), test_xattrName); has != false || err != nil {
        t.Fatalf("Has failed: %v\n", err)
    }

    if err := SetBytes(tmpfile.Name(), test_xattrName, test_xattrValue); err != nil {
        t.Fatalf("SetBytes failed: %v\n", err)
    }

    if has, err := Has(tmpfile.Name(), test_xattrName); has != true || err != nil {
        t.Fatalf("Has failed: %v\n", err)
    }

    if value, err := GetBytes(tmpfile.Name(), test_xattrName); !reflect.DeepEqual(value, test_xattrValue) {
        t.Fatalf("GetBytes failed: %v\n\tExpected: '%v'\n\tFound: '%v'\n", err, test_xattrValue, value)
    }

    if err := Remove(tmpfile.Name(), test_xattrName); err != nil {
        t.Fatalf("Remove failed: %v\n", err)
    }

    if has, err := Has(tmpfile.Name(), test_xattrName); has != false || err != nil {
        t.Fatalf("Has failed: %v\n", err)
    }
}
