package winrt_test

import (
	"testing"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/storage"
)

func Test_GetStorageFolder(t *testing.T) {
	err := ole.RoInitialize(1)
	if err != nil {
		t.Fatal(err)
	}
	op, err := storage.StorageFolderGetFolderFromPathAsync("c:\\Users\\Grill Balázs")
	if err != nil {
		t.Fatal(err)
	}
	ptr, err := op.GetResults()
	if err != nil {
		t.Fatal(err)
	}
	if uintptr(ptr) == 0 {
		t.Fatal("got nil pointer")
	}
	//folder := (*storage.IStorageFolder)(unsafe.Pointer(ptr))

}
