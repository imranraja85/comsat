# Comsat

Comsat is a command line tool that enables you parallely execute your CI/CD commands in a containerized, distributed environment. It can greatly improve your CI/CD pipeline by allowing you to execute your test or build commands in parallel. Since everything is containerized, you no longer have to configure or setup any slave dependencies.

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

All output is sent to stdout.
