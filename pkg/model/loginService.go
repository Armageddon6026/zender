package model

type LoginService struct {
	Uuid             string                     `json:"uuid,omitempty"`
	Name             string                     `json:"name,omitempty"`
	Ip               string                     `json:"ip,omitempty"`
	ScanFiles        map[string]ScanFile        `json:"scanFiles,omitempty"`
	ScanApplications map[string]ScanApplication `json:"scanApplications,omitempty"`
}

type ScanFile struct {
	Path         string `json:"path,omitempty"`
	DataCount    int32  `json:"dataCount,omitempty"`
	LastDataTime int64  `json:"lastDataTime,omitempty"`
}

type ScanApplication struct {
	Name         string `json:"Name,omitempty"`
	LastDataTime int64  `json:"lastDataTime,omitempty"`
}

type LoginServiceList []LoginService
