package main

import (
	"fmt"

	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/agents"
	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/issues"
	inputs "github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/userinputs"
)

func main() {
	fmt.Println("Welcome to Phonepe Customer Resolution System")

	// Firstly, create agents who would handle the issues
	inputs.CreateAgentsFromInputs()
	inputs.CreateIssueFromInputs()

	// Assign a given issue to first available agent
	agents.AssignIssue("T1")

	// Issue against a filter
	issues.GetIssue("T1")

	// TODO: Implement updateIssues()
	// TODO: Implement resolveIssues()
	// TODO: Implement viewAgentsWorkHistory()
}
