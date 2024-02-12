# Oracle Monitor

Oracle Monitor is monitoring program for operator of program which bring off-chain data to on-chain like Oracle.

## Monitoring Support Chains

- Umee
  - Oracle

## Quick Guide

1. Build

```bash
go build
```

2. Configure config.toml file

```bash
# You can get a example of config.toml file by below command
cp config.toml.example config.toml
```

3. Execute

```bash
./oraclemon -config config.toml
```
