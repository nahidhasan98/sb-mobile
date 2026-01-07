package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	discordtexthook "github.com/nahidhasan98/discord-text-hook"
	"github.com/nahidhasan98/sb-mobile/app/api"
	"github.com/nahidhasan98/sb-mobile/config"
)

var counterName map[string]string = map[string]string{
	"129": "Alamdanga",
	"79":  "Allardarga",
	"86":  "Allardarga 2",
	"44":  "Amla",
	"116": "Amla Fram Road",
	"53":  "Baipail",
	"132": "Baipail-2",
	"47":  "Bamundi",
	"142": "Bamundi-2",
	"78":  "Baut Bazer",
	"43":  "Bheramara",
	"135": "Bismail",
	"51":  "Bittipara",
	"124": "Bittipara-2",
	"131": "Bollovpur",
	"77":  "Bonpara",
	"72":  "Bus Stuff",
	"52":  "Chandra",
	"119": "Chandra-2",
	"118": "Chariarbil Bazer",
	"137": "Chourhash",
	"83":  "Dangmorka Bazer",
	"58":  "Dashuria",
	"141": "Dharmodah",
	"50":  "Fulbaria",
	"37":  "Gabtoli",
	"110": "Gabtoli 3",
	"48":  "Gangni",
	"57":  "Garagong",
	"114": "Garagong-2",
	"87":  "Garura",
	"125": "Genda Savar",
	"130": "Halsa",
	"128": "Hardi",
	"127": "Hatboalia",
	"75":  "Hemayetpur",
	"133": "Hemayetpur-2",
	"81":  "Hosenabad",
	"143": "Jamjami Bazar",
	"112": "Jhenidah-1",
	"136": "Jorepukuria",
	"38":  "Kallyanpur",
	"85":  "Kazipur",
	"1":   "Khalek Pump",
	"60":  "Khalisakundi",
	"107": "Khoksha",
	"106": "Kumarkhali",
	"42":  "Kushtia",
	"120": "Kushtia-3",
	"126": "Lahini Bottola",
	"40":  "Lakkhipur",
	"67":  "Magura",
	"89":  "Magura Oabda More",
	"138": "MANIKGANJ",
	"64":  "Meherpur",
	"46":  "Mirpur",
	"121": "Mirpur-10",
	"82":  "Mothurapur",
	"71":  "Nabinagar",
	"134": "Nabinagar-2",
	"109": "Online-2",
	"140": "Palli Bidyut Nabinagor",
	"84":  "Pragpur",
	"59":  "Ruppur",
	"49":  "Savar",
	"56":  "Shailkupa",
	"123": "Shailkupa-2",
	"41":  "Shekpara",
	"111": "SREEPUR",
	"55":  "Tangail",
	"80":  "Taragunia",
	"74":  "Vadalia",
	"54":  "Zirani Bazar",
	"139": "Zirani Bazar-2",
}

var stationName map[string]string = map[string]string{
	"54": "Airpot",
	"45": "Alamdanga",
	"40": "Allardorga",
	"36": "Arappur",
	"22": "Baipail",
	"1":  "Bheramara",
	"47": "Bollovpur",
	"4":  "Bonpara",
	"20": "Chandra",
	"43": "Chittagong",
	"34": "Chuadanga",
	"2":  "Dashuria",
	"25": "Dhaka",
	"26": "Gabtoli",
	"28": "Gangni",
	"18": "Garagonj",
	"50": "GULISTAN",
	"46": "Halsa",
	"44": "Hatboalia",
	"51": "JATRABARI",
	"32": "Jhenaidah",
	"23": "Jirani",
	"27": "Kallayanpur",
	"39": "Kazipur",
	"13": "Khoksha",
	"49": "KOLA BAGAN",
	"14": "Kumarkhali",
	"15": "Kushtia",
	"17": "Lakkhipur",
	"33": "Magura",
	"9":  "Manikgonj",
	"31": "Meherpur",
	"35": "Mujibnagar",
	"37": "Nabinagar",
	"10": "Pangsa",
	"12": "Pangsha",
	"21": "Paturia",
	"41": "Pragpur",
	"11": "Rajbari",
	"30": "Ruppur",
	"8":  "Savar",
	"52": "Sayedabad",
	"19": "Shailkupa",
	"16": "Sheikhpara",
	"6":  "Sirajgonj",
	"7":  "Tangail",
	"24": "Trimohini",
}

type schedulePayload struct {
	CounterId   string `json:"counterId"`
	StationId   string `json:"stationId"`
	JourneyDate string `json:"journeyDate"`
}

func GetStationsByCounter(c *gin.Context) {
	counterId := c.Param("id")
	_, err := strconv.Atoi(counterId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid counter id",
		})
		return
	}

	stations, err := api.GetStationsByCounter(counterId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   stations,
	})
}

func sendToDiscord(logMsg string) {
	msg := "```md\n"
	msg += logMsg
	msg += "```"

	// innitializing webhook
	webhook := discordtexthook.NewDiscordTextHookService(config.WebhookIDLogger, config.WebhookTokenLogger)

	// sending msg to discord
	_, err := webhook.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}
}

func GetSchedule(c *gin.Context) {
	var postData schedulePayload
	err := c.BindJSON(&postData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid JSON object",
		})
		return
	}
	userId := "10000603100"

	schedule, err := api.GetSchedule(postData.CounterId, postData.StationId, postData.JourneyDate, userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "something went wrong, please try again",
		})
		return
	}

	// logging client ip address and request payload
	clientIP := c.ClientIP()
	logMsg := fmt.Sprintf("Client IP address: %v | Payload: %v > %v > %v", clientIP, counterName[postData.CounterId], stationName[postData.StationId], postData.JourneyDate)
	log.Println(logMsg)
	go sendToDiscord(logMsg)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   schedule,
	})
}

var ticket map[string]string = map[string]string{
	"operator":      "SB Super Deluxe",
	"date":          "05-Apr-2025",
	"time":          "11:00 AM",
	"coach":         "5015-KE",
	"seat":          "E1, E2",
	"type":          "AC (Economy)",
	"price":         "2400",
	"route":         "Bheramara-Kushtia-Rajbari-Dhaka",
	"contact":       "01760781145",
	"facebook":      "https://www.facebook.com/nahid.achromatic98",
	"facebook_name": "Nahid Hasan",
}

func Index(c *gin.Context) {
	// Check if request is from old domain to show migration notice
	host := c.Request.Host
	showMigrationNotice := host == "sb-mobile.ajudge.net"

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":               "Mobile | SB Super Deluxe",
		"Notice":              false,
		"Ticket":              ticket,
		"ShowMigrationNotice": showMigrationNotice,
	})
}
