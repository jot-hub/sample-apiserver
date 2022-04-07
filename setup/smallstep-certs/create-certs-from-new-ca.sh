#init a ca with root and intermediate
#$ step ca init --helm > values.yaml
#save certs locally
#$ yq .inject.certificates.root_ca values.yaml > root_ca
#$ yq .inject.certificates.intermediate_ca values.yaml > intermediate_ca
#$ yq .inject.secrets.x509.intermediate_ca_key values.yaml > intermediate_ca_key

#prepare a cert chain concatenating intermediateca and rootca in that order
#$ cat intermediate_ca > ca-cert-chain
#$ cat root_ca >> ca-cert-chain  

# supply this chain file to --client-ca-file
#$ CACERTCHAIN=$(base64 ca-cert-chain) yq -i '.data.cert-chain=strenv(CACERTCHAIN)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/base/templates/secret.yaml

#create csr for apiserver and sign it with intermediate to get cert file
#$ step certificate create --csr --insecure --no-password --san 127.0.0.1 --san 0.0.0.0 --san localhost "standalone-kube-apiserver.wardle.svc.cluster.local" standalone-kube-apiserver.csr standalone-kube-apiserver.key
#$ step certificate sign --not-after 168h standalone-kube-apiserver.csr intermediate_ca.crt intermediate_ca_key.key > standalone-kube-apiserver.crt
#  - to be supplied to tls-cert-file and tls-private-key-file
#  - needs complete cert chain
#$ cat standalone-kube-apiserver.crt > standalone-kube-apiserver-cert-chain
#$ cat intermediate_ca >> standalone-kube-apiserver-cert-chain
#$ cat root_ca >> standalone-kube-apiserver-cert-chain
#$ KUBEAPISERVERCERT=$(base64 standalone-kube-apiserver-cert-chain) yq -i 'select(documentIndex == 0).data.tls-cert=strenv(KUBEAPISERVERCERT)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/templates/secret.yaml
#$ KUBEAPISERVERKEY=$(base64 standalone-kube-apiserver.key) yq -i 'select(documentIndex == 0).data.tls-key=strenv(KUBEAPISERVERKEY)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/templates/secret.yaml

#create csr for client and sign it with intermediate to get cert file
#$ step certificate create --csr --insecure --no-password --san curl --san kubeconfig "client" client-for-kube-apiserver.csr client-for-kube-apiserver.key
#$ step certificate sign --not-after 168h client-for-kube-apiserver.csr intermediate_ca.crt intermediate_ca_key.key > client-for-kube-apiserver.crt
#  - to be supplied in curl & kubeconfig together with unencrypted key
#for within the cluster
#$ CLIENTFORKUBEAPISERVERKEY=$(base64 client-for-kube-apiserver.key) yq -i '.users[0].user.client-key-data=strenv(CLIENTFORKUBEAPISERVERKEY)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/misc/constructed-kubeconfig-incluster
#$ CLIENTFORKUBEAPISERVERCERT=$(base64 client-for-kube-apiserver.crt) yq -i '.users[0].user.client-certificate-data=strenv(CLIENTFORKUBEAPISERVERCERT)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/misc/constructed-kubeconfig-incluster
#from CLI
#$ CLIENTFORKUBEAPISERVERKEY=$(base64 client-for-kube-apiserver.key) yq -i '.users[0].user.client-key-data=strenv(CLIENTFORKUBEAPISERVERKEY)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/misc/constructed-kubeconfig
#$ CLIENTFORKUBEAPISERVERCERT=$(base64 client-for-kube-apiserver.crt) yq -i '.users[0].user.client-certificate-data=strenv(CLIENTFORKUBEAPISERVERCERT)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/standalone-kube-apiserver/misc/constructed-kubeconfig

#create csr for kine and sign it with intermediate to get cert file
#$ step certificate create --csr --insecure --no-password --san 127.0.0.1 --san 0.0.0.0 --san localhost "kine-with-mysql-service.wardle.svc.cluster.local" kube-kine.csr kube-kine.key
#$ step certificate sign --not-after 168h --template kube-kine-cert-template.tpl kube-kine.csr intermediate_ca.crt intermediate_ca_key.key > kube-kine.crt
#  - to be supplied in kine together with unencrypted key
#$ KUBEKINECERT=$(base64 kube-kine.crt) yq -i 'select(documentIndex == 0).data.tls-cert=strenv(KUBEKINECERT)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/kine-with-mysql/templates/secret.yaml
#$ KUBEKINEKEY=$(base64 kube-kine.key) yq -i 'select(documentIndex == 0).data.tls-key=strenv(KUBEKINEKEY)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/kine-with-mysql/templates/secret.yaml

#create csr for mysql and sign it with intermediate to get cert file
#$ step certificate create --csr --insecure --no-password --san 127.0.0.1 --san 0.0.0.0 --san localhost "127.0.0.1" kube-mysql.csr kube-mysql.key
#$ step certificate sign --not-after 168h kube-mysql.csr intermediate_ca.crt intermediate_ca_key.key > kube-mysql.crt
#  - to be supplied in mysql together with unencrypted key
#$ KUBEMYSQLCERT=$(base64 kube-mysql.crt) yq -i 'select(documentIndex == 1).data.tls-cert=strenv(KUBEMYSQLCERT)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/kine-with-mysql/templates/secret.yaml
#$ KUBEMYSQLKEY=$(base64 kube-mysql.key) yq -i 'select(documentIndex == 1).data.tls-key=strenv(KUBEMYSQLKEY)' ~/github.com/jot-hub/k8s-experiments/kafee/setup/kine-with-mysql/templates/secret.yaml