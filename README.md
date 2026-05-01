[![Ad-list-processing](https://github.com/asiaqa/ag-rule/actions/workflows/ad-list.yml/badge.svg)](https://github.com/asiaqa/ag-rule/actions/workflows/ad-list.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.19-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![Updates Per Day](https://img.shields.io/badge/Updates-3%2Fday-brightgreen)](https://github.com/asiaqa/ag-rule/actions)
[![Rules](https://img.shields.io/badge/Rules-1M%2B-blue)](https://github.com/asiaqa/ag-rule)
[![GitHub Forks](https://img.shields.io/github/forks/asiaqa/ag-rule?style=social)](https://github.com/asiaqa/ag-rule/network/members)

# ag-rule

Consolidated adblock rules for AdGuard Home, Pi-hole, dnsmasq, and Unbound. Updated 3 times daily. Because ads never sleep, neither do we.

整合多個來源的廣告封鎖規則，適用於 AdGuard Home、Pi-hole、dnsmasq 及 Unbound。每日更新 3 次，因為廣告不睡覺，我們也不睡。

---

## :sparkles: Features / 功能

| Feature | Description |
|---------|-------------|
| Multi-source | Consolidates 30+ filter lists into one (we did the homework so you don't have to) |
| Multi-format | AdGuard Home, hosts, dnsmasq, Unbound, uBlock Origin, domain-only |
| 4 tiers | Super, Full, Medium, Min — pick your pain level |
| DNS filtering | Removes dead domains to reduce file size (ghost hunting for NXDOMAIN) |
| Transparency | Dead domains list published for your review and curiosity |

---

## :arrow_down: Download Links / 下載連結

All formats, all tiers, one table. Pick your weapon.

| Tier | AdGuard Home | Hosts | dnsmasq | Unbound | uBlock Origin | Domains |
|------|-------------|-------|---------|---------|---------------|---------|
| :fire: **Super** (982k) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_hosts.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_hosts.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_dnsmasq.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_dnsmasq.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_unbound.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_unbound.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_ublock.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_ublock.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_domains.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_domains.txt) |
| :zap: **Full** (927k) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_hosts.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_hosts.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_dnsmasq.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_dnsmasq.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_unbound.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_unbound.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_ublock.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_ublock.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_domains.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_domains.txt) |
| :bulb: **Medium** (352k) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_hosts.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_hosts.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_dnsmasq.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_dnsmasq.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_unbound.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_unbound.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_ublock.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_ublock.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_domains.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_domains.txt) |
| :seedling: **Min** (231k) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_hosts.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_hosts.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_dnsmasq.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_dnsmasq.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_unbound.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_unbound.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_ublock.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_ublock.txt) | [raw](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_domains.txt) \| [cdn](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_domains.txt) |

### :gift: Other Goodies / 其他好東西

| List | Raw URL | CDN |
|------|---------|-----|
| GitHub Hosts (China) | [hosts](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/hosts) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/hosts) |
| DoH Blocklist | [output_doh_block.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_doh_block.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_doh_block.txt) |
| Dead Domains | [dead_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/dead_domains.txt) | (the graveyard) |

---

## :page_facing_up: Format Description / 格式說明

| Format | Example | Use Case |
|--------|---------|----------|
| AdGuard Home | `\|\|ads.example.com^` | AdGuard Home, AdGuard DNS |
| Hosts | `0.0.0.0 ads.example.com` | Pi-hole, `/etc/hosts` |
| dnsmasq | `address=/ads.example.com/#` | OpenWrt, dnsmasq |
| Unbound | `local-zone: "ads.example.com" NXDOMAIN` | Unbound DNS resolver |
| uBlock Origin | `\|\|ads.example.com^` | Browser extensions |
| Domain-only | `ads.example.com` | For the DIY crowd |

---

## :shield: DNS NXDOMAIN Filtering / DNS 過濾

We use [puredns](https://github.com/d3mondev/puredns) to check if domains actually exist. If they don't (NXDOMAIN), we kick them out. Less junk = smaller files = faster DNS.

使用 puredns 檢查網域是否真實存在。如果不存在（NXDOMAIN），就把它們踢出去。更少垃圾 = 更小檔案 = 更快 DNS。

- Dead domains listed in `rules/dead_domains.txt` (the hall of shame)
- Results cached for 24 hours (even robots need rest)
- 1M+ domains? ~5 minutes. Coffee break not included.

---

## :rocket: How to Use / 使用方法

**Pi-hole**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_hosts.txt
```

**dnsmasq (OpenWrt)**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_dnsmasq.txt
```

**Unbound**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_unbound.txt
```

**AdGuard Home**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full.txt
```

---

## :bar_chart: Tier Guide / 等級指南

| Tier | Description | Who should use it |
|------|-------------|-------------------|
| :seedling: **Min** | Essential ad blocking | People who just want to browse in peace |
| :bulb: **Medium** | Balanced protection | Most humans with internet access |
| :zap: **Full** | Aggressive blocking | People who really, really hate ads |
| :fire: **Super** | Maximum blocking | Daredevils who enjoy troubleshooting broken websites |

---

## :heart: Thanks / 致謝

These folks make the internet bearable:

- [AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)
- [adaway.org](https://adaway.org/hosts.txt)
- [easylist](https://easylist-downloads.adblockplus.org/easylist.txt) / [easylistchina](https://easylist-downloads.adblockplus.org/easylistchina.txt) / [easyprivacy](https://easylist-downloads.adblockplus.org/easyprivacy.txt)
- [hagezi/dns-blocklists](https://github.com/hagezi/dns-blocklists)
- [neodevpro/neodevhost](https://github.com/neodevpro/neodevhost)
- [521xueweihan/GitHub520](https://github.com/521xueweihan/GitHub520)
- [dibdot/DoH-IP-blocklists](https://github.com/dibdot/DoH-IP-blocklists)
- [Perflyst/PiHoleBlocklist](https://github.com/Perflyst/PiHoleBlocklist)
- [cjx82630/cjxlist](https://github.com/cjx82630/cjxlist)
- [filter.futa.gg](https://filter.futa.gg/hosts.txt)
- [XIU2/TrackersListCollection](https://github.com/XIU2/TrackersListCollection)

---

## :balance_scale: License

MIT — do whatever you want, just don't blame us if Super tier breaks your favorite website.
