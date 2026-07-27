//go:build windows

package edge

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2ProcessFailedEventArgs2Vtbl struct {
	_ICoreWebView2ProcessFailedEventArgsVtbl
	GetReason                     ComProc
	GetExitCode                   ComProc
	GetProcessDescription         ComProc
	GetFrameInfosForFailedProcess ComProc
}

type ICoreWebView2ProcessFailedEventArgs2 struct {
	vtbl *iCoreWebView2ProcessFailedEventArgs2Vtbl
}

func (i *ICoreWebView2ProcessFailedEventArgs) GetICoreWebView2ProcessFailedEventArgs2() *ICoreWebView2ProcessFailedEventArgs2 {
	var result *ICoreWebView2ProcessFailedEventArgs2
	iid := NewGUID("{4dab9422-46fa-4c3e-a5d2-41d2071d3680}")
	hr, _, _ := i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&result)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil
	}
	return result
}

func (i *ICoreWebView2ProcessFailedEventArgs2) Release() uint32 {
	result, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return uint32(result)
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetReason() (COREWEBVIEW2_PROCESS_FAILED_REASON, error) {
	var reason COREWEBVIEW2_PROCESS_FAILED_REASON
	hr, _, _ := i.vtbl.GetReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return reason, nil
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetExitCode() (int32, error) {
	var exitCode int32
	hr, _, _ := i.vtbl.GetExitCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&exitCode)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return exitCode, nil
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetProcessDescription() (string, error) {
	var description *uint16
	hr, _, _ := i.vtbl.GetProcessDescription.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&description)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	defer CoTaskMemFree(unsafe.Pointer(description))
	return UTF16PtrToString(description), nil
}
