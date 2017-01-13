# Comsat

Comsat is a command line tool that enables you to execute your CI/CD commands in a containerized, distributed environment. It can greatly improve your CI/CD pipeline by allowing you to execute your test or build commands in parallel. Since everything is containerized, you no longer have to configure or setup any slave dependencies.

## Dependencies
 * Docker Engine or Docker Swarm

## Benefits
 * Slaveless: No slaves to configure. Execution is done by creating a container of your code base on a node (or nodes if you're running a Swarm)
 * Concurrency: Execute multiple commands at the same time. For example, a Rails app can execute it's test suite, rubocop and brakeman at the same time rather than in a the typical synchronous fashion.

## Usage:

Once you add a comsat.yml config file, execute the following:

```
  comsat
```

## Example:

If your comsat.yml looks like this

```yml
image: hacker_registration:ctest
command: 
 - rubocop
 - rspec
 - brakeman
```

When you execute `comsat` it will create 3 containers using the specified image and execute each command in parallel.

## Configuration 
In a yaml file, you specify the image you want to test and a list of commands to be executed. Then you execute ``comsat`` which concurrently executes each of the commands in it's own container.

All output is sent to stdout.

Sample comsat.yaml

```yaml
image: your-image-name
command:
 - rspec
 - rubocop
```

## Todos:
* Be able to specify a different DOCKER_HOST endpoint.
* Need a way to specify container dependencies (such as database dependency)
* Need a way to specify setup commands (such as creating or migrating a datbase)

This will require a change to the comsat.yml file. Proposed update:

```yaml
  postgres:
		image: mongo
  web:
    image: myrailsapp:test
    setup: 
      - rake db:create
      - rake db:migrate
    commands:
      - rspec
      - rubocop
```
