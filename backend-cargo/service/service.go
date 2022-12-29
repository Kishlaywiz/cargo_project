package service

import (
	"backend/dtos"
	"backend/globals"
	g "backend/globals"
	m "backend/service/model"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *gorm.DB

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

func CreateAccount(req *dtos.Account) (*uuid.UUID, error) {
	id := uuid.New()
	account := dtos.Account{
		Id:          id,
		UserName:    req.UserName,
		Password:    req.Password,
		AccountType: req.AccountType,
		Email:       req.Email,
		Mobile:      req.Mobile,
		Address:     req.Address,
		Country:     req.Country,
	}
	err := dbConnect.Table("config").Save(&account).Error
	if err != nil {
		log.Println("Database Error", err)
		return nil, err
	}
	return &id, nil
}

func Login(req *dtos.Login) (*dtos.Account, error) {
	var account dtos.Account
	err := dbConnect.Raw("SELECT * FROM config WHERE email=?", req.Email).Scan(&account).Error
	if err != nil {
		log.Println("Database error", err)
		return nil, err
	}
	fmt.Println(req.Password)
	fmt.Println(account.Password)
	if req.Password != account.Password {
		err := errors.New("Unauthorised Access")
		log.Println("User Access error", err)
		return nil, err
	}
	res := dtos.Account{
		Id:          account.Id,
		UserName:    account.UserName,
		AccountType: account.AccountType,
		Email:       account.Email,
		Mobile:      account.Mobile,
		Address:     account.Address,
	}
	return &res, nil
}

func GetCustomers() ([]dtos.Account, error) {
	var accounts []dtos.Account
	err := dbConnect.Raw("select * from config where account_type='Customer'").Scan(&accounts).Error
	if err != nil {
		log.Println("Databse Error", err)
		return nil, err
	}
	resp := []dtos.Account{}
	for _, account := range accounts {
		res := dtos.Account{
			Id:          account.Id,
			UserName:    account.UserName,
			AccountType: account.AccountType,
			Email:       account.Email,
			Mobile:      account.Mobile,
			Country:     account.Country,
		}
		resp = append(resp, res)
	}
	return resp, nil
}

func GetBookingTasks() map[string][]string {
	bookingTasks := make(map[string][]string)
	bookingTasks[g.TaskBookingCreated] = []string{}
	bookingTasks[g.TaskBookingConfirmed] = []string{}
	bookingTasks[g.TaskCargoPickup] = []string{"Pickup Acknowledgement"}
	bookingTasks[g.TaskCargoArrived] = []string{"Is Gated In"}
	bookingTasks[g.TaskVesselDeparted] = []string{"Voyage No.", "Vessel No."}
	bookingTasks[g.TaskVesselArrived] = []string{}
	bookingTasks[g.TaskCargoDeparted] = []string{}
	bookingTasks[g.TaskCargoDelivered] = []string{}
	bookingTasks[g.TaskBookingCompleted] = []string{}
	return bookingTasks
}

func AssignBookingTasks(bookingId uuid.UUID, by uuid.UUID) error {
	tasks := GetBookingTasks()
	time := time.Now().UTC().Unix()

	var bookingTask dtos.Task
	for task, subTasks := range tasks {
		if len(subTasks) > 0 {
			for _, subTask := range subTasks {
				bookingTask = dtos.Task{
					Id:        uuid.New(),
					BookingId: bookingId,
					Name:      task,
					SubTask:   subTask,
					Status:    g.StatusPending,
					CreatedAt: time,
					CreatedBy: by,
					UpdatedAt: time,
					UpdatedBy: by,
				}
				fmt.Println("Upserting Task : ", task)
				err := dbConnect.Table("task").Save(&bookingTask).Error
				if err != nil {
					log.Println("databse function error", err)
					return err
				}
			}
		} else {
			bookingTask = dtos.Task{
				Id:        uuid.New(),
				BookingId: bookingId,
				Name:      task,
				SubTask:   task,
				Status:    g.StatusPending,
				CreatedAt: time,
				CreatedBy: by,
				UpdatedAt: time,
				UpdatedBy: by,
			}
			fmt.Println("Upserting Task : ", task)
			err := dbConnect.Table("task").Save(&bookingTask).Error
			if err != nil {
				log.Println("databse function error", err)
				return err
			}
		}
	}

	return nil
}

func CreateBookingRequest(req *dtos.Booking) (*uuid.UUID, error) {
	booking := dtos.Booking{
		Id:                 uuid.New(),
		BookedOn:           req.BookedOn,
		CustomerId:         req.CustomerId,
		TermsOfShipment:    req.TermsOfShipment,
		Incoterms:          req.Incoterms,
		OriginPort:         req.OriginPort,
		OriginAddress:      req.OriginAddress,
		DestinationPort:    req.DestinationPort,
		DestinationAddress: req.DestinationAddress,
		DoorPickup:         req.DoorPickup,
		DoorDelivery:       req.DoorDelivery,
		OriginCustoms:      req.OriginCustoms,
		DestinationCustoms: req.DestinationCustoms,
		CargoReadyDate:     req.CargoReadyDate,
		CargoIsDangerous:   req.CargoIsDangerous,
		CargoIsStackable:   req.CargoIsStackable,
		CargoDimensionUnit: req.CargoDimensionUnit,
		CargoCount:         req.CargoCount,
		CargoWeight:        req.CargoWeight,
		CargoLength:        req.CargoLength,
		CargoHeight:        req.CargoHeight,
		CargoWidth:         req.CargoWidth,
		CargoHsCode:        req.CargoHsCode,
		Remarks:            req.Remarks,
		BookingStatus:      req.BookingStatus,
	}
	err := dbConnect.Table("booking").Create(&booking)
	if err.Error != nil {
		log.Println("databse function error", err.Error)
		return nil, err.Error
	}

	return &booking.Id, nil
}

func UpdateBookingRequest(req *dtos.Booking, by uuid.UUID) error {
	var booking dtos.Booking
	err := dbConnect.Raw("select * from booking where id=?", req.Id).First(&booking).Error
	if err != nil {
		log.Println("Database Error", err)
		return err
	}

	if booking.BookingStatus == globals.StatusBookingConfirmed {
		err = errors.New("Booking Already Confirmed\nCannot Update Booking")
		log.Println("Database Error", err)
		return err
	}

	bookingUpdates := dtos.Booking{
		Id:                 req.Id,
		BookedOn:           req.BookedOn,
		CustomerId:         req.CustomerId,
		TermsOfShipment:    req.TermsOfShipment,
		Incoterms:          req.Incoterms,
		OriginPort:         req.OriginPort,
		OriginAddress:      req.OriginAddress,
		DestinationPort:    req.DestinationPort,
		DestinationAddress: req.DestinationAddress,
		DoorPickup:         req.DoorPickup,
		DoorDelivery:       req.DoorDelivery,
		OriginCustoms:      req.OriginCustoms,
		DestinationCustoms: req.DestinationCustoms,
		CargoReadyDate:     req.CargoReadyDate,
		CargoIsDangerous:   req.CargoIsDangerous,
		CargoIsStackable:   req.CargoIsStackable,
		CargoDimensionUnit: req.CargoDimensionUnit,
		CargoCount:         req.CargoCount,
		CargoWeight:        req.CargoWeight,
		CargoLength:        req.CargoLength,
		CargoHeight:        req.CargoHeight,
		CargoWidth:         req.CargoWidth,
		CargoHsCode:        req.CargoHsCode,
		Remarks:            req.Remarks,
		BookingStatus:      req.BookingStatus,
	}

	err = dbConnect.Table("booking").Save(&bookingUpdates).Error
	if err != nil {
		log.Println("Database error", err)
		return err
	}

	if bookingUpdates.BookingStatus == g.StatusBookingConfirmed {
		errs := AssignBookingTasks(booking.Id, by)
		if errs != nil {
			log.Println("Database error", errs)
			return errs
		}
	}
	return nil
}

func AllBookingRequest() ([]dtos.Booking, error) {
	var res []dtos.Booking
	err := dbConnect.Raw("select * from booking").Scan(&res)
	if err.Error != nil {
		log.Println("dberror", err.Error)
		return nil, err.Error
	}
	return res, nil
}

func GetBookingRequest(id uuid.UUID) (dtos.Booking, error) {
	var res dtos.Booking
	err := dbConnect.Raw("select * from booking where id=?", id).Scan(&res)
	if err.Error != nil {
		log.Println("dberror", err.Error)
	}
	return res, nil

}

func GetBookingQuote(id uuid.UUID) (dtos.Quote, error) {
	var res m.Quote
	err := dbConnect.Raw("select * from quotes where id=?", id).Scan(&res)
	if err.Error != nil {
		log.Println("dberror", err.Error)
	}
	var charginfo []dtos.Chargesinfo
	json.Unmarshal(res.ChargesInfo, &charginfo)
	fmt.Println(charginfo, "res\n->>>", res)
	quoteres := dtos.Quote{
		ID:              res.ID,
		Currency:        res.Currency,
		Partner:         res.Partner,
		BookingId:       res.BookingId,
		Validity:        res.Validity,
		Liner:           res.Liner,
		TransitDays:     res.TransitDays,
		FreeDays:        res.FreeDays,
		OriginDate:      res.OriginDate,
		DestinationDate: res.DestinationDate,
		ChargesInfo:     charginfo,
		Remarks:         res.Remarks,
		QuoteStatus:     res.QuoteStatus,
	}
	return quoteres, nil
}

func GetBookingAllQuotes(bid uuid.UUID) ([]dtos.Quote, error) {
	var res []m.Quote
	var allquotes []dtos.Quote
	err := dbConnect.Raw("select * from quotes where booking_id=?", bid).Scan(&res)
	if err.Error != nil {
		log.Println("dberror", err.Error)
		return nil, err.Error
	}
	for _, val := range res {
		var charginfo []dtos.Chargesinfo
		json.Unmarshal(val.ChargesInfo, &charginfo)

		quoteres := dtos.Quote{
			ID:              val.ID,
			Currency:        val.Currency,
			Partner:         val.Partner,
			BookingId:       val.BookingId,
			Validity:        val.Validity,
			Liner:           val.Liner,
			TransitDays:     val.TransitDays,
			FreeDays:        val.FreeDays,
			OriginDate:      val.OriginDate,
			DestinationDate: val.DestinationDate,
			ChargesInfo:     charginfo,
			Remarks:         val.Remarks,
			QuoteStatus:     val.QuoteStatus,
		}
		allquotes = append(allquotes, quoteres)
	}

	return allquotes, nil
}

func CreateQuote(req *dtos.Quote) (*uuid.UUID, error) {
	chargereq, _ := json.Marshal(req.ChargesInfo)
	id:=uuid.New()
	quote := m.Quote{
		ID:              id,
		Currency:        req.Currency,
		Partner:         req.Partner,
		BookingId:       req.BookingId,
		Validity:        req.Validity,
		Liner:           req.Liner,
		TransitDays:     req.TransitDays,
		FreeDays:        req.FreeDays,
		OriginDate:      req.OriginDate,
		DestinationDate: req.DestinationDate,
		ChargesInfo:     chargereq,
		Remarks:         req.Remarks,
		QuoteStatus:     req.QuoteStatus,
	}
	er := dbConnect.Table("quotes").Create(&quote)
	if er.Error != nil {
		log.Println("databse function error", er.Error)
		return nil, er.Error
	}
	fmt.Println(quote.QuoteStatus,"req",req.QuoteStatus)
	if req.QuoteStatus=="Approved"{
		fmt.Println("inside if ************",quote.ID,quote.BookingId)
		er := dbConnect.Table("booking").Where("id = ?", quote.BookingId).Update("confirmed_quote",quote.ID)
		if er.Error != nil {
		log.Println("databse function error", er.Error)
		return nil, er.Error
		}

	}

	return &quote.ID, nil
}

func UpdateQuote(req *dtos.Quote) error {
	chargereq, _ := json.Marshal(req.ChargesInfo)
	quote := m.Quote{
		ID:              req.ID,
		Currency:        req.Currency,
		Partner:         req.Partner,
		BookingId:       req.BookingId,
		Validity:        req.Validity,
		Liner:           req.Liner,
		TransitDays:     req.TransitDays,
		FreeDays:        req.FreeDays,
		OriginDate:      req.OriginDate,
		DestinationDate: req.DestinationDate,
		ChargesInfo:     chargereq,
		Remarks:         req.Remarks,
		QuoteStatus:     req.QuoteStatus,
	}
	err := dbConnect.Table("quotes").Save(quote).Error
	if err != nil {
		log.Println("Database Error", err)
		return err
	}
	if quote.QuoteStatus=="Approved"{
		er := dbConnect.Table("booking").Raw("update booking set confirmed_quote=?",req.ID)
		if er.Error != nil {
		log.Println("databse function error", er.Error)
		return  er.Error
		}
	}
	return nil
}

func GetBookingTask(id uuid.UUID) (dtos.Task, error) {
	var res dtos.Task
	err := dbConnect.Raw("select * from task where id=?", id).Scan(&res)
	if err.Error != nil {
		log.Println("dberror", err.Error)
	}
	return res, nil
}

func GetBookingAllTask(id uuid.UUID) ([]dtos.Task, error) {
	var res []dtos.Task
	err := dbConnect.Raw("select * from task where booking_id=?", id).Scan(&res).Order("created_at ASC")
	if err.Error != nil {
		log.Println("dberror", err.Error)
		return nil, err.Error
	}
	return res, nil
}

func UpdateBookingTask(req *dtos.Task, by uuid.UUID) error {
	req.UpdatedAt = time.Now().Unix()
	req.UpdatedBy = by
	err := dbConnect.Table("task").Save(req).Error
	if err != nil {
		log.Println("databse function error", err)
		return err
	}
	return nil
}









