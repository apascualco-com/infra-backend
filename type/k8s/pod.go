package k8s

type Pod struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Ready     string `json:"ready"`
	Status    string `json:"status"`
	Restarts  string `json:"restarts"`
	Date      string `json:"date"`
	HostIp    string `json:"host-ip"`
	PodIp     string `json:"pod-ip"`
}
