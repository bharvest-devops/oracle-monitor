package server

type Response struct {
	UmeeOracle struct {
		Status     bool            `json:"status"`
		AcceptList map[string]bool `json:"accept-list"`
		Window     string          `json:"window (current window / window)"`
		Uptime     string          `json:"uptime (uptime / minimum uptime)"`
	} `json:"umee-oracle"`
}
