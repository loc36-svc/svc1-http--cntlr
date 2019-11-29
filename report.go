package cntlr

import (
	"fmt"
	"github.com/loc36-svc/svc1-svc1--svc"
	"github.com/qamarian-dtp/err"
	"net/http"
	"strconv"
)
func init () {
	if initReport != nil { return }

	if svc.InitReport != nil {
		initReport = err.New (`Package "github.com/loc36-svc/svc1-svc1--svc" init failed.`,
			nil, nil, svc.InitReport ())
		return
	}
}

func Report (resChan http.ResponseWriter, req *http.Request) {
	if req.FormValue ("state") == "" || req.FormValue ("sensor") == "" || req.FormValue ("sensorPass") == "" || req.FormValue ("serviceID") == "" || req.FormValue ("serviceVer") == "" {

		output := fmt.Sprintf (responseFormat, "Incomplete request data.", "c")
		resChan.Write ([]byte (output))
		return
	}

	if req.FormValue ("serviceID") != serviceID {
		output := fmt.Sprintf (responseFormat, "Service requested from the wrong service.", "d")
		resChan.Write ([]byte (output))
		return
	}

	if req.FormValue ("serviceVer") != serviceVer {
		output := fmt.Sprintf (responseFormat, "Unsupported service version.", "e")
		resChan.Write ([]byte (output))
		return
	}

	state, errX := strconv.Atoi (req.FormValue ("state"))
	if errX != nil {
		output := fmt.Sprintf (responseFormat, "State provided seems invalid.", "f")
		resChan.Write ([]byte (output))
		return
	}

	errY := svc.Service (state, req.FormValue ("sensor"), req.FormValue ("sensorPass"))
	if errY != nil {
		output := fmt.Sprintf (responseFormat, "An error occured.", "b")
		resChan.Write ([]byte (output))
		return
	}

	output := fmt.Sprintf (responseFormat, "State updated successfully!.", "a")
	resChan.Write ([]byte (output))
}
var (
	serviceID = "1"
	serviceVer = "0.1.0"
	responseFormat = `
		{
			response: "%s",
			responseCode: "%s"
		}
	`
)
