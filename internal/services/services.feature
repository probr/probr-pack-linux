@ubuntu-service
@probes/ubuntu/ubuntu
Feature: VM Services best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-service-001
    Scenario: Ensure that VM services are CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                                                                                                                           |     OUTPUT            |
| sudo -s \|dpkg -s vsftpd \| grep -E                                                                                               |   is not installed    |
| sudo -s \|dpkg -s apache2 | grep -E                                                                                               |   is not installed    |
| grep -E -s "^\s*net\.ipv4\.ip_forward\s*=\s*1" /etc/sysctl.conf /etc/sysctl.d/*.conf /usr/lib/sysctl.d/*.conf /run/sysctl.d/*.conf|                       |

      