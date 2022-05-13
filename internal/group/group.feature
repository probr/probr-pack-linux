@ubuntu-group
@probes/ubuntu/group
Feature: VM user and Group best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-group-001
    Scenario: Ensure that VM User and Group are CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                                                                               | OUTPUT  |
| awk -F: '($2 != "x" ) { print $1 " is not set to shadowed passwords "}' /etc/passwd   |         |

      