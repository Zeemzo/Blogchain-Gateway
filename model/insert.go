package model

type InsertDataResponse struct {
	TDPID     string
	ProfileID string
	TxnType   string
	Error     Error
}

type CAResponse struct {
	PrivateKey     string
	PublicKey string
	Error     Error
}

type InsertGenesisResponse struct {
	ProfileTxn  string
	GenesisTxn  string
	Identifiers string
	TxnType     string
	Error       Error
}
