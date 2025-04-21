package model

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Report struct {
	Requests  int
	Start     time.Time
	Status200 int
	Status404 int
	Status429 int
	Status500 int
}

var (
	colTitleTime    = "Tempo de Execução"
	colTitleRequest = "Total de Requests"
	colTitle200     = "Status 200"
	colTitle404     = "Status 404"
	colTitle429     = "Status 429"
	colTitle500     = "Status 500"
	rowHeader       = table.Row{colTitleTime, colTitleRequest, colTitle200, colTitle404, colTitle429, colTitle500}
)

func TableShow(timeExec time.Duration, requests, status200, status404, status429, status500 int) {
	tw := table.NewWriter()
	row := table.Row{timeExec, requests, status200, status404, status429, status500}
	tw.AppendHeader(rowHeader)
	tw.AppendRows([]table.Row{row})

	tw.SetStyle(table.StyleLight)
	tw.Style().Title.Align = text.AlignCenter
	fmt.Println(tw.Render())
}
