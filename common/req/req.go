package req

type JsonObj struct {
	WorkflowId   int    `json:"workflowId"`
	Creator      string `json:"creator"`
	BusinessId   int    `json:"businessId"`
	BusinessData string `json:"businessData"`
}
