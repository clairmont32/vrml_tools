package vrml

import (
	"encoding/json"
	"github.com/clairmont32/httplib"
	"github.com/sirupsen/logrus"
)

func GetSponsors(h []httplib.Headers) (sponsors Sponsors, err error) {
	req := &httplib.FormRequest{
		BaseURL:  "https://api.vrmasterleague.com/",
		Endpoint: "Sponsors",
		Payload:  nil,
		Method:   "GET",
	}

	// []headers are added into the request through DefaultRequest
	resp, err := httplib.DefaultRequest(req, h)

	// unmarshal into sponsors and return
	err = json.Unmarshal(resp, &sponsors)
	if err != nil {
		logrus.Infoln("error unmarshalling response. check the api schema and types")
		return nil, err
	}

	return
}
