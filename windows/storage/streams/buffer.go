// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package streams

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureBuffer string = "rc(Windows.Storage.Streams.Buffer;{905a0fe0-bc53-11df-8c49-001e4fc686da})"

type Buffer struct {
	ole.IUnknown
}

func (impl *Buffer) GetCapacity() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIBuffer))
	defer itf.Release()
	v := (*IBuffer)(unsafe.Pointer(itf))
	return v.GetCapacity()
}

func (impl *Buffer) GetLength() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIBuffer))
	defer itf.Release()
	v := (*IBuffer)(unsafe.Pointer(itf))
	return v.GetLength()
}

func (impl *Buffer) SetLength(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIBuffer))
	defer itf.Release()
	v := (*IBuffer)(unsafe.Pointer(itf))
	return v.SetLength(value)
}

const GUIDiBufferFactory string = "71af914d-c10f-484b-bc50-14bc623b3a27"
const SignatureiBufferFactory string = "{71af914d-c10f-484b-bc50-14bc623b3a27}"

type iBufferFactory struct {
	ole.IInspectable
}

type iBufferFactoryVtbl struct {
	ole.IInspectableVtbl

	BufferCreate uintptr
}

func (v *iBufferFactory) VTable() *iBufferFactoryVtbl {
	return (*iBufferFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}

func BufferCreate(capacity uint32) (*Buffer, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Storage.Streams.Buffer", ole.NewGUID(GUIDiBufferFactory))
	if err != nil {
		return nil, err
	}
	v := (*iBufferFactory)(unsafe.Pointer(inspectable))

	var out *Buffer
	hr, _, _ := syscall.SyscallN(
		v.VTable().BufferCreate,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(capacity),             // in uint32
		uintptr(unsafe.Pointer(&out)), // out Buffer
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
