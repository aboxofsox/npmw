### Watch a directory and run NPM scripts when changes are detected.

#### Install
`go install github.com/aboxofsox/npmw`

*or*

Download the binary for your system from [releases](https://github.com/aboxofsox/npmw/releases).

#### Building
**Linux**

`./scripts/build.sh`

**Windows**

`./scripts/build.ps1`

#### Usage
`npwm watch --script <string> --root <string>`

`npmw watch -s build -r .`