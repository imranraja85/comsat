# Comsat

Comsat is a command line tool that enables you execute test suites inside in containers on a Docker swarm. This allows you to greatly simplify your CI/CD process as the whole test environment can be configured by a Dockerfile rather than configuring slaves.

## Usage:

```
  comsat
```

## Configuration 
In a yaml file, you specify the image you want to test and a list of test runner commands to be executed. Then you execute ``comsat`` which runs a container and executes your test runners.

All output is sent to stdout.

Sample comsat.yaml

```yaml
image: your-image-name
command:
 - rspec
 - rubocop
```

## TODOS:
Version 2 comsat.yml:

 - How to make this OO and add tests????
 - Get the right exit status after each container run
 - exit immediately if an exit status of 1?

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
