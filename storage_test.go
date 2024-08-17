package winrt_test

import (
	"testing"

	"github.com/saltosystems/winrt-go/windows/storage"
	"golang.org/x/sys/windows"
)

func Test_GetStorageFolder(t *testing.T) {
	err := windows.CoInitializeEx(0, windows.COINIT_MULTITHREADED)
	if err != nil {
		t.Fatal(err)
	}
	op, err := storage.StorageFolderGetFolderFromPathAsync("c:")
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
