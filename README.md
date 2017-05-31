[ ![Codeship Status for anthonyikeda/docker-env](https://app.codeship.com/projects/71874790-2852-0135-51bc-3e07a28a8e4e/status?branch=master)](https://app.codeship.com/projects/223289)

# docker-env
Manage different docker environments from the command line

# Usage

Save new config
---

```bash
$ docker-env -host=myhost.aws.com:2376 -tls-verify=true -cert-path=~/mymachine/.certs -name=development save
Saved configuration for development!
```

List configs
---

```bash
$ docker-env list
NAME            HOST                     ENABLED          RUNNING
development     192.168.99.100:2376      true             true
testing         192.168.99.101:2376      false            true
production      docker.ecs.aws.com:2376  true             true
```

Switching machines
---

```bash
$ docker-env use development
Switching to development...
DOCKER_HOST=192.168.99.100:2376
```

* You'll have to forgive the quality of the code, I'm not a typical Go developer :)
