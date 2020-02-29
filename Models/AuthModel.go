package Models

type Exception struct {
	Message string `json:"message"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Auth struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}
