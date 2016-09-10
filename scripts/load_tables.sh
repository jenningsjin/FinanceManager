#!/bin/bash

echo "load sql tables"
MySQL -u "finance" --password="password" "FINANCE_APP" < clean.sql