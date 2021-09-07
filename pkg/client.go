package pkg

type IClient interface {
	Register() error
	Discovery(serName string) ([]string, error)
	Destroy() error
}

type Client struct {
	InsClient IClient
}

func NewClient(instance IClient) *Client {
	return &Client{
		InsClient: instance,
	}
}

func (this *Client) Register() error {
	return this.InsClient.Register()
}

func (this *Client) Discovery(serName string) ([]string, error) {
	return this.InsClient.Discovery(serName)
}

func (this *Client) Destroy() error {
	return this.InsClient.Destroy()
}
