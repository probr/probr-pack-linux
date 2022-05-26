@ubuntu-useraccount
@probes/ubuntu/ubuntu
Feature: VM User accounts and Environment best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-useraccount-001
    Scenario: Ensure that VM User accounts and Environment are CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                                 |     OUTPUT          |
| grep PASS_MAX_DAYS /etc/login.defs      |  PASS_MAX_DAYS 365  |
| useradd -D \| grep INACTIVE             |  INACTIVE=30        |

      