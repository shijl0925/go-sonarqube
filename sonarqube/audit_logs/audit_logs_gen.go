package audit_logs

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DownloadRequest Returns security related audits of this SonarQube instance. Logs are returned in JSON format.<br/>Requires the system administration permission<br/>
type DownloadRequest struct {
	From string `url:"from"` // Date in ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hh:mm) from which the logs will be returned. Inclusive.
	To   string `url:"to"`   // Date in ISO 8601 datetime format (YYYY-MM-DDThh:mm:ss±hh:mm) until which the logs will be returned. Inclusive.
}

// DownloadResponse is the response for DownloadRequest
type DownloadResponse struct {
	AuditLogs []struct {
		Category  string `json:"category,omitempty"`
		CreatedAt string `json:"createdAt,omitempty"`
		NewValue  struct {
			ComponentKey  string `json:"componentKey,omitempty"`
			ComponentName string `json:"componentName,omitempty"`
			ComponentUuid string `json:"componentUuid,omitempty"`
			IsEnabled     bool   `json:"isEnabled,omitempty"`
			Qualifier     string `json:"qualifier,omitempty"`
		} `json:"newValue,omitempty"`
		Operation     string `json:"operation,omitempty"`
		UserLogin     string `json:"userLogin,omitempty"`
		UserTriggered bool   `json:"userTriggered,omitempty"`
		UserUuid      string `json:"userUuid,omitempty"`
		PreviousValue struct {
			TokenName string `json:"tokenName,omitempty"`
			UserLogin string `json:"userLogin,omitempty"`
			UserUuid  string `json:"userUuid,omitempty"`
		} `json:"previousValue,omitempty"`
	} `json:"audit_logs,omitempty"`
}
