kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-custom
  namespace: postgresql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres-custom
  template:
    metadata:
      labels:
        app: postgres-custom
    spec:
      securityContext:
        fsGroup: 999
      containers:
      - name: postgres
        image: cedrickiama/pg17-4-timescaledb:2.19.3-slim
        env:
        - name: POSTGRES_DB
          valueFrom:
            secretKeyRef:
              name: postgres-credentials
              key: POSTGRES_DB
        - name: POSTGRES_USER
          valueFrom:
            secretKeyRef:
              name: postgres-credentials
              key: POSTGRES_USER
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: postgres-credentials
              key: POSTGRES_PASSWORD
        ports:
        - containerPort: 5432
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql
      volumes:
      - name: data
        emptyDir: {}
EOF
