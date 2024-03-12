package config

import (
	"mediafile/storage_node/utils"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Name              string `yaml:"name"`               // Node name
	ReplicationFactor int    `yaml:"replication_factor"` // Amount of data replicas
	BindHost          string `yaml:"bind_host"`          // Node host
	BindPort          int    `yaml:"bind_port"`          // Node port
	StorageLocation   string `yaml:"storage_location"`   // path of the storage location
	MaxStorage        uint64 `yaml:"max_storage"`        //	max storage size
}

func GetDefault() Node {
	node := Node{}

	// set default node config
	user, err := user.Current()
	if err != nil {
		node.Name = "Node"
	} else {
		node.Name = user.Username
	}

	node.ReplicationFactor = 1
	node.BindHost = "localhost"
	node.BindPort = 5010
	node.StorageLocation = filepath.Join(user.HomeDir, "mediafile", "storage")
	node.MaxStorage = 1073741824

	// set node config from enviorenment vars
	name := os.Getenv("name")
	if name != "" {
		node.Name = name
	}
	replicationFactor := os.Getenv("replicationFactor")
	if replicationFactor != "" {
		res, err := strconv.Atoi(replicationFactor)
		if err == nil {
			node.ReplicationFactor = res
		}
	}
	bindHost := os.Getenv("bindHost")
	if bindHost != "" {
		node.BindHost = bindHost
	}
	bindPort := os.Getenv("bindPort")
	if bindPort != "" {
		res, err := strconv.Atoi(bindPort)
		if err == nil {
			node.BindPort = res
		}
	}
	storageLocation := os.Getenv("storageLocation")
	if storageLocation != "" {
		node.StorageLocation = storageLocation
	}
	maxStorage := os.Getenv("maxStorage")
	if maxStorage != "" {
		res, err := strconv.ParseUint(maxStorage, 10, 64)
		if err == nil {
			node.MaxStorage = res
		}
	}

	return node
}

func GetConfig() (Node, error) {
	// Get default node config
	defaultConfig := GetDefault()

	// Get configuration location
	configLocation := os.Getenv("CONFIG_LOCATION")

	if configLocation == "" {
		user, err := user.Current()
		if err != nil {
			return defaultConfig, err
		}
		configLocation = filepath.Join(user.HomeDir, "mediafile")

		err = utils.CreateFolder(configLocation)
		if err != nil {
			return defaultConfig, err
		}

		configLocation = filepath.Join(configLocation, "node_config.yaml")
	}

	yamlFile, err := os.ReadFile(configLocation)
	if err != nil {
		outputFile, err := os.Create(configLocation)
		if err != nil {
			return defaultConfig, err
		}
		defer outputFile.Close()

		yamlData, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			return defaultConfig, err
		}

		_, err = outputFile.Write(yamlData)
		if err != nil {
			return defaultConfig, err
		}

		err = utils.CreateFolder(defaultConfig.StorageLocation)
		if err != nil {
			return defaultConfig, err
		}
		return defaultConfig, nil
	}

	var node Node
	err = yaml.Unmarshal(yamlFile, &node)
	if err != nil {
		return defaultConfig, err
	}

	if node.Name == "" {
		node.Name = defaultConfig.Name
	}
	if node.ReplicationFactor == 0 {
		node.ReplicationFactor = defaultConfig.ReplicationFactor
	}
	if node.BindHost == "" {
		node.BindHost = defaultConfig.BindHost
	}
	if node.BindPort == 0 {
		node.BindPort = defaultConfig.BindPort
	}
	if node.StorageLocation == "" {
		node.StorageLocation = defaultConfig.StorageLocation
	}
	if node.MaxStorage == 0 {
		node.MaxStorage = defaultConfig.MaxStorage
	}

	err = utils.CreateFolder(node.StorageLocation)
	if err != nil {
		return node, err
	}

	return node, nil
}
