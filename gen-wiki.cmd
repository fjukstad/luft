REM @ECHO OFF
PUSHD "%~dp0."
PUSHD wiki
npm install
node airbit.wiki.js
POPD
IF EXIST public\wiki RD /Q public\wiki
MKLINK /J public\wiki wiki\_site
POPD