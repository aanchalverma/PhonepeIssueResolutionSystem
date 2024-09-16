/*
Constants to be used within the Issues resolution system
*/
package constants

// Enumerate the value of issue type
type IssueType string

const (
	IssueTypePayment    IssueType = "Payment"
	IssueTypeMutualFund IssueType = "Mutual Fund"
	IssueTypeGold       IssueType = "Gold"
	IssueTypeInsurance  IssueType = "Insurance"
)
