kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-production
  namespace: postgresql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres-production
  template:
    metadata:
      labels:
        app: postgres-production
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "9187"
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
          name: postgresql
        volumeMounts:
        - name: postgres-data
          mountPath: /var/lib/postgresql
        resources:
          requests:
            cpu: "1"
            memory: "2Gi"
          limits:
            cpu: "4"
            memory: "8Gi"
        livenessProbe:
          exec:
            command:
            - pg_isready
            - -U
            - postgres
          initialDelaySeconds: 60
          periodSeconds: 20
          timeoutSeconds: 5
          failureThreshold: 6
        readinessProbe:
          exec:
            command:
            - pg_isready
            - -U
            - postgres
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
      - name: metrics
        image: prometheuscommunity/postgres-exporter:latest
        env:
        - name: DATA_SOURCE_NAME
          value: "postgresql://postgres:\${POSTGRES_PASSWORD}@localhost:5432/\${POSTGRES_DB}?sslmode=disable"
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: postgres-credentials
              key: POSTGRES_PASSWORD
        - name: POSTGRES_DB
          valueFrom:
            secretKeyRef:
              name: postgres-credentials
              key: POSTGRES_DB
        ports:
        - containerPort: 9187
          name: metrics
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "200m"
            memory: "256Mi"
      volumes:
      - name: postgres-data
        persistentVolumeClaim:
          claimName: postgres-data-pvc
EOF
