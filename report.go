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

	errX := svc.InitReport ()
	if errX != nil {
		initReport = err.New (`Package "github.com/loc36-svc/svc1-svc1--svc" " +
			" init failed.`, nil, nil, svc.InitReport ())
		return
	}
}

// Function Report () records the reported state of a sensor. It expects its request to be
// a POST request. The request must provide the following field data: state, sensor,
// sensorPass, serviceId, and serviceVer.
//
//	state:      Should be the state of the sensor.
//	sensor:     Should be the id of the sensor.
//	sensorPass: Should be the pass of the sensor.
//	serviceId:  Should be the id of the service the client thinks its interacting with.
//	serviceVer: Should be the service version the client needs.
//
// All the field data above are mandatory.
//
// Response code would ordinarily be: 200. Any other error code should be treated as fatal
// error.
// When response code is 200, a JSON data would be output:
//
// 	{
//		response: "{x}",
//		responseCode: "{y}"
//	}
//
// where '{x}' is the response of the request; and '{y}' is a code for the response.
//
// Possible codes are:
// 	a: State updated successfully! 
//	b: An error occured.
//	c: Incomplete request data.
//	d: Service requested from the wrong service.
//	e: Unsupported service version.
//	f: State provided seems invalid.
//
func Report (resChan http.ResponseWriter, req *http.Request) {
	if req.FormValue ("state") == "" || req.FormValue ("sensor") == "" ||
	req.FormValue ("sensorPass") == "" || req.FormValue ("serviceId") == "" ||
	req.FormValue ("serviceVer") == "" {

		output := fmt.Sprintf (responseFormat, "Incomplete request data.", "c")
		resChan.Write ([]byte (output))
		return
	}

	if req.FormValue ("serviceId") != serviceId {
		output := fmt.Sprintf (responseFormat, "Service requested from the " +
			"wrong service.", "d")
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
		output := fmt.Sprintf (responseFormat, "State provided seems invalid.",
			"f")
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
	serviceId = "1"
	serviceVer = "0.1.0"
	responseFormat = `
		{
			response: "%s",
			responseCode: "%s"
		}
	`
)
