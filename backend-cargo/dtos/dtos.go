package dtos

import (
	"github.com/google/uuid"
)

type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

type Account struct {
	Id          uuid.UUID `json:"id"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	AccountType string    `json:"account_type"`
	Email       string    `json:"email"`
	Mobile      string    `json:"mobile"`
	Address     string    `json:"address"`
	Country     string    `json:"country"`
}

type Quote struct {
	ID              uuid.UUID     `json:"id"`
	Currency        string        `json:"currency"`
	Partner         string        `json:"partner"`
	BookingId       uuid.UUID     `json:"booking_id"`
	Validity        string        `json:"validity"`
	Liner           string        `json:"liner"`
	TransitDays     int           `json:"transit_days"`
	FreeDays        int           `json:"free_days"`
	OriginDate      string        `json:"origin_date"`
	DestinationDate string        `json:"destination_date"`
	ChargesInfo     []Chargesinfo `json:"charges_info"`
	Remarks         string        `json:"remarks"`
	QuoteStatus     string        `json:"quote_status"`
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

type Booking struct {
	Id                 uuid.UUID `json:"id"`
	BookedOn           string    `json:"booked_on"`
	CustomerId         uuid.UUID `json:"customer_id"`
	TermsOfShipment    string    `json:"terms_of_shipment"`
	Incoterms          string    `json:"incoterms"`
	OriginPort         string    `json:"origin_port"`
	OriginAddress      string    `json:"origin_address"`
	DestinationPort    string    `json:"destination_port"`
	DestinationAddress string    `json:"destination_address"`
	DoorPickup         bool      `json:"door_pickup"`
	DoorDelivery       bool      `json:"door_delivery"`
	OriginCustoms      bool      `json:"origin_customs"`
	DestinationCustoms bool      `json:"destination_customs"`
	CargoReadyDate     string    `json:"cargo_ready_date"`
	CargoIsDangerous   bool      `json:"cargo_is_dangerous"`
	CargoIsStackable   bool      `json:"cargo_is_stackable"`
	CargoDimensionUnit string    `json:"cargo_dimension_unit"`
	CargoCount         int       `json:"cargo_count"`
	CargoWeight        float32   `json:"cargo_weight"`
	CargoLength        float32   `json:"cargo_length"`
	CargoHeight        float32   `json:"cargo_height"`
	CargoWidth         float32   `json:"cargo_width"`
	CargoHsCode        int       `json:"cargo_hs_code"`
	Remarks            string    `json:"remarks"`
	BookingStatus      string    `json:"booking_status"`
	ConfirmedQuote     uuid.UUID `json:"confirmed_quote"`
}

type Task struct {
	Id        uuid.UUID `json:"id"`
	BookingId uuid.UUID `json:"booking_id"`
	Name      string    `json:"name"`
	SubTask   string    `json:"sub_task"`
	Info      string    `json:"info"`
	Status    string    `json:"status"`
	CreatedAt int64     `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt int64     `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}
