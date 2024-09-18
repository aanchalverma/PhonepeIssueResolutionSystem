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
	fmt.Println("agent his", agents.AgentsOccupied)

	// Issue against a filter
	issues.GetIssue("T1")

	// Update issue with status and resolution
	issues.UpdateIssue("T1", "Resolved", "Able to access now")

	// Resolve an issue
	agents.ResolveIssue("T1", "Able to access now")

	// TODO: Implement viewAgentsWorkHistory()
	agents.ViewAgentsWorkHistory("Tushar")
}
