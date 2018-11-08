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

		builderReprompt.Say("You can say a telephone number, such as ")
		builderReprompt.PN("2678157599")
		builderReprompt.Say(", or say the name of a Be Anywhere device, such as My Cell. ")

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
