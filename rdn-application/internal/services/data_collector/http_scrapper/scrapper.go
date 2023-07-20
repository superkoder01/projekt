package http_scrapper

import (
	"RDN-application/internal/model"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "02-01-2006"

type collector struct {
	scrapper *colly.Collector
	config   config.AppConfig
	logger   logger.Logger
	headers  map[string][]string
}

func NewDataCollector(config config.AppConfig, logger logger.Logger) *collector {
	collector := &collector{
		scrapper: colly.NewCollector(),
		config:   config,
		logger:   logger,
		headers:  convertHeaderTitlesToCamelCase(config.GetCollectorConfig().Headers, config.GetCollectorConfig().HeaderNames),
	}
	collector.scrapper.AllowURLRevisit = true
	collector.withOnRequestLogs()
	return collector
}

func (collector *collector) withOnRequestLogs() {
	collector.scrapper.OnRequest(func(req *colly.Request) {
		collector.logger.Infof("OnRequest: ---------------------")
		collector.logger.Infof("Url: %v", req.URL)
		collector.logger.Infof("Method: %v", req.Method)
		collector.logger.Infof("Headers(%v): ", len(*req.Headers))
		for key, value := range *req.Headers {
			collector.logger.Infof("%v: %v", key, value)
		}
	})

	collector.scrapper.OnResponse(func(resp *colly.Response) {
		collector.logger.Infof("OnResponse: ---------------------")
		collector.logger.Infof("Status Code: %v", resp.StatusCode)
		collector.logger.Infof("Headers: ")
		for key, value := range *resp.Headers {
			collector.logger.Infof("%s: %s", key, value)
		}
		//collector.logger.Debugf("Body: %v", string(resp.Body))
	})

	collector.scrapper.OnError(func(r *colly.Response, err error) {
		collector.logger.Infof("Error:", r.StatusCode, err)
	})
}

func (collector *collector) CollectDataFromDate(ctx context.Context, date time.Time, chanel chan model.Collection) {
	dayCollection := &model.DayCollection{
		HourData: map[string]float64{},
	}

	scrapperCfg := collector.config.GetCollectorConfig().Scrapper

	collector.scrapper.OnHTML(scrapperCfg.MainSelector, func(container *colly.HTMLElement) {
		collector.logger.Debugf("%v: ---------------------", scrapperCfg.MainSelector)

		container.ForEach(scrapperCfg.ConditionSelector.SelectorName, func(_ int, header *colly.HTMLElement) {
			collector.logger.Debugf("%v, %v", scrapperCfg.ConditionSelector.SelectorName, header.Text)

			if strings.Contains(header.Text, scrapperCfg.ConditionSelector.ConditionValue) {
				collector.logger.Infof("Fetching: %v", header.Text)
				collectedDate := strings.Fields(header.Text)
				collector.logger.Debugf("Data date: %v", collectedDate[len(collectedDate)-1])

				parse, err := time.Parse(DateFormat, collectedDate[len(collectedDate)-1])
				if err != nil {
					parse = date.AddDate(0, 0, 1)
				}
				dayCollection.AppendDate(parse)

				if strings.Contains(strings.TrimSpace(container.Text), "BRAK DANYCH") {
					collector.logger.Errorf("No data in Data Selectors !! Cannot fetch data")
					return
				}

				container.ForEach(scrapperCfg.DataSelector.SelectorName, func(_ int, table *colly.HTMLElement) {
					var hour string
					var fixing float64
					var err error = nil

					for _, dataElement := range scrapperCfg.DataSelector.Data {
						switch dataElement.DataType {
						case config.Hour:
							data := collector.fetchSingleRow(table, dataElement.SelectorName)
							hour, err = collector.convertStringToHour(data)
						case config.Fixing:
							data := collector.fetchSingleRow(table, dataElement.SelectorName)
							fixing, err = collector.convertStringToFixing(data)
						}
					}
					if err == nil {
						dayCollection.AppendData(hour, fixing)
						//collector.logger.Debugf("Collected data: %v", fmt.Sprintf("%v", dayCollection.HourData))
					} else {
						collector.logger.Errorf("Data not saved hour: %v, fixing: %v error occurred: %v", hour, fixing, err)
					}
				})
			}
		})
	})

	select {
	case <-ctx.Done():
		chanel <- dayCollection
	default:
		err := collector.scrapper.Request(collector.config.GetCollectorConfig().RequestMethod,
			collector.config.GetCollectorConfig().DestinationUrl+"?dateShow="+date.Format(DateFormat)+"&dateAction=next",
			strings.NewReader(``),
			colly.NewContext(),
			collector.headers,
		)
		if err != nil {
			collector.logger.Errorf("%v", err)
		}
	}

	collector.logger.Infof(fmt.Sprintf("Collected: %v", dayCollection.HourData))

	chanel <- dayCollection
}

func (collector *collector) convertStringToHour(value string) (string, error) {
	if len(value) > 1 {
		return value, nil
	} else {
		return "", errors.New("invalid string value")
	}
}

func (collector *collector) convertStringToFixing(value string) (float64, error) {
	if len(value) > 1 {
		fixing, err := strconv.ParseFloat(strings.Replace(value, ",", ".", 1), 64)
		if err != nil {
			collector.logger.Errorf("Failed to convert %v to FLOAT %v", value, err)
			return 0, err
		}
		return fixing, nil
	} else {
		return 0, errors.New("invalid string value")
	}
}

func (collector *collector) fetchSingleRow(table *colly.HTMLElement, selector string) string {
	attr := table.ChildText(selector)
	//collector.logger.Debugf("Data Attribute: %v", attr)
	return strings.TrimSpace(attr)
}

func convertHeaderTitlesToCamelCase(headerMap map[string][]string, headerNames []string) map[string][]string {
	var newKeySlice = map[string]string{}
	for key, _ := range headerMap {
		for _, name := range headerNames {
			if strings.ToLower(key) == strings.ToLower(name) {
				newKeySlice[key] = name
			}
		}
	}

	for key, newKey := range newKeySlice {
		headerMap[newKey] = headerMap[key]
		delete(headerMap, key)
	}
	return headerMap
}
