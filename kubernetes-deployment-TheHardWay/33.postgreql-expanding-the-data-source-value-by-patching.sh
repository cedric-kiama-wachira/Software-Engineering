kubectl patch statefulset postgresql -n postgresql --type=json -p='[
  {
    "op": "replace", 
    "path": "/spec/template/spec/containers/1/env/0", 
    "value": {
      "name": "DATA_SOURCE_NAME", 
      "value": "postgresql://postgres:$(POSTGRES_PASSWORD)@localhost:5432/postgres?sslmode=disable"
    }
  }
]'
