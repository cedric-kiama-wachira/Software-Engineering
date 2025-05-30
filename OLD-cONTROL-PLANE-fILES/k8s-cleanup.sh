#!/bin/bash

set -xe

echo "[1/6] Resetting kubeadm (if initialized)"
sudo kubeadm reset -f || true

echo "[2/6] Stopping kubelet and containerd"
sudo systemctl stop kubelet || true
sudo systemctl stop containerd || true

echo "[3/6] Removing kubelet, kubeadm, kubectl"
sudo apt-mark unhold kubelet kubeadm kubectl
sudo apt-get purge -y kubelet kubeadm kubectl

echo "[4/6] Removing containerd and Docker (if installed)"
sudo apt-get purge -y containerd.io docker-ce docker-ce-cli docker-buildx-plugin docker-compose-plugin
sudo rm -rf /etc/containerd /var/lib/containerd /etc/docker /var/lib/docker

echo "[5/6] Removing CNI configs and network state"
sudo rm -rf /etc/cni /opt/cni /var/lib/cni /var/lib/kubelet /etc/kubernetes /etc/systemd/system/kubelet.service.d

echo "[6/6] Cleaning apt lists and leftover files"
sudo apt-get autoremove -y
sudo apt-get clean
sudo rm -rf ~/.kube /etc/sysctl.d/k8s.conf /etc/modules-load.d/k8s.conf /etc/apt/keyrings/kubernetes-apt-keyring.gpg
sudo cd /etc/apt/keyrings && rm -Rf docker.asc kubernetes-apt-keyring.gpg && cd /etc/apt/sources.list.d && rm -Rf docker.list kubernetes.list && apt update && apt -y upgrade
sudo ipvsadm --clear || true

echo "System cleaned. You may want to reboot before re-bootstrap."
reboot
