package horizon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	// "log"
	"regexp"

	"github.com/manucorporat/sse"
	"github.com/nullstyle/coinop/entity"
)

var endEvent = regexp.MustCompile("(\r\n|\r|\n){2}")

type payment struct {
	Type        string `json:"type"`
	PagingToken string `json:"paging_token"`

	Links struct {
		Transaction struct {
			Href string `json:"href"`
		} `json:"transaction"`
	} `json:"_links"`

	// create_account fields
	Account         string `json:"account"`
	Funder          string `json:"funder"`
	StartingBalance string `json:"starting_balance"`

	// payment/path_payment fields
	From        string `json:"from"`
	To          string `json:"to"`
	AssetType   string `json:"asset_type"`
	AssetCode   string `json:"asset_code"`
	AssetIssuer string `json:"asset_issuer"`
	Amount      string `json:"amount"`

	// transaction fields
	Memo struct {
		Type  string `json:"memo_type"`
		Value string `json:"memo"`
	}
}

func (p *payment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *payment) Entity() (ent entity.Payment, err error) {
	if p.Type == "create_account" {
		ent.From = entity.AccountID(p.Funder)
		ent.To = entity.AccountID(p.Account)
		ent.Asset.Type = "native"
		ent.Asset.Code = ""
		ent.Asset.Issuer = ""
		ent.Amount = entity.MustParseAmount(p.StartingBalance)
	} else {
		ent.From = entity.AccountID(p.From)
		ent.To = entity.AccountID(p.To)
		ent.Asset.Type = p.AssetType
		ent.Asset.Code = p.AssetCode
		ent.Asset.Issuer = p.AssetIssuer
		ent.Amount = entity.MustParseAmount(p.Amount)
	}

	ent.PagingToken = p.PagingToken
	ent.Memo.Type = p.Memo.Type
	ent.Memo.Value = p.Memo.Value

	return
}

func loadMemo(p *payment) error {
	res, err := http.Get(p.Links.Transaction.Href)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&p.Memo)
}

func parseEvent(data []byte) (result sse.Event, err error) {
	r := bytes.NewReader(data)
	events, err := sse.Decode(r)
	if err != nil {
		return
	}

	if len(events) != 1 {
		err = fmt.Errorf("only expected 1 event, got: %d", len(events))
		return
	}

	result = events[0]
	return
}

func splitSSE(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, nil
	}

	if loc := endEvent.FindIndex(data); loc != nil {
		return loc[1], data[0:loc[1]], nil
	}

	return 0, nil, nil
}
