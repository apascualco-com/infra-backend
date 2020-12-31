package k8s

type Node struct {
	Name             string `json:"name"`
	Status           string `json:"status"`
	Roles            string `json:"roles"`
	Date             string `json:"date"`
	Version          string `json:"version"`
	Os               string `json:"os"`
	KernelVersion    string `json:"kernel-version"`
	ContainerRuntime string `json:"container-runtime"`
	Architecture     string `json:"architecture"`
}
