#!/usr/bin/env python3
import string


with open('pl.json') as file1:
    pl = file1.readlines()
with open('en.json') as file2:
    en = file2.readlines()
error_lines = {}
for n in range(en.__len__()):
    key_pl = pl[n][pl[n].find('\"'):pl[n].find(":")]
    key_en = en[n][en[n].find('\"'):en[n].find(":")]
    if(key_pl != key_en):
        error_lines[n+1] = key_pl + "-" + key_en

if(error_lines):
    print('Theese lines have different keys')
    for error in error_lines:
        print(error, ":" ,error_lines[error])
else:
    print("Everything seems ok")