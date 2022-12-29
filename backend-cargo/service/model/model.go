package m

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	t, err := json.Marshal(j)
	return t, err
}

func (p *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

type Quote struct {
	ID              uuid.UUID `json:"id"`
	Currency        string    `json:"currency"`
	Partner         string    `json:"partner"`
	BookingId       uuid.UUID `json:"booking_id"`
	Validity        string    `json:"validity"`
	Liner           string    `json:"liner"`
	TransitDays     int       `json:"transit_days"`
	FreeDays        int       `json:"free_days"`
	OriginDate      string    `json:"origin_date"`
	DestinationDate string    `json:"destination_date"`
	ChargesInfo     []byte    `sql:"type:jsonb" json:"charges_info"`
	Remarks         string    `json:"remarks"`
	QuoteStatus     string    `json:"quote_status"`
}
type Chargesinfo struct {
	Name     string  `json:"name"`
	Freight  string  `json:"freight"`
	BuyRate  float32 `json:"buy_rate"`
	BuyTax   float32 `json:"buy_tax"`
	Unit     int     `json:"unit"`
	SellRate float32 `json:"sell_rate"`
	SellTax  float32 `json:"sell_tax"`
}
