package models

type Currency struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

type SymbolData struct {
	ID           string `json:"id"`
	BaseCurrency string `json:"baseCurrency"`
	FeeCurrency  string `json:"feeCurrency"`
}

type CurrencyData struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

// TradeData represents the structure of our documents in MongoDB
type TradeData struct {
	ID          string `bson:"_id,omitempty"`
	Symbol      string `bson:"symbol"`
	SymbolID    string `bson:"symbolID"`
	FeeCurrency string `bson:"feeCurrency"`
	FullName    string `bson:"fullName"`
}
