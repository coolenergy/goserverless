package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/sns"
)

func main() {
	template := cloudformation.NewTemplate()

	template.Resources["MyTopic"] = &sns.Topic{
		TopicName: "my-topic-" + strconv.FormatInt(time.Now().Unix(), 10),
	}

	template.Resources["MyTopicSubscription"] = &sns.Subscription{
		TopicArn: cloudformation.Ref("MyTopic"),
		Protocol: "email",
		Endpoint: "some.email@example.com",
	}

	// Let's see the JSON AWS CloudFormation template
	j, err := template.JSON()
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(j))
	}

	// and also the YAML AWS CloudFormation template
	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}
}
