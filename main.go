package main

import (
  "flag"
  "fmt"
  "os"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "errors"
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
  switchCommand := flag.NewFlagSet("switch", flag.ExitOnError)

  nameCommandPtr := saveCommand.String("name", "", "Name to save config as")
  hostCommandPtr := saveCommand.String("host", "", "Docker Host and port")
  tlsVerifyCommandPtr := saveCommand.Bool("tls-verify", true, "TLS Verify")
  certsPathCommandPtr := saveCommand.String("cert-path", "", "Path to certs")

  if len(os.Args) < 2 {
    flag.PrintDefaults()
    os.Exit(1)
  }

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
  case "switch":
    switchCommand.Parse(os.Args[2:])
    fmt.Printf("Switching to machine \"%s\"\r\n", os.Args[2])
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

    services, loadErr := LoadConfig(CONFIG_PATH)

    if loadErr != nil {
      fmt.Println(loadErr)
      services = Services {
        Services: []DockerConf {
          DockerConf {*nameCommandPtr, *hostCommandPtr, *tlsVerifyCommandPtr, *certsPathCommandPtr},
        },
      }
    } else {
      services.Services = append(services.Services, DockerConf {*nameCommandPtr, *hostCommandPtr, *tlsVerifyCommandPtr, *certsPathCommandPtr})
    }

    err := SaveConfig(services, CONFIG_PATH)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }

  if listCommand.Parsed() {
    services, loadErr := LoadConfig(CONFIG_PATH)

    if loadErr != nil {
      fmt.Println(loadErr)
      os.Exit(1)
    }

    err := ListConfig(services);
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }
}

func ListConfig(services Services) error {

  if services.Services == nil {
    return errors.New("Services is not initialised")
  }

  fmt.Printf("%s | %s | \r\n", "Name", "Host")
  for _, v := range services.Services {
    fmt.Printf("%s | %s | \r\n", v.Name, v.Host)
  }

  return nil
}

func SaveConfig(c Services, filename string) error {
  bytes, err := yaml.Marshal(c)
  if err != nil {
    return err
  }
  return ioutil.WriteFile(filename, bytes, 0644)
}

func LoadConfig(filename string) (Services, error) {
  var s Services

  services := Services{}
  bytes, err := ioutil.ReadFile(filename)

  if err != nil  {
    return s, err
  }

  error := yaml.Unmarshal(bytes, &services)

  if error != nil {
    return s, error
  }

  return services, nil
}
