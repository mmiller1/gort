package config

import (
	"fmt"
	"testing"

	"github.com/clockworksoul/cog2/data"
	yaml "gopkg.in/yaml.v1"
)

func Test(t *testing.T) {
	config := data.CogConfig{}

	config.SlackProviders = make([]data.SlackProvider, 1)

	config.SlackProviders[0] = data.SlackProvider{
		BotName:       "Cog2",
		Name:          "ClockworkSoul",
		IconURL:       "https://emoji.slack-edge.com/T025151EM/dragongopher/cdc9b1bd1a7752eb.png",
		SlackAPIToken: "SlackAPIToken",
	}

	config.BundleConfigs = make([]data.Bundle, 1)

	config.BundleConfigs[0] = data.Bundle{
		Name:        "test",
		Description: "A description",
	}

	config.BundleConfigs[0].Docker = data.BundleDocker{
		Image: "clockworksoul/foo",
		Tag:   "latest",
	}

	config.BundleConfigs[0].Commands = make(map[string]data.BundleCommand)

	config.BundleConfigs[0].Commands["splitecho"] = data.BundleCommand{
		Description: "Echos back anything sent to it, one parameter at a time.",
		Executable:  []string{"/opt/app/splitecho.sh"},
	}

	config.BundleConfigs[0].Commands["curl"] = data.BundleCommand{
		Description: "The official curl command",
		Executable:  []string{"/usr/bin/curl"},
	}

	config.BundleConfigs[0].Commands["echo"] = data.BundleCommand{
		Description: "Echos back anything sent to it, all at once.",
		Executable:  []string{"/bin/echo"},
	}

	y, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println(string(y))
}
