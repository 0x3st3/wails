//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri       ComProc
	GetChallenge ComProc
	GetResponse  ComProc
	GetCancel    ComProc
	PutCancel    ComProc
	GetDeferral  ComProc
}

type ICoreWebView2BasicAuthenticationRequestedEventArgs struct {
	vtbl *_ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetUri() (string, error) {
	var _uri *uint16
	hr, _, _ := i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", windows.Errno(hr)
	}
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetChallenge() (string, error) {
	var _challenge *uint16
	hr, _, _ := i.vtbl.GetChallenge.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_challenge)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", windows.Errno(hr)
	}
	challenge := windows.UTF16PtrToString(_challenge)
	windows.CoTaskMemFree(unsafe.Pointer(_challenge))
	return challenge, nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) PutCancel(cancel bool) error {
	var _cancel uintptr
	if cancel {
		_cancel = 1
	}
	hr, _, _ := i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		_cancel,
	)
	if windows.Handle(hr) != windows.S_OK {
		return windows.Errno(hr)
	}
	return nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetResponse() (*ICoreWebView2BasicAuthenticationResponse, error) {
	var response *ICoreWebView2BasicAuthenticationResponse
	hr, _, _ := i.vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, windows.Errno(hr)
	}
	return response, nil
}
