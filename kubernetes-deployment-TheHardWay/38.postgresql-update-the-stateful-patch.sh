kubectl patch statefulset postgresql -n postgresql --type=json -p='[
  {
    "op": "add", 
    "path": "/spec/template/spec/containers/0/volumeMounts/-", 
    "value": {
      "name": "init-script", 
      "mountPath": "/docker-entrypoint-initdb.d/init-extensions.sql",
      "subPath": "init-extensions.sql"
    }
  },
  {
    "op": "add", 
    "path": "/spec/template/spec/containers/0/volumeMounts/-", 
    "value": {
      "name": "init-runner", 
      "mountPath": "/docker-entrypoint-initdb.d/init-db.sh",
      "subPath": "init-db.sh"
    }
  },
  {
    "op": "add", 
    "path": "/spec/template/spec/volumes/-", 
    "value": {
      "name": "init-script", 
      "configMap": {
        "name": "postgres-init-script"
      }
    }
  },
  {
    "op": "add", 
    "path": "/spec/template/spec/volumes/-", 
    "value": {
      "name": "init-runner", 
      "configMap": {
        "name": "postgres-init-runner",
        "defaultMode": 0755
      }
    }
  }
]'
