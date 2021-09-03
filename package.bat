@echo off
set time_hh=%time:~0,2%
if /i %time_hh% LSS 10 (set time_hh=0%time:~1,1%)
set filename=%date:~,4%%date:~5,2%%date:~8,2%
zip -r -q go_tenancy_linux_%filename%.zip ./www  ./main