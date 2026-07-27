package application

import "fmt"

type WebViewProcessError struct {
	WindowName              string
	WindowID                uint
	RuntimeVersion          string
	Kind                    string
	Reason                  string
	ExitCode                int32
	ProcessDescription      string
	FailureReportFolderPath string
}

func (e *WebViewProcessError) Error() string {
	return fmt.Sprintf(
		"WebView2 process failed: window=%q id=%d runtime=%q kind=%s reason=%s exitCode=%d description=%q failureReportFolder=%q",
		e.WindowName,
		e.WindowID,
		e.RuntimeVersion,
		e.Kind,
		e.Reason,
		e.ExitCode,
		e.ProcessDescription,
		e.FailureReportFolderPath,
	)
}
