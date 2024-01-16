package main

import (
	"fmt"
	"os"

	// "github.com/julienschmidt/httprouter"
	// "github.com/skip2/go-qrcode"
	// "github.com/yeqown/go-qrcode"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

var websiteURL = "https://github.com/srijith13" // GoDevTips URL
var imageSize = 1024                            // 256 x 256 pixels  = 256

func main() {
	// fmt.Println(os.Args, os.Args[1])
	if len(os.Args) == 2 && os.Args[1] == "file" { // Determine which is option to launch ///////// go run main.go file

		////  Generate only the raw PNG bytes for skip2
		// outputFileName := "qr_code2.png"
		// fmt.Printf("outputFileName ", outputFileName)
		// taskError := qrcode.WriteFile(websiteURL, qrcode.High, imageSize, outputFileName)

		// if taskError != nil {
		// 	log.Fatal("Error generating QR code to file. ", taskError)
		// } else {
		// 	fmt.Println("Qrcode file created!")
		// }

		////  Generate only the raw PNG bytes for yeqown

		qrc, err := qrcode.New(websiteURL)
		if err != nil {
			fmt.Printf("could not generate QRCode: %v", err)
			return
		}

		w, err := standard.New("qr_code4.png", standard.WithLogoImageFileJPEG("./S.jpeg"), standard.WithLogoSizeMultiplier(1))
		if err != nil {
			fmt.Printf("standard.New failed: %v", err)
			return
		}

		// save file
		if err = qrc.Save(w); err != nil {
			fmt.Printf("could not save image: %v", err)
		}

	}
	//  else {
	// 	startServer()
	// }
}

// func startServer() {

// 	router := httprouter.New()  // router instance
// 	router.GET("/", sendQRCode) // landing page

// 	fmt.Println("server listening on port 3000")
// 	serverError := http.ListenAndServe(":3000", router)
// 	if serverError != nil {
// 		log.Fatal("Unable to start web server, cause: ", serverError)
// 	}
// }

// func sendQRCode(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

// 	// Generate QR code HTML image tag
// 	qr_code_image_tag := generateQRCodeHtmlImageTag()
// 	// Set response header to Status Ok (200)
// 	response.WriteHeader(http.StatusOK)

// 	// HTML image with embedded QR code
// 	responsePageHtml := "<!DOCTYPE html><html><body><h1>QR Code example Go Dev Tips</h1>" + qr_code_image_tag + "</body></html>"

// 	// Send HTML response back to client
// 	_, _ = response.Write([]byte(responsePageHtml))
// }

// func generateQRCodeHtmlImageTag() string {

// 	qrCodeImageData, taskError := qrcode.Encode(websiteURL, qrcode.High, imageSize)

// 	if taskError != nil {
// 		log.Fatalln("Error generating QR code. ", taskError)
// 	}

// 	// Encode raw QR code data to base 64
// 	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)
// 	fmt.Println("encodedData ", encodedData)
// 	return "<img src=\"data:image/png;base64, " + encodedData + "\">"
// }
