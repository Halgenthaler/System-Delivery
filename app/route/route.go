package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string      `json:"routeId"`
	ClientID  string      `json:"clientId"`
	Positions []Positions `json:"positions"`
}

type Positions struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialPositions struct {
	ID        string    `json:"routeId"`
	ClientID  string    `json:"clientId"`
	Positions []float64 `json:"positions"`
	Finished  bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route id not informed")
	}
	//TO DO revisar o processo de abertura de documentos
	f, err := os.Open("destination/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Positions{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

//Gera a nossa lista de posição em json/sting
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialPositions
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Positions = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))

	}
	return result, nil
}
