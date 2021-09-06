package functions

import (
	"encoding/json"
	"facebook-connector/pkg/model"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func CreatingPost(message string, token string) (*model.Responsestring, error) {
	client := resty.New()
	resp , err:= client.R().
		SetHeaders(map[string]string{
			"Accept": "*/*" ,
			"Content-Type": "application/json",
			"Accept-Encoding": "gzip, deflate, br",
			"Connection":"keep-alive",
		}).
		Post("https://graph.facebook.com/v11.0/102279725521691/feed?message="+message+"&access_token=" + token)

	var jsonModel *model.Responsestring
	error_ := json.Unmarshal([]byte(resp.String()), jsonModel)

	if error_ != nil{
		fmt.Println(error_)
	}

	return jsonModel, err
}
