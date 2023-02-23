#!/usr/bin/env bash

su - db2inst1 -c "
  db2 create database g2 using codeset utf-8 territory us;
"
