package values

var (
	Name            = "myapp"
	Namespace       = "csip-dns"
	Replicas  int32 = 1
	Labels          = map[string]string{
		"app": "myapp",
	}
	Image = "nginx:latest"
)
