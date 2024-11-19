# DeBeAndo Agent for ProxySQL

Database monitoring tool designed for small environments, adapted for Kubernetes and send metrics to InfluxDB.

## Image Description

This image is maintained by DeBeAndo and will be updated regularly on best-effort basis. The image is based on Alpine Linux and only contains the build result of this repository.

## Run

To run container:

```bash
docker run \
	--name debeando-agent-proxysql \
	--env DEBUG=true \
	--env INTERVAL=10 \
	--env INFLUXDB_HOST="http://com-env-influxdb-observability-node01.aws.com" \
	--env INFLUXDB_TOKEN="abc123cde456==" \
	--env PROXYSQL_HOST="com-env-proxysql-stack-node01.aws.com" \
	--env PROXYSQL_USER="radmin" \
	--env PROXYSQL_PASSWORD="<radmin-pass>" \
	--env SERVER="com-env-proxysql-stack-node01" \
	debeando/agent-proxysql
```

## ProxySQL Config

Allow and configure remote access for a `radmin` user to agent. Remember, `radmin` has administrator privileges.

```sql
SET admin-admin_credentials = "admin:<admin-pass>;radmin:<radmin-pass>";
LOAD ADMIN VARIABLES TO RUNTIME;
SAVE ADMIN VARIABLES TO DISK;
```

Please, change default password.
