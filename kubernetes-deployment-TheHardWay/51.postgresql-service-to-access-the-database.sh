kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: postgresql
  namespace: postgresql
  labels:
    app: postgres-persistent
spec:
  type: ClusterIP
  ports:
  - port: 5432
    targetPort: 5432
    name: postgresql
  selector:
    app: postgres-persistent
EOF
