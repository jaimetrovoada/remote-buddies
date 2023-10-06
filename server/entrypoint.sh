#!/bin/bash

goose -v -dir /usr/src/app/migrations postgres $DATABASE_URL up
air
