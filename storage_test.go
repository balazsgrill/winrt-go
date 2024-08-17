package winrt_test

import (
	"math"
	"runtime/debug"
	"testing"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/storage"
)

func Test_GetStorageFolder(t *testing.T) {
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)
	err := ole.RoInitialize(1)
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
