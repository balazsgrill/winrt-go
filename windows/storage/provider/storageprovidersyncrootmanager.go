// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package provider

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
	"github.com/saltosystems/winrt-go/windows/storage"
)

const GUIDiStorageProviderSyncRootManagerStatics2 string = "efb6cfee-1374-544e-9df1-5598d2e9cfdd"
const SignatureiStorageProviderSyncRootManagerStatics2 string = "{efb6cfee-1374-544e-9df1-5598d2e9cfdd}"

type iStorageProviderSyncRootManagerStatics2 struct {
	ole.IInspectable
}

type iStorageProviderSyncRootManagerStatics2Vtbl struct {
	ole.IInspectableVtbl

	StorageProviderSyncRootManagerIsSupported uintptr
}

func (v *iStorageProviderSyncRootManagerStatics2) VTable() *iStorageProviderSyncRootManagerStatics2Vtbl {
	return (*iStorageProviderSyncRootManagerStatics2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func StorageProviderSyncRootManagerIsSupported() (bool, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics2))
	if err != nil {
		return false, err
	}
	v := (*iStorageProviderSyncRootManagerStatics2)(unsafe.Pointer(inspectable))

	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerIsSupported,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiStorageProviderSyncRootManagerStatics string = "3e99fbbf-8fe3-4b40-abc7-f6fc3d74c98e"
const SignatureiStorageProviderSyncRootManagerStatics string = "{3e99fbbf-8fe3-4b40-abc7-f6fc3d74c98e}"

type iStorageProviderSyncRootManagerStatics struct {
	ole.IInspectable
}

type iStorageProviderSyncRootManagerStaticsVtbl struct {
	ole.IInspectableVtbl

	StorageProviderSyncRootManagerRegister                        uintptr
	StorageProviderSyncRootManagerUnregister                      uintptr
	StorageProviderSyncRootManagerGetSyncRootInformationForFolder uintptr
	StorageProviderSyncRootManagerGetSyncRootInformationForId     uintptr
	StorageProviderSyncRootManagerGetCurrentSyncRoots             uintptr
}

func (v *iStorageProviderSyncRootManagerStatics) VTable() *iStorageProviderSyncRootManagerStaticsVtbl {
	return (*iStorageProviderSyncRootManagerStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func StorageProviderSyncRootManagerRegister(syncRootInformation *StorageProviderSyncRootInfo) error {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics))
	if err != nil {
		return err
	}
	v := (*iStorageProviderSyncRootManagerStatics)(unsafe.Pointer(inspectable))

	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerRegister,
		uintptr(unsafe.Pointer(v)),                   // this
		uintptr(unsafe.Pointer(syncRootInformation)), // in StorageProviderSyncRootInfo
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func StorageProviderSyncRootManagerUnregister(id string) error {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics))
	if err != nil {
		return err
	}
	v := (*iStorageProviderSyncRootManagerStatics)(unsafe.Pointer(inspectable))

	idHStr, err := ole.NewHString(id)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerUnregister,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(idHStr),            // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func StorageProviderSyncRootManagerGetSyncRootInformationForFolder(folder *storage.IStorageFolder) (*StorageProviderSyncRootInfo, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics))
	if err != nil {
		return nil, err
	}
	v := (*iStorageProviderSyncRootManagerStatics)(unsafe.Pointer(inspectable))

	var out *StorageProviderSyncRootInfo
	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerGetSyncRootInformationForFolder,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(folder)), // in storage.IStorageFolder
		uintptr(unsafe.Pointer(&out)),   // out StorageProviderSyncRootInfo
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func StorageProviderSyncRootManagerGetSyncRootInformationForId(id string) (*StorageProviderSyncRootInfo, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics))
	if err != nil {
		return nil, err
	}
	v := (*iStorageProviderSyncRootManagerStatics)(unsafe.Pointer(inspectable))

	var out *StorageProviderSyncRootInfo
	idHStr, err := ole.NewHString(id)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerGetSyncRootInformationForId,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(idHStr),               // in string
		uintptr(unsafe.Pointer(&out)), // out StorageProviderSyncRootInfo
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func StorageProviderSyncRootManagerGetCurrentSyncRoots() (*collections.IVectorView, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Provider.StorageProviderSyncRootManager", ole.NewGUID(GUIDiStorageProviderSyncRootManagerStatics))
	if err != nil {
		return nil, err
	}
	v := (*iStorageProviderSyncRootManagerStatics)(unsafe.Pointer(inspectable))

	var out *collections.IVectorView
	hr, _, _ := syscall.SyscallN(
		v.VTable().StorageProviderSyncRootManagerGetCurrentSyncRoots,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVectorView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
