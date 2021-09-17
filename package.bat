@echo off
@REM set time_hh=%time:~0,2%
@REM if /i %time_hh% LSS 10 (set time_hh=0%time:~1,1%)
set filename=%date:~3,4%%date:~8,2%%date:~11,2%
echo %filename%
zip -r -q go_tenancy_linux_%filename%.zip ./www  ./main