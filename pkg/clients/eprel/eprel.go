package eprel

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func getEndpoint(path string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("clients.eprel"), path)
}

func GetEprelSpecifications(eprelCode string) (*TyreResponse, error) {
	response, err := http.Get(getEndpoint("products/tyres/" + eprelCode))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var tyreResponse TyreResponse
	err = json.NewDecoder(response.Body).Decode(&tyreResponse)
	if err != nil {
		return nil, err
	}

	return &tyreResponse, nil
}
