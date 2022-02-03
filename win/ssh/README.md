# Run SSH Server on Windows

Since you cannot mount a USB into WSL2 I thought I'd get an SSH server
running on it and just use the ssh form of Git URLs to access the
private repo on it.

> lbgdn: Ok, the trick is to put the public key in
> `%PROGRAMDATA%\ssh\administrators_authorized_keys` and change security
> on it to only SYSTEM and Administrators full access. 

You might have to disable inherited permissions on the file in order to
remove the "Authenticated User" permission.

* OpenSSH Server configuration for Windows \| Microsoft Docs  
  <https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_server_configuration>
* Key-based Authentication for OpenSSH on Windows - Concurrency  
  <https://www.concurrency.com/blog/may-2019/key-based-authentication-for-openssh-on-windows>
