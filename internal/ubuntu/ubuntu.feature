@ubuntu-net
@probes/ubuntu/ubuntu
Feature: Cluster networking best practice

  Background:
    Given an Ubuntu VM must be up

   @ubuntu-net-001
    Scenario:  ensure SSH root login is disabled
      Then ensure SSH root login is disabled

    @ubuntu-net-002
    Scenario:  Ensure ufw firewall is configured
      Then Ensure ufw firewall is configured

   



  Examples:
| TEST        | 
| test            | 
      