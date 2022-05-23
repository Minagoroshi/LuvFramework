<p align="center">
  <img width="400" height="400" src="https://i.ibb.co/Hx1YH7h/optimizedluv.gif">
</p>

## 🚧 THIS FRAMEWORK IS UNDER DEVELOPMENT 🚧

## 🔎💵 LOOKING FOR COLLABORATORS AND MAINTAINERS, depending on the amount of work you offer I may even pay you! See [Contributing](#contributing) 💵🔎


# Luv Framework
Luv Framework is a framework for developing malware and pen-testing software to be used by researchers specifically targeting Windows machines



## Table of Contents
- [Table of Contents](#table-of-contents)
- [Installation](#installation)
- [Features](#features)
- [Docs](#documentation)
  - [Evasion](#evasion)
  - [Recon](#recon)
  - [Exploitations](#exploitation)
  - [Persistence](#persistence)
  - [Anti-Sandbox](#anti-sandbox)
  - [Anti-Memory](#anti-memory)
  - [Anti-Debugging](#anti-debugging)
  - [Third-Party](#third-party)
    - [Discord](#discord)
  - [Wireless](#wireless)
    - [Credential Stuffing Library](#credential-stuffing)
  - [Toolchain](#toolchain)
- [Upcoming Features](#upcoming-features)
- [Contributing](#contributing)
- [Disclaimer](#disclaimer)
- [License](#license)

## Installation
`go get github.com/Minagoroshi/LuvFramework`

## Features
- [x] Evasion functions
- [x] Reconnaissance Functions
- [x] Exploitation functions
- [x] Anti Sandbox functions
- [x] Anti Memory Inspection functions
- [x] Anti Debugging functions

## Documentation
### Evasion
- HopBypass (Manipulates machine memory to evade AVs)

### Recon
- GetDrives (Returns all drives on the machine)
- CheckRoot (Checks if the current user is root)
- GetUsers (Returns a list of users on the machine)
- ListFiles (Returns a list of files to the provided path)
- GetRegKeyData (Returns a list of registry values to the provided key)
- GetFileData (Returns a list of file data to the provided path)
- GetFileSize (Returns the size of the file)
- MachineGuid (Returns the machine's GUID)
- RegKeyData (Returns a list of registry values to the provided key)
- GetWindowsVersion (Returns the current Windows version as an integer)

### Exploitation
- ForkBomb (Exponential growth of goroutines to deplete system resources)
- NetworkDisconnect (Disconnects network connections)
- Wipe (Wipes the filesystem of the provided path)
- ClearLogs (Clears windows logs)
- Shutdown (Shuts down the machine)
- SluiBypass (UAC bypass using slui.exe)
- ComputerDefaults (UAC bypass using computerdefaults.exe)
- SdcltControl (UAC bypass using sdclt.exe)
- FodHelper (UAC bypass using fodhelper.exe)
- TryEscalate (Attempts to escalate privileges, returns true if successful)

### Persistence
- [ ] Persistence Functions (TODO)

### Anti Sandbox
- [ ] Anti Sandbox functions (TODO)

### Anti Memory
- [ ] Anti Memory Inspection functions (TODO)

### Anti Debugging
- [ ] Anti Debugging functions (TODO)

### Third-Party
#### Discord
- TokenLogEvasive (Logs discord tokens to a discord webhook, uses function name obfuscation and encoding to evade AVs)

### Wireless
Various functions to assist in penetration testing of wireless networks

#### Credential Stuffing
## Usage

```go
package main

import "github.com/Minagoroshi/LuvFramework"

func main() {
w := luvcheckerlib.NewWorker()
w.AddWork(checker)
w.StartWorker()
}

func checker() {
// Your Code Here
}
```
## ToolChain
- [ ] ToolChain (TODO)

## Upcoming Features
- [ ] Persistence Functions (TODO)
- [ ] Anti Sandbox functions (TODO)
- [ ] Anti Memory Inspection functions (TODO)
- [ ] Anti Debugging functions (TODO)
- [ ] Wireless functions (TODO)
- [ ] Third Party exploits (TODO)
- [ ] Toolchain  (TODO)

## Contributing
Any contributions are welcome! Feel free to fork and submit pull requests.
Please make sure to follow our [Guidelines](CONTRIBUTING.md)

I have a VERY busy lifestyle with a fulltime job, and college. 
Any help at all would be very appreciated, if you would like to be part of the team and become a collaborator please message me on discord `tawp#0001` 
depending on the amount of effort put in I may even offer payment, but don't expect a job out of it.



## Disclaimer
This software is distributed in the hope that it will be useful, the usage of this software is at your own risk.
Any damage caused by this software is not the responsibility of the author.
This software is provided "as is" with no guarantees or warranties of any kind.

In regard to the [Third Party](#third-party) functions, the author is not responsible for any damages caused by the use of these functions.
They should only be used on your own machines and accounts that you have created yourself. Any damage caused by these functions is the responsibility of the user.
These are solely to be used for educational purposes in hopes to prompt an exploit to be discovered and resolved.

## License

This software is licensed under the GPLv3 license.
