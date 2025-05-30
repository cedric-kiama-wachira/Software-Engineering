kubectl patch statefulset postgresql -n postgresql --type=json -p='[
  {
    "op": "add", 
    "path": "/spec/template/spec/initContainers/-", 
    "value": {
      "name": "init-extensions",
      "image": "cedrickiama/pg17-4-timescaledb:2.19.3-slim",
      "command": ["/bin/bash", "/scripts/init-extensions.sh"],
      "volumeMounts": [
        {
          "name": "init-extensions",
          "mountPath": "/scripts"
        }
      ]
    }
  },
  {
    "op": "add", 
    "path": "/spec/template/spec/volumes/-", 
    "value": {
      "name": "init-extensions",
      "configMap": {
        "name": "postgres-init-extensions",
        "defaultMode": 0755
      }
    }
  }
]'
