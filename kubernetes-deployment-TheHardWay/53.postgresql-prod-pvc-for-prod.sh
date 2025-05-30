kubectl apply -f - <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgres-production-pvc
  namespace: postgresql
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: postgresql-block
  resources:
    requests:
      storage: 20Gi
EOF
