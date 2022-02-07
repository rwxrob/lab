# Find an executable anywhere, even Windows

Using the `exec.LookPath("VirtualBox.exe")` should be enough to
determine if VirtualBox has been installed on a Windows system whence
the Go program is executing (no matter what shell).
