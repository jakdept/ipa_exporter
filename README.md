# ipa_exporter

A Prometheus Exporter to watch a FreeIPA server

Expected coverage / things that will be checked:

```
$ cipa -d ipa.example.com -W ********
+--------------------+----------+----------+----------+-----------+----------+----------+-------+
| FreeIPA servers:   | ipa01    | ipa02    | ipa03    | ipa04     | ipa05    | ipa06    | STATE |
+--------------------+----------+----------+----------+-----------+----------+----------+-------+
| Active Users       | 1199     | 1199     | 1199     | 1199      | 1199     | 1199     | OK    |
| Stage Users        | 0        | 0        | 0        | 0         | 0        | 0        | OK    |
| Preserved Users    | 0        | 0        | 0        | 0         | 0        | 0        | OK    |
| Hosts              | 357      | 357      | 357      | 357       | 357      | 357      | OK    |
| Services           | 49       | 49       | 49       | 49        | 49       | 49       | OK    |
| User Groups        | 55       | 55       | 55       | 55        | 55       | 55       | OK    |
| Host Groups        | 29       | 29       | 29       | 29        | 29       | 29       | OK    |
| Netgroups          | 11       | 11       | 11       | 11        | 11       | 11       | OK    |
| HBAC Rules         | 3        | 3        | 3        | 3         | 3        | 3        | OK    |
| SUDO Rules         | 2        | 2        | 2        | 2         | 2        | 2        | OK    |
| DNS Zones          | 114      | 114      | 114      | 114       | 114      | 114      | OK    |
| Certificates       | 0        | 0        | 0        | 0         | 0        | 0        | OK    |
| LDAP Conflicts     | 0        | 0        | 0        | 0         | 0        | 0        | OK    |
| Ghost Replicas     | 0        | 0        | 0        | 0         | 0        | 0        | OK    |
| Anonymous BIND     | ON       | ON       | ON       | ON        | ON       | ON       | OK    |
| Microsoft ADTrust  | False    | False    | False    | False     | False    | False    | OK    |
| Replication Status | ipa03 0  | ipa03 0  | ipa04 0  | ipa03 0   | ipa03 0  | ipa04 0  | OK    |
|                    | ipa04 0  | ipa04 0  | ipa05 0  | ipa01 0   | ipa01 0  |          |       |
|                    | ipa05 0  | ipa05 0  | ipa01 0  | ipa02 0   | ipa02 0  |          |       |
|                    | ipa02 0  | ipa01 0  | ipa02 0  | ipa06 0   |          |          |       |
+--------------------+----------+----------+----------+-----------+----------+----------+-------+
```

All of these things should be exported to a prometheus exporter
