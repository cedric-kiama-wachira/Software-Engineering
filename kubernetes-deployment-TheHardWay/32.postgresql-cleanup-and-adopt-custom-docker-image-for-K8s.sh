# Delete all PostgreSQL-related resources except PVCs
kubectl delete pod -n postgresql postgresql-debug
kubectl delete pod -n postgresql postgresql-0
kubectl delete deployment -n postgresql --all
kubectl delete statefulset -n postgresql --all --cascade=orphan

kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: postgresql
  namespace: postgresql
spec:
  serviceName: postgresql
  replicas: 1
  selector:
    matchLabels:
      app: postgresql
  template:
    metadata:
      labels:
        app: postgresql
    spec:
      # This is crucial - the initContainer runs as root to set permissions
      # but the main container will run as postgres (user 999)
      securityContext:
        fsGroup: 999
      initContainers:
      - name: init-permissions
        image: busybox
        command:
        - /bin/sh
        - -c
        - |
          echo "Preparing directories..."
          # Create directories owned by postgres user
          mkdir -p /var/lib/postgresql/17/main
          mkdir -p /var/lib/pgbackrest
          mkdir -p /var/log/pgbackrest
          # Set proper ownership
          chown -R 999:999 /var/lib/postgresql
          chown -R 999:999 /var/lib/pgbackrest
          chown -R 999:999 /var/log/pgbackrest
          chmod 700 /var/lib/postgresql/17/main
          ls -la /var/lib/postgresql/17
          echo "Initialization complete"
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql
        - name: backrest
          mountPath: /var/lib/pgbackrest
        - name: logs
          mountPath: /var/log/pgbackrest
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
        ports:
        - containerPort: 5432
          name: postgresql
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql
        - name: backrest
          mountPath: /var/lib/pgbackrest
        - name: backrest-config
          mountPath: /etc/pgbackrest
        - name: logs
          mountPath: /var/log/pgbackrest
        - name: extensions
          mountPath: /docker-entrypoint-initdb.d/init-extensions.sh
          subPath: init-extensions.sh
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
          value: "postgresql://postgres:\$(POSTGRES_PASSWORD)@localhost:5432/\$(POSTGRES_DB)?sslmode=disable"
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
      - name: backrest-config
        emptyDir: {}
      - name: extensions
        configMap:
          name: postgres-initdb
          defaultMode: 0755
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: postgresql-block
      resources:
        requests:
          storage: 50Gi
  - metadata:
      name: backrest
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: postgresql-block
      resources:
        requests:
          storage: 20Gi
  - metadata:
      name: logs
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: postgresql-block
      resources:
        requests:
          storage: 5Gi
EOF
