@echo off
@REM 清屏
cls
@REM 指定BAT文件为UTF-8编码
chcp 65001
@REM 指定系统类型
set GOOS=windows
@REM 指定系统架构
set GOARCH=amd64
@REM 指定当前为开发模式
set BEDISK_IS_DEV=1
@REM 执行程序运行
go run main.go