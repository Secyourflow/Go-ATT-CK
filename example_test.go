package attck_test

import (
	"fmt"

	"github.com/Secyourflow/attck"
)

func Example() {
	if tactic, ok := attck.Objects["TA0004"]; ok {
		fmt.Println(tactic.Name)
		fmt.Println(tactic.URL)
	}

	if technique, ok := attck.Objects["T1548.002"]; ok {
		fmt.Println(technique.Name)
		fmt.Println(technique.FullName)
		fmt.Println(technique.URL)
	}

	// Output:
	// Privilege Escalation
	// https://attack.mitre.org/tactics/TA0004
	// Bypass User Account Control
	// Abuse Elevation Control Mechanism: Bypass User Account Control
	// https://attack.mitre.org/techniques/T1548/002
}
