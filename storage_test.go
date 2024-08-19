package winrt_test

import (
	"math"
	"runtime/debug"
	"testing"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage"
)

func Test_GetStorageFolder(t *testing.T) {
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)
	err := ole.RoInitialize(1)
	if err != nil {
		t.Fatal(err)
	}
	op, err := storage.StorageFolderGetFolderFromPathAsync("c:\\")
	if err != nil {
		t.Fatal(err)
	}
	semaphore := make(chan bool)
	iid := winrt.ParameterizedInstanceGUID(foundation.GUIDAsyncOperationCompletedHandler, storage.SignatureStorageFolder)
	handler := foundation.NewAsyncOperationCompletedHandler(ole.NewGUID(iid), func(instance *foundation.AsyncOperationCompletedHandler, asyncInfo *foundation.IAsyncOperation, asyncStatus foundation.AsyncStatus) {
		semaphore <- true
	})
	err = op.SetCompleted(handler)
	if err != nil {
		t.Fatal(err)
	}
	<-semaphore
	ptr, err := op.GetResults()
	if err != nil {
		t.Fatal(err)
	}
	if uintptr(ptr) == 0 {
		t.Fatal("got nil pointer")
	}
	//folder := (*storage.IStorageFolder)(unsafe.Pointer(ptr))

}
