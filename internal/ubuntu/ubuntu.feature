@ubuntu-net
@probes/ubuntu/ubuntu
Feature: Cluster networking best practice

  Background:
    Given an Ubuntu VM must be up

    @ubuntu-net-001
    Scenario:  gnome display Manager is disabled
      Then gnome display Manager is disabled
      