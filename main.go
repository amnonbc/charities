package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.charitycommission.gov.uk/register/api"
	apiKey  = "2334eb606a7346c3b5c7aec621843a24"
)

func getTrustees(charityNumber string) (string, []string, error) {
	log := slog.With(
		slog.String("method", "allcharitydetailsV2"),
		slog.String("charityNumber", charityNumber))
	url := fmt.Sprintf("%s/allcharitydetailsV2/%s/0", baseURL, charityNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		log.Error("failed to get charity", "status", resp.Status, "body", body)
		return "", nil, errors.New(resp.Status)
	}

	var c Charity
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		log.Error("error decoding CharityDetails: ", "error", err)
		return "", nil, err
	}

	names := make([]string, len(c.TrusteeNames))
	for i, t := range c.TrusteeNames {
		names[i] = t.TrusteeName
	}
	return c.CharityName, names, nil
}

const exampleCharityNo = "202280"

func loadCharityNumbers(r io.Reader) ([]string, error) {
	cr := csv.NewReader(r)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(records))
	for i, record := range records {
		ids[i] = record[0]
	}
	return ids, nil
}

func main() {
	infileName := flag.String("i", "", "input csv file")
	outfileName := flag.String("o", "out.csv", "output csv file")
	flag.Parse()
	charityNumbers := []string{exampleCharityNo}
	if *infileName != "" {
		f, err := os.Open(*infileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		charityNumbers, err = loadCharityNumbers(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	var cw *csv.Writer
	if *outfileName != "" {
		f, err := os.Create(*outfileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		cw = csv.NewWriter(f)
		cw.Write([]string{"CharityNumber", "Name"})
	}

	for _, id := range charityNumbers {
		name, trustees, err := getTrustees(id)
		if err != nil {
			slog.Error("failed to get trustees", "exampleCharityNo", exampleCharityNo, "error", err)
		}
		rec := []string{id}
		rec = append(rec, name)
		cw.Write(rec)
		for _, t := range trustees {
			rec := []string{"", t}
			cw.Write(rec)
		}
	}
	cw.Flush()
}
