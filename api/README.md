## API

Currently the service supports one request.
You can just ask service about status.

```bash
curl host:8080/api/status | jq .
{
  "version": "v0.0.1",
  "version_api": "0.1",
  "build": "master-ba034efc-2020.05.24-01:24:14",
  "uptime": "20s"
}
```