package ai

func ClassifyIssue(issue string) (urgency string, draftMsg string) {
	if len(issue) > 20 {
		return "EMERGENCY", "Urgent plumbing issue reported. Please attend ASAP."
	}
	return "ROUTINE", "Routine maintenance request received."
}
