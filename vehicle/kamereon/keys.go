package kamereon

import "fmt"

const keyStore = "https://renault-wrd-prod-1-euw1-myrapp-one.s3-eu-west-1.amazonaws.com/configuration/android/config_%s.json"

type configResponse struct {
	Servers configServers
}

type configServers struct {
	GigyaProd configServer `json:"gigyaProd"`
	WiredProd configServer `json:"wiredProd"`
}

type configServer struct {
	Target string `json:"target"`
	APIKey string `json:"apikey"`
}

func Keys(region string) error {
	uri := fmt.Sprintf(keyStore, region)

	var cr configResponse
	err := v.GetJSON(uri, &cr)
	if err != nil {
		// Use of old fixed keys if keyStore is not accessible
		v.gigya = configServer{"https://accounts.eu1.gigya.com", "3_7PLksOyBRkHv126x5WhHb-5pqC1qFR8pQjxSeLB6nhAnPERTUlwnYoznHSxwX668"}
		v.kamereon = configServer{"https://api-wired-prod-1-euw1.wrd-aws.com", "VAX7XYKGfa92yMvXculCkEFyfZbuM7Ss"}
		return nil
	} else {
		v.gigya = cr.Servers.GigyaProd
		v.kamereon = cr.Servers.WiredProd
		// Temporary fix of wrong kamereon APIKey in keyStore
		v.kamereon.APIKey = "VAX7XYKGfa92yMvXculCkEFyfZbuM7Ss"
		return err
	}
}
