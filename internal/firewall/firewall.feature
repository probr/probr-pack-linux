@ubuntu-firewall
@probes/ubuntu/firewall
Feature: VM Firewall best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-firewall-001
    Scenario: Ensure that VM firewall is CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                         |     OUTPUT                       |
| sudo ufw status \| grep Status  | inactive                         | 

      