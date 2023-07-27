package models

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
	Roles  string `json:"roles"`
}

type User struct {
	UserID      string `json:"userId"`
	Firstname   string `json:"firstname"`
	Midname     string `json:"midname"`
	Lastname    string `json:"lastname"`
	Sex         string `json:"sex"`
	BirthDay    string `json:"birthDay"`
	Citizenship string `json:"citizenship"`
	DocType     string `json:"docType"`
	RegionID    string `json:"regionId"`
	Doc         string `json:"doc"`
	YourSelf    bool   `json:"yourSelf"`
}


type ReqOrder struct {
	Backward      bool `json:"backward"`
	OrdersRequest []OrdersRequest `json:"ordersRequest"`
	WithInsurance       bool   `json:"withInsurance"`
	WithSmsNotification bool   `json:"withSmsNotification"`
	Phone               string `json:"phone"`
	Passengers          []UserInfo2 `json:"passengers"`
}

type UserInfo2 struct {
	Children     interface{} `json:"children"`
	Valid        bool        `json:"valid"`
	Child5From10 bool        `json:"child5From10"`
	BirthDay     string      `json:"birthDay"`
	Citizenship  string      `json:"citizenship"`
	Doc          string      `json:"doc"`
	DocType      string      `json:"docType"`
	Firstname    string      `json:"firstname"`
	Lastname     string      `json:"lastname"`
	Midname      string      `json:"midname"`
	RegionID     string      `json:"regionId"`
	Pinfl        string      `json:"pinfl"`
	Sex          string      `json:"sex"`
}


type Car struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}


type Requirements struct {
	SeatsRange string `json:"seatsRange"`
}

type OrdersRequest struct {
	StationFrom string `json:"stationFrom"`
	StationTo   string `json:"stationTo"`
	DepDate     string `json:"depDate"`
	Train      Train `json:"train"`
	Car  Car `json:"car"`
	Requirements  Requirements `json:"requirements"`
}

type Cart struct {
	PaySysID       string `json:"paySysId"`
	CardCVC        string `json:"cardCVC"`
	CardExpire     string `json:"cardExpire"`
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	RedirectURL    string `json:"redirectUrl"`
}


type OrderPayResponse struct {
	PaySysID       string  `json:"paySysId"`
	OrderID        string  `json:"orderId"`
	TotalAmountSum float64 `json:"totalAmountSum"`
	CurrencyRates  []struct {
		Code      string      `json:"code"`
		BuyPrice  float64     `json:"buyPrice"`
		SellPrice interface{} `json:"sellPrice"`
		Date      interface{} `json:"date"`
	} `json:"currencyRates"`
	Percent float64 `json:"percent"`
}

type Order struct {
	OrderID string `json:"orderId"`
}


type ResponseVagon struct {
	HasError          bool        `json:"hasError"`
	Type              string      `json:"type"`
	ShowWithoutPlaces interface{} `json:"showWithoutPlaces"`
	Direction         []struct {
		HasError  bool   `json:"hasError"`
		Type      string `json:"type"`
		PassRoute struct {
			From     string `json:"from"`
			CodeFrom string `json:"codeFrom"`
			CodeTo   string `json:"codeTo"`
			To       string `json:"to"`
		} `json:"passRoute"`
		Trains []struct {
			Date  string `json:"date"`
			Train struct {
				TrainSchemaType string `json:"trainSchemaType"`
				TrainBrandName  string `json:"trainBrandName"`
				Length          string `json:"length"`
				Type            string `json:"type"`
				Number          string `json:"number"`
				Departure       struct {
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
				ElRegPossible string      `json:"elRegPossible"`
				Parom         interface{} `json:"parom"`
				Bus           interface{} `json:"bus"`
				FirmName      interface{} `json:"firmName"`
				Comments      interface{} `json:"comments"`
				Number2       string      `json:"number2"`
				Route         struct {
					Station []string `json:"station"`
				} `json:"route"`
				Cars []struct {
					Type  string `json:"type"`
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
					Car []struct {
						TypeOfCarSchema struct {
							BarndName    string `json:"barndName"`
							ClassService string `json:"classService"`
							CarNumber    string `json:"carNumber"`
							SeatsCount   string `json:"seatsCount"`
							SchemaName   string `json:"schemaName"`
						} `json:"typeOfCarSchema"`
						Number        string `json:"number"`
						ElRegPossible struct {
							Uk  interface{} `json:"uk"`
							Akp interface{} `json:"akp"`
						} `json:"elRegPossible"`
						Station interface{} `json:"station"`
						SubType string      `json:"subType"`
						Seats   struct {
							SeatsUndef     string      `json:"seatsUndef"`
							SeatsUp        interface{} `json:"seatsUp"`
							SeatsLateralDn interface{} `json:"seatsLateralDn"`
							SeatsLateralUp interface{} `json:"seatsLateralUp"`
							FreeComp       interface{} `json:"freeComp"`
							SeatsDn        interface{} `json:"seatsDn"`
						} `json:"seats"`
						Places     string      `json:"places"`
						Swim       interface{} `json:"swim"`
						NonSmoking interface{} `json:"nonSmoking"`
						SelBedding interface{} `json:"selBedding"`
						PayFood    interface{} `json:"payFood"`
						SelFood    interface{} `json:"selFood"`
						AddFood    interface{} `json:"addFood"`
						DesignCar  interface{} `json:"designCar"`
						SeatsSex   interface{} `json:"seatsSex"`
						TypePlaces interface{} `json:"typePlaces"`
					} `json:"car"`
					ClassService struct {
						Content string `json:"content"`
						Service struct {
							Conditioner bool `json:"conditioner"`
						} `json:"service"`
						Type string `json:"type"`
					} `json:"classService"`
					Carrier struct {
						Name string `json:"name"`
					} `json:"carrier"`
					ComissionFee             string      `json:"comissionFee"`
					Tariff                   string      `json:"tariff"`
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
					TypeShow        string      `json:"typeShow"`
					TrainLetter     string      `json:"trainLetter"`
					Uz              interface{} `json:"uz"`
					ClassServiceInt interface{} `json:"classServiceInt"`
					AddSigns        interface{} `json:"addSigns"`
					Tariff2         interface{} `json:"tariff2"`
				} `json:"cars"`
				Brand     string `json:"brand"`
				TimeInWay string `json:"timeInWay"`
			} `json:"train"`
		} `json:"trains"`
		Content string      `json:"content"`
		Code    interface{} `json:"code"`
	} `json:"direction"`
	ReqExpressZK       string      `json:"reqExpressZK"`
	ReqLocalRecv       string      `json:"reqLocalRecv"`
	ReqAddress         string      `json:"reqAddress"`
	ReqExpressDateTime string      `json:"reqExpressDateTime"`
	ReqLocalSend       string      `json:"reqLocalSend"`
	Content            string      `json:"content"`
	Code               interface{} `json:"code"`
}