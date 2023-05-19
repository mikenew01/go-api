package services

type ClientePontoVendaCadastro struct {
	MemberNumber string `json:"memberNumber" validate:"required,max=30"`
	OptinStix    string `json:"optinStix" validate:"required,max=30"`
	Loja         int    `json:"loja"`
	PinCode      string `json:"pinCode" validate:"required,min=4,max=4"`
}

func CadastrarCliente(cliente *ClientePontoVendaCadastro) (ClientePontoVendaCadastro, error) {
	return ClientePontoVendaCadastro{
		MemberNumber: cliente.MemberNumber,
		OptinStix:    cliente.OptinStix,
		Loja:         cliente.Loja,
		PinCode:      cliente.PinCode,
	}, nil
}
