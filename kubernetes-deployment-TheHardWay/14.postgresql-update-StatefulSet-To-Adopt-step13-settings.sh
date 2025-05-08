kubectl patch statefulset -n postgresql postgresql --type=json -p='[{"op": "add", "path": "/spec/template/spec/containers/0/volumeMounts/-", "value": {"name": "postgresql-config", "mountPath": "/etc/postgresql/17/main/conf.d/custom.conf", "subPath": "postgresql.conf"}}]'

kubectl patch statefulset -n postgresql postgresql --type=json -p='[{"op": "add", "path": "/spec/template/spec/volumes/-", "value": {"name": "postgresql-config", "configMap": {"name": "postgresql-config"}}}]'
