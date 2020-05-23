package apiserver

// StatusResponse response in status request
type StatusResponse struct {
	Version    string `json:"version"`
	VersionAPI string `json:"version_api"`
	Build      string `json:"build"`
	Uptime     string `json:"uptime"`
}
