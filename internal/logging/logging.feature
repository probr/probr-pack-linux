@ubuntu-network
@probes/ubuntu/logging
Feature: VM Logging and Auditing best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-logging-001
    Scenario: Ensure that VM logging is CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                        | OUTPUT                 |
| systemctl is-enabled auditd    | enabled                |
| systemctl is-enabled rsyslog   | enabled                |
| dpkg -s rsyslog                | install ok installed   |


      