kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: postgresql-production
  namespace: postgresql
  labels:
    app: postgres-production
spec:
  type: ClusterIP
  ports:
  - port: 5432
    targetPort: postgresql
    name: postgresql
  - port: 9187
    targetPort: metrics
    name: metrics
  selector:
    app: postgres-production
EOF
