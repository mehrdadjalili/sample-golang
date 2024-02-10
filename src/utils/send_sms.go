package utils

import (
	"bytes"
	"encoding/json"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"net/http"
)

func SendSms(mobile, code string) error {
	var apiKey = ""
	posturl := "https://api.sms.ir/v1/send/verify"

	data := map[string]interface{}{
		"Mobile":     mobile,
		"TemplateId": "100000",
		"Parameters": []map[string]string{
			{
				"Name":  "CODE",
				"Value": code,
			},
		},
	}

	body, err := json.Marshal(data)
	if err != nil {
		return derrors.InternalError()
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return derrors.InternalError()
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-API-KEY", apiKey)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return derrors.InternalError()
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return derrors.InternalError()
	}

	return nil
}
