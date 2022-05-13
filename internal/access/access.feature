@ubuntu-access
@probes/ubuntu/ubuntu
Feature: Acceess, authentication and authorization best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-access-001
    Scenario: Ensure that VM Acceess, authentication and authorization are CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                         |     OUTPUT                       |
|  systemctl is-enabled cron      |   enabled   |

      