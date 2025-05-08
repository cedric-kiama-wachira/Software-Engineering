# Delete the current StatefulSet but keep the PVCs
kubectl delete statefulset postgresql -n postgresql --cascade=orphan

# Create a new StatefulSet with a simpler configuration
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
          mkdir -p /var/lib/postgresql/17/main
          mkdir -p /var/lib/pgbackrest
          chown -R 999:999 /var/lib/postgresql
          chown -R 999:999 /var/lib/pgbackrest
          chmod 700 /var/lib/postgresql/17/main
          echo "Initialization complete"
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql
        - name: backrest
          mountPath: /var/lib/pgbackrest
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
      volumes:
      - name: backrest
        persistentVolumeClaim:
          claimName: data-postgresql-0
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: postgresql-block
      resources:
        requests:
          storage: 50Gi
EOF
