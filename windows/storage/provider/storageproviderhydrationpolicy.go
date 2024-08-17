// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package provider

type StorageProviderHydrationPolicy int32

const SignatureStorageProviderHydrationPolicy string = "enum(Windows.Storage.Provider.StorageProviderHydrationPolicy;i4)"

const (
	StorageProviderHydrationPolicyPartial     StorageProviderHydrationPolicy = 0
	StorageProviderHydrationPolicyProgressive StorageProviderHydrationPolicy = 1
	StorageProviderHydrationPolicyFull        StorageProviderHydrationPolicy = 2
	StorageProviderHydrationPolicyAlwaysFull  StorageProviderHydrationPolicy = 3
)
