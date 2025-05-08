cat <<EOF | kubectl apply -f -
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
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "9187"
    spec:
      serviceAccountName: postgresql-sa
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/hostname
                operator: In
                values:
                - wkone.veryeasy.dev
                - wktwo.veryeasy.dev
                - wkthree.veryeasy.dev
                - wkfour.veryeasy.dev
      securityContext:
        fsGroup: 999  # postgres group id
      initContainers:
        - name: init-permissions
          image: busybox
          command: ["sh", "-c", "chown -R 999:999 /var/lib/postgresql/17/main /var/lib/pgbackrest"]
          volumeMounts:
            - name: postgresql-data
              mountPath: /var/lib/postgresql/17/main
            - name: pgbackrest-data
              mountPath: /var/lib/pgbackrest
      containers:
        - name: postgresql
          image: cedrickiama/pg17-4-timescaledb:2.19.3-slim
          imagePullPolicy: IfNotPresent
          resources:
            requests:
              cpu: "2"
              memory: "4Gi"
            limits:
              cpu: "6"
              memory: "10Gi"
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
              name: postgresql
          volumeMounts:
            - name: postgresql-data
              mountPath: /var/lib/postgresql/17/main
            - name: pgbackrest-data
              mountPath: /var/lib/pgbackrest
            - name: pgbackrest-config
              mountPath: /etc/pgbackrest/pgbackrest.conf
              subPath: pgbackrest.conf
            - name: postgres-initdb
              mountPath: /docker-entrypoint-initdb.d/init-extensions.sh
              subPath: init-extensions.sh
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
          resources:
            requests:
              cpu: "200m"
              memory: "256Mi"
            limits:
              cpu: "500m"
              memory: "512Mi"
          env:
            - name: DATA_SOURCE_NAME
              value: "postgresql://\$(POSTGRES_USER):\$(POSTGRES_PASSWORD)@localhost:5432/\$(POSTGRES_DB)?sslmode=disable"
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
            - name: POSTGRES_DB
              valueFrom:
                secretKeyRef:
                  name: postgres-credentials
                  key: POSTGRES_DB
          ports:
            - containerPort: 9187
              name: metrics
      volumes:
        - name: pgbackrest-config
          configMap:
            name: pgbackrest-config
        - name: postgres-initdb
          configMap:
            name: postgres-initdb
            defaultMode: 0755
  volumeClaimTemplates:
    - metadata:
        name: postgresql-data
      spec:
        accessModes: [ "ReadWriteOnce" ]
        storageClassName: rook-ceph-block
        resources:
          requests:
            storage: 50Gi
    - metadata:
        name: pgbackrest-data
      spec:
        accessModes: [ "ReadWriteOnce" ]
        storageClassName: rook-ceph-block
        resources:
          requests:
            storage: 25Gi
EOF
