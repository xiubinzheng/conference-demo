package main

import (
	//"strings"

	"./alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request alexa.Request) (alexa.Response, error) {
	alexa.LogObject("Request", request)
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}

func HandleStartIntent(request alexa.Request) alexa.Response {

	var builder alexa.SSMLBuilder

	slots := request.Body.Intent.Slots
	PNCur := slots["PN"].Value
	BeAnywhereCur := slots["BeAnywhere"].Value

	if PNCur != "" || BeAnywhereCur != "" {

		if PNCur != "" {

			builder.Say("Your conference was started on ")
			builder.PN(PNCur)
			builder.Say(". ")

		} else if BeAnywhereCur != "" {
			builder.Say("Your conference was started on " + BeAnywhereCur + ". ")
		}

		return alexa.NewSSMLResponse("StartIntent Device", builder.Build(), "", true)

	} else {

		var builderReprompt alexa.SSMLBuilder

		builder.Say("Which device would you like to start your conference on? ")

		builder.Say("You can say a telephone number, such as ")
		builder.PN("2678157599")
		builder.Say(", or say a Be Anywhere device, such as My Cell. ")

		return alexa.NewSSMLResponse("StartIntent NoDevices", builder.Build(), builderReprompt.Build(), false)

	}

}

func HandleStopIntent(request alexa.Request) alexa.Response {

	var builder alexa.SSMLBuilder

	slots := request.Body.Intent.Slots
	PNCur := slots["PN"].Value
	BeAnywhereCur := slots["BeAnywhere"].Value

	if PNCur != "" {

		builder.Say("Your conference was stopped on ")
		builder.PN(PNCur)
		builder.Say(". ")

	} else if BeAnywhereCur != "" {

		builder.Say("Your conference was stopped on " + BeAnywhereCur + ". ")

	} else {

		builder.Say("Your conference was stopped. ")

	}

	return alexa.NewSSMLResponse("StopIntent", builder.Build(), "", true)

}

func HandleLaunchRequest(request alexa.Request) alexa.Response {

	var builder alexa.SSMLBuilder

	builder.Say("My Conference skill can start an Audio conference on your main line,a BeAnywhere phone, or any telephone number you like. You can say, 'Alexa, ask My conference to start a conference on my main line.' Or, to start a conference on any one of your BeAnywhere phones, such as 'mobile,' say 'Alexa, ask My Conference to start a conference on my mobile.' You will need to be a Comcast Business VoiceEdge subscriber to use this skill")
	//builder.PN("2155551234")

	return alexa.NewSSMLResponse("LaunchRequest", builder.Build(), "", true)

}

func IntentDispatcher(request alexa.Request) alexa.Response {

	var response alexa.Response

	if request.Body.Type == "LaunchRequest" {
		response = HandleLaunchRequest(request)
	} else if request.Body.Type == "IntentRequest" {

		switch request.Body.Intent.Name {
		case "StartIntent":
			response = HandleStartIntent(request)
		case "StopIntent":
			response = HandleStopIntent(request)
		}

	}

	return response

}

/*
package main

import (
	"fmt"
	//"encoding/json"
	"strings"
)

type Info struct {
	StartIntent bool
	PN string
	BeAnywhere string
}

type UserInfo struct {
	UserId string
	Info Info
}

type CardImg struct {
	Small string
	Large string
}

//<a href="https://www.iconfinder.com/icons/309047/conference_group_people_users_icon" target="_blank">"Conference, group, people, users icon"</a> by <a href="https://www.iconfinder.com/visualpharm" target="_blank">Ivan Boyko</a> is licensed under <a href="http://creativecommons.org/licenses/by/3.0" target="_blank">CC BY 3.0</a>
//"Conference, group, people, users icon" (https://www.iconfinder.com/icons/309047/conference_group_people_users_icon) by Ivan Boyko (https://www.iconfinder.com/visualpharm) is licensed under CC BY 3.0 (http://creativecommons.org/licenses/by/3.0)
var conferenceImg CardImg
conferenceImg.Small = "https://s3.amazonaws.com/audio-conference/images/conferenceSmall.png"
conferenceImg.Large = "https://s3.amazonaws.com/audio-conference/images/conferenceLarge.png"

//<a href="https://www.iconfinder.com/icons/3324959/outgoing_phone_icon" target="_blank">"Outgoing, phone icon"</a> by <a href="https://www.iconfinder.com/colebemis" target="_blank">Cole Bemis</a> is licensed under <a href="http://creativecommons.org/licenses/by/3.0" target="_blank">CC BY 3.0</a>
//"Outgoing, phone icon" (https://www.iconfinder.com/icons/3324959/outgoing_phone_icon) by Cole Bemis (https://www.iconfinder.com/colebemis) is licensed under CC BY 3.0 (http://creativecommons.org/licenses/by/3.0)
var phoneStartImg CardImg
phoneStartImg.Small = "https://s3.amazonaws.com/audio-conference/images/phoneStartSmall.png"
phoneStartImg.Large = "https://s3.amazonaws.com/audio-conference/images/phoneStartLarge.png"

//<a href="https://www.iconfinder.com/icons/3324961/missed_phone_icon" target="_blank">"Missed, phone icon"</a> by <a href="https://www.iconfinder.com/colebemis" target="_blank">Cole Bemis</a> is licensed under <a href="http://creativecommons.org/licenses/by/3.0" target="_blank">CC BY 3.0</a>
//"Missed, phone icon" (https://www.iconfinder.com/icons/3324961/missed_phone_icon) by Cole Bemis (https://www.iconfinder.com/colebemis) is licensed under CC BY 3.0 (http://creativecommons.org/licenses/by/3.0)
var phoneStopImg CardImg
phoneStopImg.Small = "https://s3.amazonaws.com/audio-conference/images/phoneStopSmall.png"
phoneStopImg.Large = "https://s3.amazonaws.com/audio-conference/images/phoneStopLarge.png"

//<a href="https://www.iconfinder.com/icons/3324960/off_phone_icon" target="_blank">"Off, phone icon"</a> by <a href="https://www.iconfinder.com/colebemis" target="_blank">Cole Bemis</a> is licensed under <a href="http://creativecommons.org/licenses/by/3.0" target="_blank">CC BY 3.0</a>
//"Off, phone icon" (https://www.iconfinder.com/icons/3324960/off_phone_icon) by Cole Bemis (https://www.iconfinder.com/colebemis) is licensed under CC BY 3.0 (http://creativecommons.org/licenses/by/3.0)
var phoneErrorImg CardImg
phoneErrorImg.Small = "https://s3.amazonaws.com/audio-conference/images/phoneErrorSmall.png"
phoneErrorImg.Large = "https://s3.amazonaws.com/audio-conference/images/phoneErrorLarge.png"

//<a href="https://www.iconfinder.com/icons/183285/help_mark_question_icon" target="_blank">"Help, mark, question icon"</a> by <a href="https://www.iconfinder.com/yanlu" target="_blank">Yannick Lung</a>
//"Help, mark, question icon" (https://www.iconfinder.com/icons/183285/help_mark_question_icon) by Yannick Lung (https://www.iconfinder.com/yanlu)
var questionImg CardImg
questionImg.Small = "https://s3.amazonaws.com/audio-conference/images/questionSmall.png"
questionImg.Large = "https://s3.amazonaws.com/audio-conference/images/questionLarge.png"


func main() {
	fmt.Println("Hello!")
}
*/
