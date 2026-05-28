//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2BasicAuthenticationResponseVtbl struct {
	_IUnknownVtbl
	GetUserName ComProc
	PutUserName ComProc
	GetPassword ComProc
	PutPassword ComProc
}

type ICoreWebView2BasicAuthenticationResponse struct {
	vtbl *_ICoreWebView2BasicAuthenticationResponseVtbl
}

func (i *ICoreWebView2BasicAuthenticationResponse) Release() uintptr {
	ret, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return ret
}

func (i *ICoreWebView2BasicAuthenticationResponse) PutUserName(userName string) error {
	_userName, err := windows.UTF16PtrFromString(userName)
	if err != nil {
		return err
	}
	hr, _, _ := i.vtbl.PutUserName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userName)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return windows.Errno(hr)
	}
	return nil
}

func (i *ICoreWebView2BasicAuthenticationResponse) PutPassword(password string) error {
	_password, err := windows.UTF16PtrFromString(password)
	if err != nil {
		return err
	}
	hr, _, _ := i.vtbl.PutPassword.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_password)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return windows.Errno(hr)
	}
	return nil
}
