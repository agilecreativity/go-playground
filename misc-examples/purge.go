// Mimic the bash script 'find /home/user/tmp -mtime +90 exec rm -f {} \;'
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var path string // path to purge
var days int    // age, in days, to purge
var test bool   // test mode toggle

// isDir accepts a string (file path) and returns
// a boolean which indicates if the path is
// a valid directory.
func isDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return stat.IsDir()
}

func init() {
	// Set up command-line flags.
	flag.StringVar(&path, "p", "", "path to purge")
	flag.IntVar(&days, "d", -1, "number of days")
	flag.BoolVar(&test, "t", false, "test mode -- don't delete and log to stdout")
	flag.Parse()
	// Validate flags.
	if path == "" {
		log.Fatal("No path passed.")
	}
	if !isDir(path) {
		log.Fatal(path, "is not a valid path.")
	}
	if days < 1 {
		log.Fatal("Number of days must be a positive, non-zero number. This is for your own protection.")
	}
	// Set up logging.
	if test {
		log.Println("Test mode. Will not delete anything.")
		return
	}
	logfile := fmt.Sprintf("purge_%s.log", time.Now().Format("2006-01-02_150405"))
	logPath := filepath.Join(os.TempDir(), logfile)
	writer, err := os.Create(logPath)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(writer)
}

// walker implements filepath.WalkFunc.
func walker(path string, info os.FileInfo, err error) error {
	cutoff, err := time.ParseDuration(fmt.Sprintf("%dh", days*24))
	if err != nil {
		log.Fatal(err)
	}
	if time.Now().Sub(info.ModTime()) < cutoff {
		return nil
	}
	// Don't delete a directory unless it's empty.
	if info.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		if len(files) > 0 {
			log.Printf("directory %s contains files; skipping\n", path)
			return nil
		}
		log.Printf("directory %s is empty\n", path)
	}
	log.Printf("deleting %s\n", path)
	if !test {
		err = os.Remove(path)
	}
	return nil
}

func main() {
	filepath.Walk(path, walker)
}
