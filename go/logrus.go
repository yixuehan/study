package main

import (
	"git.ucloudadmin.com/securityhouse/dataflow/pkg/log"
)

func main() {
	log.SetLevel("info")
	log.Infof("aha")
	// log.Infof("abc %s\n", "123")
	// log.Debugf("12dd %v\n", true)
	// log.Fatalf("fatal")
	// fmt.Printf("end main")
}
