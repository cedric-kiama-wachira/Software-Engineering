# Delete the current deployment
kubectl delete deployment -n postgresql postgresql-custom
kubectl delete statefulset postgresql -n postgresql
# Create a new deployment
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgresql-custom
  namespace: postgresql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgresql-custom
  template:
    metadata:
      labels:
        app: postgresql-custom
    spec:
      containers:
        - name: postgresql
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
            - name: PGDATA
              value: /var/lib/postgresql/17/main
          ports:
            - containerPort: 5432
          volumeMounts:
            - name: data-volume
              mountPath: /var/lib/postgresql/17/main
            - name: pgbackrest-volume
              mountPath: /etc/pgbackrest
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
      volumes:
        - name: data-volume
          emptyDir: {}
        - name: pgbackrest-volume
          emptyDir: {}
EOF
