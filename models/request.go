package models

type Stop struct{
	StopFor string  `json:"stop"`
}

type RequestBody struct {
	Direction []Direction `json:"direction"`
	StationFrom       string `json:"stationFrom"`
	StationTo         string `json:"stationTo"`
	DetailNumPlaces   int    `json:"detailNumPlaces"`
	ShowWithoutPlaces int    `json:"showWithoutPlaces"`
	Wagon WagonType  `json:"wagon"`
	Time string   `json:"time"`
	Types string `json:"types"`
}

type RequestSearch struct {
	Direction []Direction `json:"direction"`
	StationFrom       string `json:"stationFrom"`
	StationTo         string `json:"stationTo"`
	DetailNumPlaces   int    `json:"detailNumPlaces"`
	ShowWithoutPlaces int    `json:"showWithoutPlaces"`
}

type Direction struct{
	DepDate string `json:"depDate"`
	Fullday bool   `json:"fullday"`
	Type    string `json:"type"`
}

type WagonType struct{
	Name string 
}

type Time struct{
	Time string
}

type Response struct{
	Resp string `json:"resp"`
}

type ResponseAll struct {
	Express struct {
		HasError     bool   `json:"hasError"`
		Type         string `json:"type"`
		ReqExpressZK string `json:"reqExpressZK"`
		Direction    []struct {
			Type      string `json:"type"`
			PassRoute struct {
				From     string `json:"from"`
				CodeFrom string `json:"codeFrom"`
				CodeTo   string `json:"codeTo"`
				To       string `json:"to"`
			} `json:"passRoute"`
			Trains []struct {
				Date  string `json:"date"`
				Train []struct {
					Length string `json:"length"`
					Type   string `json:"type"`
					Number string `json:"number"`
					Brand  string `json:"brand"`
					Route  struct {
						Station []string `json:"station"`
					} `json:"route"`
					Places struct {
						Cars []struct {
							Type  string `json:"type"`
							Seats struct {
								SeatsUndef     string      `json:"seatsUndef"`
								SeatsUp        interface{} `json:"seatsUp"`
								SeatsLateralDn interface{} `json:"seatsLateralDn"`
								SeatsLateralUp interface{} `json:"seatsLateralUp"`
								FreeComp       interface{} `json:"freeComp"`
								SeatsDn        interface{} `json:"seatsDn"`
							} `json:"seats"`
							Tariffs struct {
								Tariff []struct {
									Owner struct {
										Type    string `json:"type"`
										Country struct {
											Name string `json:"name"`
											Code string `json:"code"`
										} `json:"country"`
										Railway struct {
											Name string `json:"name"`
											Code string `json:"code"`
										} `json:"railway"`
									} `json:"owner"`
									Station      interface{} `json:"station"`
									ClassService struct {
										Content string `json:"content"`
										Service struct {
										} `json:"service"`
										Type string `json:"type"`
									} `json:"classService"`
									Seats struct {
										SeatsUndef     string      `json:"seatsUndef"`
										SeatsUp        interface{} `json:"seatsUp"`
										SeatsLateralDn interface{} `json:"seatsLateralDn"`
										SeatsLateralUp interface{} `json:"seatsLateralUp"`
										FreeComp       interface{} `json:"freeComp"`
										SeatsDn        interface{} `json:"seatsDn"`
									} `json:"seats"`
									Carrier struct {
										Name string `json:"name"`
									} `json:"carrier"`
									ComissionFee    string      `json:"comissionFee"`
									Tariff          string      `json:"tariff"`
									ClassServiceInt interface{} `json:"classServiceInt"`
									AddSigns        interface{} `json:"addSigns"`
									ElRegPossible   struct {
										Uk  interface{} `json:"uk"`
										Akp interface{} `json:"akp"`
									} `json:"elRegPossible"`
									Swim                     interface{} `json:"swim"`
									SelBedding               interface{} `json:"selBedding"`
									PayFood                  interface{} `json:"payFood"`
									SelFood                  interface{} `json:"selFood"`
									AddFood                  interface{} `json:"addFood"`
									DesignCar                interface{} `json:"designCar"`
									SeatsSex                 interface{} `json:"seatsSex"`
									Tariff2                  interface{} `json:"tariff2"`
									TariffService            interface{} `json:"tariffService"`
									Ud                       interface{} `json:"ud"`
									SellingInternetForbidden interface{} `json:"sellingInternetForbidden"`
									SaleOnTwo                interface{} `json:"saleOnTwo"`
									SaleOnFour               interface{} `json:"saleOnFour"`
									Modificators             interface{} `json:"modificators"`
									PassVok                  struct {
										StationFrom struct {
											Code string `json:"code"`
										} `json:"stationFrom"`
										StationTo interface{} `json:"stationTo"`
									} `json:"passVok"`
								} `json:"tariff"`
							} `json:"tariffs"`
							IndexType string `json:"indexType"`
							FreeSeats string `json:"freeSeats"`
							TypeShow  string `json:"typeShow"`
						} `json:"cars"`
					} `json:"places"`
					ElRegPossible string `json:"elRegPossible"`
					TimeInWay     string `json:"timeInWay"`
					Number2       string `json:"number2"`
					Departure     struct {
						Time      string      `json:"time"`
						LocalTime string      `json:"localTime"`
						Stop      interface{} `json:"stop"`
						LocalDate string      `json:"localDate"`
					} `json:"departure"`
					Arrival struct {
						Time      string `json:"time"`
						LocalTime string `json:"localTime"`
						Date      string `json:"date"`
						Stop      string `json:"stop"`
						LocalDate string `json:"localDate"`
					} `json:"arrival"`
					DepartureTrain struct {
						Date string `json:"date"`
					} `json:"departureTrain"`
					Parom    interface{} `json:"parom"`
					Bus      interface{} `json:"bus"`
					FirmName interface{} `json:"firmName"`
					Comments interface{} `json:"comments"`
				} `json:"train"`
			} `json:"trains"`
			NotAllTrains interface{} `json:"notAllTrains"`
		} `json:"direction"`
		ShowWithoutPlaces  interface{} `json:"showWithoutPlaces"`
		ReqExpressDateTime string      `json:"reqExpressDateTime"`
		ReqLocalSend       string      `json:"reqLocalSend"`
		ReqLocalRecv       string      `json:"reqLocalRecv"`
		ReqAddress         string      `json:"reqAddress"`
		Content            string      `json:"content"`
		Code               interface{} `json:"code"`
	} `json:"express"`
	Discount []interface{} `json:"discount"`
}


type TrainAddresses struct {
	StationFrom     string `json:"stationFrom"`
	StationTo       string `json:"stationTo"`
	DetailNumPlaces int    `json:"detailNumPlaces"`
	Direction []Direction2 `json:"direction"`
}


type Direction2 struct{
	DepDate string `json:"depDate"`
	Fullday bool   `json:"fullday"`
	Type    string `json:"type"`
	Train  Train   `json:"train"`
}

type Train struct {
	Number string `json:"number"`
}


type Test13 struct {
	ID                 string `json:"id"`
	PaymentID          string `json:"paymentId"`
	Type               string `json:"type"`
	WithLoyaltyProgram bool   `json:"withLoyaltyProgram"`
}