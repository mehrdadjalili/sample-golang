package utils

import (
	"bytes"
	"encoding/json"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"net/http"
)

func SendEmail(email, code string) error {
	var apiKey = ""
	var fromEmail = "no_reply@jibito.org"

	posturl := "https://api.elasticemail.com/v4/emails/transactional"

	msg := "کد فعالسازی شما "
	msg = msg + code
	msg = msg + " میباشد. "

	data := map[string]interface{}{
		"Recipients": map[string]interface{}{
			"To": []string{email},
		},
		"Content": map[string]interface{}{
			"Subject": "کد فعالسازی",
			"From":    fromEmail,
			"Body": []map[string]interface{}{
				{
					"ContentType": "HTML",
					"Content":     msg,
					"Charset":     "string",
				},
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
	r.Header.Add("X-ElasticEmail-ApiKey", apiKey)

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
