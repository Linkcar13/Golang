package main

import (
	"fmt"
	"log"
	"os/exec" //execute commands
)

//https://img.buzzfeed.com/buzzfeed-static/static/2017-08/8/10/campaign_images/buzzfeed-prod-fastlane-02/39-bellisimos-tatuajes-de-harry-potter-que-te-har-2-19843-1502203268-2_dblbig.jpg?resize=1200:*
const (
	image  string = "39-bellisimos-tatuajes-de-harry-potter-que-te-har-2-19843-1502203268-2_dblbig.jpg"
	domain string = "https://img.buzzfeed.com"
	path   string = "/buzzfeed-static/static/2017-08/8/10/campaign_images/buzzfeed-prod-fastlane-02/"
	url    string = domain + path + image
)

func main() {
	Output, err := exec.Command("wget", url).CombinedOutput()
	if err != nil {
		log.Fatalln("error, File doesn't exist")
	}
	fmt.Println(string(Output))
}
