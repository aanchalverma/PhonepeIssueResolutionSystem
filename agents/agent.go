/*
Agent package : Defines features and attributes of agent
*/
package agents

import (
	"fmt"

	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/constants"
	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/issues"
)

type AgentAvailable struct {
	AgentName         string
	AgentEmail        string
	ExpertiseInIssues []constants.IssueType
	WorkHistory       []issues.Issue
}

type AgentOccupied struct {
	AgentName         string
	AgentEmail        string
	ExpertiseInIssues []constants.IssueType
	WorkHistory       []issues.Issue
}

var Agents []AgentAvailable
var AgentsOccupied []AgentOccupied

func AddAgent(agent AgentAvailable) {
	// Add a new agent to the agents queue
	Agents = append(Agents, agent)
}

func AssignIssue(transactionId string) {
	var newAgentList []AgentAvailable
	var agentAssigned AgentOccupied
	indexOfAssignedAgent := -1
	// Find the issue corresponding to transactionId
	issueToBeResolved, isIssueFound := issues.FindIssueWithTransactionId(transactionId)
	if !isIssueFound {
		fmt.Errorf("Issue with transaction id %s could not be found", transactionId)
	} else {
		// If issue found, assign an agent
		for i, agent := range Agents {
			for _, issueType := range agent.ExpertiseInIssues {
				// Check if the issue type matches with agent expertise
				if issueType == issueToBeResolved.IssueType {
					// Assign the issue to the first available agent
					agentAssigned = AgentOccupied(agent)
					AgentsOccupied = append(AgentsOccupied, agentAssigned)
					indexOfAssignedAgent = i
					// TODO: Change this to pass by reference to update the actual issue object
					issueToBeResolved.AgentHandled = agentAssigned.AgentName
					issueToBeResolved.Status = "Resolving"
					break
				}
			}
		}
		// Remove agent from available queue and move to occupied queue
		newAgentList = append(newAgentList, Agents[:indexOfAssignedAgent]...)
		newAgentList = append(newAgentList, Agents[indexOfAssignedAgent+1:]...)
		Agents = newAgentList
		if indexOfAssignedAgent == -1 {
			fmt.Println("No agents available currently. Please wait....")
		} else {
			fmt.Println("Agent", agentAssigned.AgentName, "has been assigned for the issue", transactionId)
		}
	}
}
