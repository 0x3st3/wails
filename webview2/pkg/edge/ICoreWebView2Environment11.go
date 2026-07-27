//go:build windows

package edge

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2Environment2Vtbl struct {
	iCoreWebView2EnvironmentVtbl
	CreateWebResourceRequest ComProc
}

type iCoreWebView2Environment3Vtbl struct {
	iCoreWebView2Environment2Vtbl
	CreateCoreWebView2CompositionController ComProc
	CreateCoreWebView2PointerInfo           ComProc
}

type iCoreWebView2Environment4Vtbl struct {
	iCoreWebView2Environment3Vtbl
	GetAutomationProviderForWindow ComProc
}

type iCoreWebView2Environment5Vtbl struct {
	iCoreWebView2Environment4Vtbl
	AddBrowserProcessExited    ComProc
	RemoveBrowserProcessExited ComProc
}

type iCoreWebView2Environment6Vtbl struct {
	iCoreWebView2Environment5Vtbl
	CreatePrintSettings ComProc
}

type iCoreWebView2Environment7Vtbl struct {
	iCoreWebView2Environment6Vtbl
	GetUserDataFolder ComProc
}

type iCoreWebView2Environment8Vtbl struct {
	iCoreWebView2Environment7Vtbl
	AddProcessInfosChanged    ComProc
	RemoveProcessInfosChanged ComProc
	GetProcessInfos           ComProc
}

type iCoreWebView2Environment9Vtbl struct {
	iCoreWebView2Environment8Vtbl
	CreateContextMenuItem ComProc
}

type iCoreWebView2Environment10Vtbl struct {
	iCoreWebView2Environment9Vtbl
	CreateCoreWebView2ControllerOptions                ComProc
	CreateCoreWebView2ControllerWithOptions            ComProc
	CreateCoreWebView2CompositionControllerWithOptions ComProc
}

type iCoreWebView2Environment11Vtbl struct {
	iCoreWebView2Environment10Vtbl
	GetFailureReportFolderPath ComProc
}

type ICoreWebView2Environment11 struct {
	vtbl *iCoreWebView2Environment11Vtbl
}

func (i *ICoreWebView2Environment) GetICoreWebView2Environment11() *ICoreWebView2Environment11 {
	var result *ICoreWebView2Environment11
	iid := NewGUID("{f0913dc6-a0ec-42ef-9805-91dff3a2966a}")
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

func (i *ICoreWebView2Environment11) Release() uint32 {
	result, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return uint32(result)
}

func (i *ICoreWebView2Environment11) GetFailureReportFolderPath() (string, error) {
	var path *uint16
	hr, _, _ := i.vtbl.GetFailureReportFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&path)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	defer CoTaskMemFree(unsafe.Pointer(path))
	return UTF16PtrToString(path), nil
}
