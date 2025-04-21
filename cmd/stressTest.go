/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gvillela7/stresstest/internal/handler"
	tableshow "github.com/gvillela7/stresstest/internal/model"
	"github.com/spf13/cobra"
	"github.com/useinsider/go-pkg/insrequester"

	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress test url",
	Long:  `Test to evaluate the number of parallel requests to a given URL`,
	Run: func(cmd *cobra.Command, args []string) {
		requester := insrequester.NewRequester().Load()
		req := requests
		if concurrency == 1 {
			var Report tableshow.Report
			Report.Start = time.Now()
			for i := 0; i < req; i++ {
				response, _ := requester.Get(insrequester.RequestEntity{Endpoint: url})
				switch response.StatusCode {
				case 200:
					Report.Status200 += 1
				case 404:
					Report.Status404 += 1
				case 429:
					Report.Status429 += 1
				case 500:
					Report.Status500 += 1
				}
			}
			timeExec := time.Since(Report.Start)
			tableshow.TableShow(timeExec, req, Report.Status200, Report.Status404, Report.Status429, Report.Status500)
		} else {
			handler.Concurrency(url, requests, concurrency)
		}
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "URL test required")
	err := stressCmd.MarkPersistentFlagRequired("url")
	if err != nil {
		fmt.Println("usage: --url")
		return
	}
	stressCmd.PersistentFlags().IntVarP(&requests, "requests", "r", 1, "Quantity of requests")
	stressCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 1, "Quantity of concurrency")
}
