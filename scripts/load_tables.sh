#!/bin/bash

echo "load sql tables"
mysql -u root -p myDatabase < my_backup.sql
