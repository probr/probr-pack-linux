@ubuntu-network
@probes/ubuntu/ubuntu
Feature: Cluster networking best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-network-001
    Scenario: Ensure that VM network is CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                         |     OUTPUT                       |
| dpkg -s sudo \| grep Status     |   Status: install ok installed   |

      