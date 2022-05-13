@ubuntu-net
@probes/ubuntu/filesystem
Feature: VM file systems best practice

  Background:
    Given an Ubuntu VM must be up
    
    @ubuntu-filesystem-001
    Scenario: Ensure that VM File System is CIS compliance
        Then run "<COMMAND>" and then output is  "<OUTPUT>"
    
  Examples:
| COMMAND                         |     OUTPUT                       |
| dpkg -s sudo \| grep Status     |   Status: install ok installed   |
| sudo -s \|  df -h               | /var                             | 
| sudo -s \|  df -h               | /var/tmp                         | 
| sudo -s \|  df -h               | /var/log                         | 
| sudo -s \|  df -h               | /var/audit                       | 
| sudo -s \|  df -h               | /home                            | 
      