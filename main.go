package main

import (
  "flag"
  "fmt"
  "os"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type Services struct {
  Services []DockerConf `yaml:"services"`
}
type DockerConf struct {
  Name string `yaml:"name"`
  Host string `yaml:"host"`
  TlsVerify bool `yaml:"tlsVerify"`
  CertPath string `yaml:"certPath"`
}

func main() {
  CONFIG_PATH := fmt.Sprintf("%s/.docker-env", os.Getenv("HOME"))
  listCommand := flag.NewFlagSet("list", flag.ExitOnError)
  saveCommand := flag.NewFlagSet("save", flag.ExitOnError)
  applyCommand := flag.NewFlagSet("apply", flag.ExitOnError)

  nameCommandPtr := saveCommand.String("name", "", "Name to save config as")
  hostCommandPtr := saveCommand.String("host", "", "Docker Host and port")
  tlsVerifyCommandPtr := saveCommand.Bool("tls-verify", true, "TLS Verify")
  certsPathCommandPtr := saveCommand.String("cert-path", "", "Path to certs")

  if len(os.Args) < 2 {
    flag.PrintDefaults()
    os.Exit(1)
  }
//  fmt.Printf("docker-env --help %s, %s, %t\r\n", *machineName, *hostPtr, *tlsVerifyPtr)
//  fmt.Printf("Command supplied: %s\r\n", command)

  switch os.Args[1] {
  case "list":
    fmt.Println("Listing saved configs...")
    listCommand.Parse(os.Args[2:])
  case "save":
    fmt.Printf("Saving config to %s\r\n", CONFIG_PATH)
    saveCommand.Parse(os.Args[2:])
    fmt.Printf("Config name %s\r\n", *nameCommandPtr)
  case "apply":
    applyCommand.Parse(os.Args[2:])
  default:
    flag.PrintDefaults()
    os.Exit(1)
  }

  if saveCommand.Parsed() {
    if *nameCommandPtr == "" {
      saveCommand.PrintDefaults()
      os.Exit(1)
    }

    if *hostCommandPtr == "" {
      saveCommand.PrintDefaults()
      os.Exit(1)
    }

    if *certsPathCommandPtr == "" {
      saveCommand.PrintDefaults()
      os.Exit(1)
    }

    services := Services {
        Services: []DockerConf {
          DockerConf {*nameCommandPtr, *hostCommandPtr, *tlsVerifyCommandPtr, *certsPathCommandPtr},
        },
      }

    err := saveConfig(services, CONFIG_PATH)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }

}

func saveConfig(c Services, filename string) error {
  bytes, err := yaml.Marshal(c)
  if err != nil {
    return err
  }
  return ioutil.WriteFile(filename, bytes, 0644)
}
