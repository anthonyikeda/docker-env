package main

import (
  "testing"
)

func TestListConfig(t *testing.T) {
  service := Services {
            Services: []DockerConf {
              DockerConf { Name: "service1", Host: "localhost:2354", TlsVerify: true, CertPath: "/home/certs" },
            },
          }

  t.Log("Given a struct and location save the file"); {
    err := ListConfig(service);

    if err != nil {
      t.Error("Error occured")
    }
  }
}
