package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Station struct {
	StationId   int    `json:"StationId"`
	StationName string `json:"StationName"`
}

type stationData struct {
	Data []Station `json:"d"`
}

type Schedule struct {
	ScheduleName   string  `json:"ScheduleName"`
	BusNo          string  `json:"BusNo"`
	BusDescription string  `json:"BusDescription"`
	BusType        string  `json:"BusType"`
	DDate          string  `json:"DDate"`
	RouteId        int     `json:"RouteId"`
	RouteName      string  `json:"RouteName"`
	Time           string  `json:"Time"`
	ScheduleId     int     `json:"SchedTimeuleId"`
	Active         string  `json:"Active"`
	TripType       string  `json:"TripType"`
	SeatFare       float32 `json:"SeatFare"`
	NumberOfSeat   int     `json:"NumberOfSeat"`
	SeatUpdate     int     `json:"SeatUpdate"`
	SeatId         int     `json:"SeatId"`
	Seats          string  `json:"Seats"`
	ReservedSeats  string  `json:"ReservedSeats"`
}

type scheduleData struct {
	Data []Schedule `json:"d"`
}

func GetStationsByCounter(counterId string) ([]Station, error) {
	url := "https://sbsuperdeluxe.com/SellTicketOnline.aspx/GetStationsByCounter"
	method := "POST"

	payload := strings.NewReader(`{
		"counterId":   ` + counterId + `
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res stationData
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getCookie(url string) (*string, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	cookieReady := "ASP.NET_SessionId="
	cookies := response.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "ASP.NET_SessionId" {
			cookieReady += cookie.Value
		}
	}

	return &cookieReady, nil
}

func GetSchedule(counterId, stationId, journeyDate, userId string) ([]Schedule, error) {
	cookie, err := getCookie("https://sbsuperdeluxe.com/SellTicketOnline.aspx")
	if err != nil {
		return nil, err
	}

	url := "https://sbsuperdeluxe.com/SellTicketOnline.aspx/ShowScheduleByCounterAndStationOldUser"
	method := "POST"

	payload := strings.NewReader(`{
		"counterId":   ` + counterId + `,
		"stationId":   ` + stationId + `,
		"journeyDate": "` + journeyDate + `",
		"userId":      "` + userId + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", *cookie)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res scheduleData
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	var schedules []Schedule
	for _, val := range res.Data {
		if val.RouteName != "" {
			schedules = append(schedules, val)
		}
	}

	return schedules, nil
}
