@ubuntu-sshaccess
@probes/ubuntu/ubuntu
Feature: VM SSH Access best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-sshaccess-001
    Scenario: Ensure that VM SSH access is CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                                                                                              | OUTPUT   |
| grep -is 'loglevel' /etc/ssh/sshd_config /etc/ssh/sshd_config.d/*.conf \| grep -Evi '(VERBOSE|INFO)' |          |

      