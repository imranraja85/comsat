# Ctest
Ctest is a command line tool that can execute a containerized test suite.

Usage:
```
  ctest
```

How does it work?

In a yaml file, you specify the image you want to test and a list of test runner commands to be executed. Then you execute ``ctest``` which runs a container and executes your test runners.

All output is sent to stdout. If command fails and exit status of 1 will be returned and the container will immediately be terminated.

Version 1 ctest.yml:
```yaml
images: yourimagename
command:
 - rspec
 - rubocop
```

Version 2 ctest.yml:

```yaml
	postgres:
		image: mongo
  web:
    image: myrailsapp:test
    setup: 
      - rake db:create
      - rake db:migrate
    tests:
      - rspec
      - rubocop
```

## Sequences of events
  - read in dockerfile.yaml
  - store contents in a struct
  - create overlay network
  - iterate over services struct and create containers (or services?)
  - execute the given user command and direct all output to stdout
  - remove services
  - remove overlay network
  - profit

# TODO:
 - How to make this OO and add tests????
 
