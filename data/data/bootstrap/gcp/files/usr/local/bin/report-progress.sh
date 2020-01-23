#!/usr/bin/env bash

KUBECONFIG="${1}"

wait_for_existance() {
	while [ ! -e "${1}" ]
	do
		sleep 5
	done
}

echo "Waiting for cvo-bootstrap to complete..."
wait_for_existance /opt/openshift/cvo-bootstrap.done

## remove the routes setup so that we can open up the blackhole
systemctl stop gcp-routes.service

echo "Waiting for bootstrap to complete..."
wait_for_existance /opt/openshift/.bootstrap.done

echo "Reporting install progress..."
while ! oc --config="$KUBECONFIG" create -f - <<-EOF
	apiVersion: v1
	kind: ConfigMap
	metadata:
	  name: bootstrap
	  namespace: kube-system
	data:
	  status: complete
EOF
do
	sleep 5
done
