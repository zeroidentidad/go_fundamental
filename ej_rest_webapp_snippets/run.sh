#!/bin/bash
export S_ADDR=":4001"
export S_DSN="webusr:passx123@tcp(localhost:3306)/pastein?parseTime=true"
export S_SECRET="z3Roh+pPbnzHbS*+9Pk8qGWhTzbpa@jf"
go run . -addr=$S_ADDR -dsn=$S_DSN -secret=$S_SECRET