package req

type JsonObj struct {
	WorkflowId   int    `json:"workflowId"`
	Creator      string `json:"creator"` //创建人
	BusinessId   int    `json:"businessId"`
	BusinessData string `json:"businessData"`
	InstanceId   int    `json:"instanceId"`
	Assignee     string `json:"assignee"` //审核人

}
