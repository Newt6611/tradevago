package bitopro

type BitoClient struct {

}

func NewBitoClient() *BitoClient {
    return &BitoClient{
    }
}

func (c *BitoClient) GetName() string {
    return "BitoPro"
}
