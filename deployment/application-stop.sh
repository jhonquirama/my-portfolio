#!/usr/bin/env bash

systemctl-exists() {
  [ $(systemctl list-unit-files "${1}*" | wc -l) -gt 3 ]
}

systemctl-exists hexagonal-scaffolding

if [ $? -eq 0 ]
then
  systemctl stop hexagonal-scaffolding
fi