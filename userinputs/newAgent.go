package inputs

import (
	"fmt"

	agents "github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/agents"
	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/constants"
	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/issues"
)

func CreateAgentsFromInputs() {
	// TODO: Let the user pass custom input for creating agents
	newAgent := agents.AgentAvailable{
		AgentName:  "Tushar",
		AgentEmail: "tushar@gmail.com",
		ExpertiseInIssues: []constants.IssueType{
			constants.IssueTypeGold,
			constants.IssueTypeInsurance,
		},
	}

	agents.AddAgent(newAgent)
	fmt.Printf("Agent %s has been added successfully", newAgent.AgentName)

	newAgent = agents.AgentAvailable{
		AgentName:  "Tanvi",
		AgentEmail: "tanvi@gmail.com",
		ExpertiseInIssues: []constants.IssueType{
			constants.IssueTypeMutualFund,
		},
	}

	agents.AddAgent(newAgent)
	fmt.Printf("Agent %s has been added successfully", newAgent.AgentName)

	newAgent = agents.AgentAvailable{
		AgentName:  "Sahil",
		AgentEmail: "sahil@gmail.com",
		ExpertiseInIssues: []constants.IssueType{
			constants.IssueTypePayment,
		},
	}

	agents.AddAgent(newAgent)

	fmt.Printf("Agent %s has been added successfully", newAgent.AgentName)
}

func CreateIssueFromInputs() {
	// TODO: Let the user pass custom input for creating issue
	newIssue := issues.Issue{
		TransactionId: "T1",
		IssueType:     constants.IssueTypeGold,
		Subject:       "Gold related issue",
		Description:   "Unable to access locker for gold",
		Email:         "bhavna@gmail.com",
	}
	issues.AddIssue(newIssue)

	newIssue = issues.Issue{
		TransactionId: "T2",
		IssueType:     constants.IssueTypeInsurance,
		Subject:       "Insurance related issue",
		Description:   "Unable to access INSURANCE",
		Email:         "bhaKTI@gmail.com",
	}
	issues.AddIssue(newIssue)

	newIssue = issues.Issue{
		TransactionId: "T3",
		IssueType:     constants.IssueTypeMutualFund,
		Subject:       "MutualFund related issue",
		Description:   "Unable to access MutualFund",
		Email:         "swati@gmail.com",
	}
	issues.AddIssue(newIssue)

	newIssue = issues.Issue{
		TransactionId: "T4",
		IssueType:     constants.IssueTypePayment,
		Subject:       "Payment related issue",
		Description:   "Unable to access Payment",
		Email:         "swati@gmail.com",
	}
	issues.AddIssue(newIssue)
}
