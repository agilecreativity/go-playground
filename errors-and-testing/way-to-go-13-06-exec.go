package main

import (
	"fmt"
	"os"
)

func showAllProcesses(cmd string, args []string, procAttr *os.ProcAttr) {
	pid, err := os.StartProcess(cmd, args, procAttr)

	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}

	fmt.Printf("The process id is %#v\n", pid)
}

func main() {
	env := os.Environ()

	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	showAllProcesses("/bin/ls", []string{"ls", "-l"}, procAttr)
	showAllProcesses("/bin/ps", []string{"aux"}, procAttr)
}

/** Output:

The process id is &os.Process{Pid:40743, handle:0x0, isdone:0x0}
The process id is &os.Process{Pid:40744, handle:0x0, isdone:0x0}
total 64
drwxr-xr-x  2 bch  staff   102  9 Jul 10:29 s13_04_parse
-rw-r--r--  1 bch  staff   575  9 Jul 10:29 way-to-go-13-01-01-user-defined.go
-rw-r--r--  1 bch  staff   586  9 Jul 10:29 way-to-go-13-01-02-fmt-errorf.go
-rw-r--r--  1 bch  staff   409  9 Jul 10:29 way-to-go-13-02-01-simple-panic.go
-rw-r--r--  1 bch  staff   242  9 Jul 10:29 way-to-go-13-02-02-panic-eg.go
-rw-r--r--  1 bch  staff   376  9 Jul 10:29 way-to-go-13-03-panic-recover.go
-rw-r--r--  1 bch  staff   695  9 Jul 10:29 way-to-go-13-05-panic-package.go
-rw-r--r--  1 bch  staff  1334  9 Jul 21:12 way-to-go-13-06-exec.go
-rw-r--r--  1 bch  staff   642  9 Jul 10:29 way-to-go-13-13-06-panic-defer.go

  PID TTY           TIME CMD
26807 ttys000    0:01.10 -/bin/zsh
36508 ttys000    0:37.21 /opt/boxen/homebrew/Cellar/macvim/HEAD/MacVim.app/Contents/MacOS/Vim way-to-go-13-06-exec.go
40738 ttys000    0:00.00 (go)
40743 ttys000    0:00.00 (ls)
40042 ttys001    0:00.19 -/bin/zsh
*/
