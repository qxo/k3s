package cmds

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

type Agent struct {
	Token                    string
	TokenFile                string
	ServerURL                string
	DataDir                  string
	NodeIP                   string
	NodeName                 string
	ClusterSecret            string
	Docker                   bool
	ContainerRuntimeEndpoint string
	NoFlannel                bool
	Debug                    bool
	AgentShared
}

type AgentShared struct {
	NodeIP string
}

var (
	appName     = filepath.Base(os.Args[0])
	AgentConfig Agent
	NodeIPFlag  = cli.StringFlag{
		Name:        "node-ip,i",
		Usage:       "(agent) IP address to advertise for node",
		Destination: &AgentConfig.NodeIP,
	}
	NodeNameFlag = cli.StringFlag{
		Name:        "node-name",
		Usage:       "(agent) Node name",
		EnvVar:      "K3S_NODE_NAME",
		Destination: &AgentConfig.NodeName,
	}
	DockerFlag = cli.BoolFlag{
		Name:        "docker",
		Usage:       "(agent) Use docker instead of containerd",
		Destination: &AgentConfig.Docker,
	}
	FlannelFlag = cli.BoolFlag{
		Name:        "no-flannel",
		Usage:       "(agent) Disable embedded flannel",
		Destination: &AgentConfig.NoFlannel,
	}
	CRIEndpointFlag = cli.StringFlag{
		Name:        "container-runtime-endpoint",
		Usage:       "(agent) Disable embedded containerd and use alternative CRI implementation",
		Destination: &AgentConfig.ContainerRuntimeEndpoint,
	}
)

func NewAgentCommand(action func(ctx *cli.Context) error) cli.Command {
	return cli.Command{
		Name:      "agent",
		Usage:     "Run node agent",
		UsageText: appName + " agent [OPTIONS]",
		Action:    action,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "token,t",
				Usage:       "Token to use for authentication",
				EnvVar:      "K3S_TOKEN",
				Destination: &AgentConfig.Token,
			},
			cli.StringFlag{
				Name:        "token-file",
				Usage:       "Token file to use for authentication",
				EnvVar:      "K3S_TOKEN_FILE",
				Destination: &AgentConfig.TokenFile,
			},
			cli.StringFlag{
				Name:        "server,s",
				Usage:       "Server to connect to",
				EnvVar:      "K3S_URL",
				Destination: &AgentConfig.ServerURL,
			},
			cli.StringFlag{
				Name:        "data-dir,d",
				Usage:       "Folder to hold state",
				Destination: &AgentConfig.DataDir,
				Value:       "/var/lib/rancher/k3s",
			},
			cli.StringFlag{
				Name:        "cluster-secret",
				Usage:       "Shared secret used to bootstrap a cluster",
				Destination: &AgentConfig.ClusterSecret,
				EnvVar:      "K3S_CLUSTER_SECRET",
			},
			DockerFlag,
			FlannelFlag,
			NodeNameFlag,
			NodeIPFlag,
			CRIEndpointFlag,
		},
	}
}
