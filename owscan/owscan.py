# !/usr/bin/env python
# -*- coding:utf-8 -*-

import sys
import requests


def who_is(website):
    url = "http://api.hackertarget.com/whois/?q=" + website
    r = requests.get(url)
    print(r.text)


def dns_lookup(website):
    url = "http://api.hackertarget.com/dnslookup/?q=" + website
    r = requests.get(url)
    print(r.text)


def zone_transfer(website):
    url = "http://api.hackertarget.com/zonetransfer/?q=" + website
    r = requests.get(url)
    print(r.text)


def mtr(website):
    url = "https://api.hackertarget.com/mtr/?q=" + website
    r = requests.get(url)
    print(r.text)


def nmap(website):
    url = "http://api.hackertarget.com/nmap/?q=" + website
    r = requests.get(url)
    print(r.text)


def page_links(website):
    url = "https://api.hackertarget.com/pagelinks/?q=http://" + website
    r = requests.get(url)
    print(r.text)


def host(website):
    url = "https://tools.keycdn.com/geo.json?host=" + website
    r = requests.get(url)
    print(r.text)


def http_headers(website):
    url = "http://api.hackertarget.com/httpheaders/?q=" + website
    r = requests.get(url)
    print(r.text)


def main(website):
    who_is(website)
    dns_lookup(website)
    zone_transfer(website)
    mtr(website)
    nmap(website)
    page_links(website)
    host(website)
    http_headers(website)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python " + sys.argv[0] + " www.example.com")
    else:
        main(sys.argv[1])
