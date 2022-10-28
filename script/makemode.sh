#!/bin/bash
goctl model mysql datasource -url="$database" -table="*" -c -dir common/entitys -style "goZero"