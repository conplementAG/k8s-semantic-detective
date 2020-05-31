package externalprobe

const TestConfigMap string = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: management-k8s-probe-volatile-test-object
  namespace: {{ namespace }}
  labels: 
    name: management-k8s-probe
    timestamp: "{{ timestamp }}" # the part that always changes
`
