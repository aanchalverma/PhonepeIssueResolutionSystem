/*
Issues package : Defines features and attributes of an issue/complaints
*/

package issues

import (
	"fmt"

	"github.com/aanchalverma/Machine-Coding/PhonepeIssueResolutionSystem/constants"
)

type Issue struct {
	TransactionId string
	IssueType     constants.IssueType
	// TODO: Check if you can validate email and define range of characters to be accepted for subject/description and an throw error if input fails to validate
	Subject      string
	Description  string
	Email        string
	AgentHandled string
	Status       string
}

var Issues []Issue

func AddIssue(issue Issue) {
	fmt.Println("Adding issue to the queue")
	Issues = append(Issues, issue)
	fmt.Println("Issue", issue.TransactionId, "added to the queue")
}

func FindIssueWithTransactionId(transactionId string) (Issue, bool) {
	for _, issue := range Issues {
		if issue.TransactionId == transactionId {
			return issue, true
		}
	}
	return Issue{}, false
}

func GetIssue(filter string) {
	found := false
	for _, issue := range Issues {
		if filter == issue.TransactionId || filter == issue.Subject || filter == issue.Description || filter == issue.Email {
			fmt.Println("Details of the issue", filter)
			fmt.Println(issue.TransactionId)
			fmt.Println(issue.Description)
			fmt.Println(issue.Email)
			fmt.Println(issue.IssueType)
			fmt.Println(issue.AgentHandled)
			fmt.Println(issue.Status)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Could not find the issue", filter)
	}
}
