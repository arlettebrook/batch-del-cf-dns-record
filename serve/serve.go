package serve

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/arlettebrook/batch-del-cf-dns-record/models"
)

var (
	conf    = GetConfig()
	baseURL string
)

func setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+conf.ApiToken)
	req.Header.Set("Content-Type", "application/json")
}

func closeRespBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		Logger.Fatal(err)
	}
}

func getRecords() (models.Result, error) {
	Logger.Info("Start fetch DNS records...")
	var result models.Result
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return result, fmt.Errorf("creating request: %w", err)
	}
	setHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("making request: %w", err)
	}
	defer closeRespBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("fetch DNS records fail: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("reding response body: %w", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("unmarshal response: %w", err)
	}
	Logger.WithFields(logrus.Fields{
		"TotalDNSRecords": result.ResultInfo.TotalCount,
	}).Info("Fetch DNS records success!")
	return result, nil
}

func verifyParams() {
	if conf.ApiToken == "" || conf.ZoneID == "" {
		Logger.Fatalf("Both api_token" +
			" and zone_id command line arguments must be provided")
	}
	baseURL = fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records",
		conf.ZoneID)
}

func deleteRecords(DNSRecord models.DNSRecord) error {
	client := &http.Client{}
	deleteURL := fmt.Sprintf(fmt.Sprintf("%s/%s", baseURL, DNSRecord.ID))
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return fmt.Errorf("creating delete request: %w", err)
	}

	setHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("making delete request: %w", err)
	}
	defer closeRespBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete request failed: %s", resp.Status)
	}
	Logger.Warnf("Successfully deleted %s that pointed to %s\n",
		DNSRecord.Name, DNSRecord.Content)

	return nil
}

func Start() {
	Logger.Info("Mission begins...")
	verifyParams()

	for {
		records, err := getRecords()
		if err != nil {
			Logger.Fatalf("Failed to fetch DNS records: %s", err)
		}

		if records.TotalCount != 0 {
			Logger.Info("Start deleting dns records!")
			var wg sync.WaitGroup
			for _, DNSRecord := range records.Result {
				wg.Add(1)
				go func(DNSRecord models.DNSRecord) {
					defer wg.Done()
					if err := deleteRecords(DNSRecord); err != nil {
						Logger.Errorf(
							"Failed to delete record %s: %v",
							DNSRecord.Name, err)
					}
				}(DNSRecord)
			}
			wg.Wait()
		} else {
			Logger.Info("All DNS records delete success!")
			break
		}
	}

	Logger.Info("Mission ended!")
}
