# Installing Docker (Virtualbox, boot2docker, docker)

1. Download and install the latest [boot2docker](https://github.com/boot2docker/osx-installer/releases) for OS X.
2. Open a new Terminal and enter `boot2docker init`
3. Run `boot2docker start`
4. You will be given an environment variable to export `export DOCKER_HOST=tcp://192.168.59.103:2375`. Add that to your bash_profile or equivalent.
5. You have a successful environment when the output of `docker version` is clean (exits 0). The output should look a bit like:

```
Client version: 1.1.0
Client API version: 1.13
Go version (client): go1.2.1
Git commit (client): 79812e3
Server version: 1.1.0
Server API version: 1.13
Go version (server): go1.2.1
Git commit (server): 79812000
```

Pull an image and verify that something simple like an `ls` command works (many layers will be pulled down if this is your first execution of docker):

```
$ docker run -i -t ubuntu bash
root@621c9e1d4fe6:/# ls
bin  boot  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```

The Telephone and Answering Machine Exercises will help bootstrap your Docker skills. First, we will fix the Telephone application, allowing it to talk to
your teacher's running Answering Machine. The results will be that you now have a basic understanding of the Dockerfile, docker commands, and you will have
posted a message on a running Answering Machine.

## Basic Docker Commands

* Help for any command - `docker help <command>`
* Building a Dockerfile in CWD - `docker build .`
* Listing running docker containers - `docker ps`
* Running a container - `docker run <container name or hash> <command>`
  * Interactively running a container - `docker run -i -t ...`
  * Detaching a running container - `docker run -d ...`

## Exercise 1 - Fixing and running the Telephone Application

1. Install [Git](http://sourceforge.net/projects/git-osx-installer/)
2. Clone the Telephone Github repository - `git clone https://github.com/KAllan357/telephone.git`
3. Let's complete the application:
4. Fix the Dockerfile according to the comments.
5. Address the FIXME's in the `telephone.rb`.

## Exercise 1 - Key Takeaways

* Basic Dockerfile syntax (FROM, MAINTAINER RUN, ADD, CMD)
* Basic docker commands (build, run)

## Exercise 1 - Help

* Depending on your FROM (base OS), there might be a Ruby package available from your OS's package manager. You might need the build-essential or ruby-dev packages!
* Layering can help your `docker build .` commands run faster.
* How do we ADD telephone.rb?
* How do we execute a Ruby file for our command?

By now, you have successfully POSTed a message to the presenter's running Answering Machine! Horray!

## Exercise 2 - Run your own Answering Machine

1. Clone the Answering Machine Github repository `git clone https://github.com/KAllan357/answering-machine.git`
2. Build and start a container running the Answering Machine application.
3. Try curling the Answering Machine. Does it work on localhost? Does it work on the container's IP? (`docker inspect --format '{{ .NetworkSettings.IPAddress }}' <container id>`)
4. Edit the Answering Machine's Dockerfile, rebuild it and learn how to EXPOSE a port.
5. Learn how to forward a port with the docker run's `-p` flag. Does curling work now?
6. The basic Answering Machine is "Docker's" - you want this to be YOUR answering machine! Notice that [the code](https://github.com/KAllan357/answering-machine/blob/master/main.go#L54) lets us know that an environment variable is exposed!
7. Learn how to pass an environment variable through docker run's `-e` flag. Who's Answering Machine is it now?

## Exercise 2 - An Advanced Answering Machine

Next, we want to Link our own Answering Machine to our own Telephone.

1. Run your Answering Machine container with a `--name` flag.
2. Next, you'll want to run your Telephone application with the `--link` flag, linking the container name you used in step 1 to what your telephone.rb tries to connect to.
3. Finally, check out what the `docker logs <container>` gives you for your Answering Machine application.

## Exercise 2 - Key Takeaways

* Exposing a port when running a container
* Setting an environment variable when running a container
* Linking two containers together
* Docker Logs
