// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package collections

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIMap string = "3c2925fe-8519-45c1-aa79-197b6718c1c1"
const SignatureIMap string = "{3c2925fe-8519-45c1-aa79-197b6718c1c1}"

type IMap struct {
	ole.IInspectable
}

type IMapVtbl struct {
	ole.IInspectableVtbl

	Lookup  uintptr
	GetSize uintptr
	HasKey  uintptr
	GetView uintptr
	Insert  uintptr
	Remove  uintptr
	Clear   uintptr
}

func (v *IMap) VTable() *IMapVtbl {
	return (*IMapVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMap) Lookup(key unsafe.Pointer) (unsafe.Pointer, error) {
	var out unsafe.Pointer
	hr, _, _ := syscall.SyscallN(
		v.VTable().Lookup,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(key),                  // in unsafe.Pointer
		uintptr(unsafe.Pointer(&out)), // out unsafe.Pointer
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IMap) GetSize() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *IMap) HasKey(key unsafe.Pointer) (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().HasKey,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(key),                  // in unsafe.Pointer
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *IMap) Insert(key unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().Insert,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(key),                  // in unsafe.Pointer
		uintptr(value),                // in unsafe.Pointer
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *IMap) Remove(key unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Remove,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(key),               // in unsafe.Pointer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IMap) Clear() error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Clear,
		uintptr(unsafe.Pointer(v)), // this
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
