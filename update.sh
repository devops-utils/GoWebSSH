#!/bin/bash
rm -rf gossh/webroot/*
cp -r webssh/dist/* gossh/webroot
git add gossh/webroot
git commit -a -m 'update'
git push -u origin main