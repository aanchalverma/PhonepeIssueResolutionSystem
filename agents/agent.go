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
					agent.WorkHistory = append(agent.WorkHistory, *issueToBeResolved)
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

func ViewAgentsWorkHistory(agentName string) {
	found := false
	for _, agent := range Agents {
		if agent.AgentName == agentName {
			if len(agent.WorkHistory) == 0 {
				fmt.Println("No work history found for agent", agentName)
				found = true
				break
			}
			fmt.Printf("Here's the work history of agent %s", agentName)
			for _, workhistory := range agent.WorkHistory {
				fmt.Println(workhistory.IssueType)
				fmt.Println(workhistory.Subject)
				fmt.Println(workhistory.Description)
				fmt.Println(workhistory.AgentHandled)
				fmt.Println(workhistory.Status)
				fmt.Println(workhistory.Resolution)
				found = true
				break
			}
		}
	}

	for _, agent := range AgentsOccupied {
		if agent.AgentName == agentName {
			if len(agent.WorkHistory) == 0 {
				fmt.Println("No work history found for agent", agentName)
				found = true
				break
			}
			fmt.Printf("Here's the work history of agent %s", agentName)
			for _, workhistory := range agent.WorkHistory {
				fmt.Println(workhistory.IssueType)
				fmt.Println(workhistory.Subject)
				fmt.Println(workhistory.Description)
				fmt.Println(workhistory.AgentHandled)
				fmt.Println(workhistory.Status)
				fmt.Println(workhistory.Resolution)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("No agent found with name", agentName)
	}
}

func ResolveIssue(transactionId string, resolution string) {
	issueToBeUpdated, ifIssueFound := issues.FindIssueWithTransactionId(transactionId)
	if !ifIssueFound {
		fmt.Errorf("Issue with id %s could not be found", transactionId)
	} else {
		issueToBeUpdated.Status = "Resolved"
		issueToBeUpdated.Resolution = resolution
		for _, agent := range AgentsOccupied {
			if agent.AgentName == issueToBeUpdated.AgentHandled {
				// Put agent back in the queue
				Agents = append(Agents, AgentAvailable(agent))
			}
		}
		fmt.Printf("Issue with id %s has been successfully resolved with resolution - %s", transactionId, resolution)
		fmt.Println()
	}
}
