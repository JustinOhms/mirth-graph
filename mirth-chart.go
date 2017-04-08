// mirth-chart project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/justinohms/mirth-chart/structs"
)

var srcDirP = flag.String("src", "/Users/justinohms/Dropbox/nc/src/salmon-mirth/src/mirth/channel-groups", "The directory containing the mirth source xml files.")

//var srcDirP = flag.String("src", "", "The directory containing the mirth source xml files.")

var channels = make(map[string]structs.MirthChannel)

func main() {
	flag.Parse()

	srcDir := ""
	fmt.Println("Mirth Chart")
	if *srcDirP == "" {
		srcDir, _ = os.Getwd()
	} else {
		srcDir = *srcDirP
	}
	fmt.Println("Source Directory:", srcDir)

	findAllXmlFiles(srcDir)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findAllXmlFiles(path string) {
	err := filepath.Walk(path, visit)
	check(err)
}

func visit(p string, f os.FileInfo, err error) error {
	//fmt.Printf("Visited: %s\n", p)

	//if it's not a dir and is an xml file we are interested
	if !f.IsDir() && strings.ToLower(path.Ext(p)) == ".xml" {
		fl, err := os.Open(p)
		defer fl.Close()
		check(err)
		//fmt.Printf("name: %s\n", f.Name())
		b1 := make([]byte, 25)
		fl.ReadAt(b1, 0)
		fl.Close()

		if strings.Contains(string(b1), "<channel version=") {
			fmt.Printf("CHANNEL %s\n", p)

			ch := structs.MirthChannel{
				FilePath: p,
			}

			channels[p] = ch

		}
		//fmt.Printf("%s\n", string(b1))
		fmt.Printf("%d\n", len(channels))
	}

	return nil
}
