#init a ca with root and intermediate
#prepare a cert chain concatenating intermediateca and rootca in that order
# supply this chain file to --client-ca-file
#create csr for apiserver and sign it with intermediate to get cert file
#  - to be supplied to tls-cert-file and tls-private-key-file
#create csr for client and sign it with intermediate to get cert file
#  - to be supplied in curl & kubeconfig together with unencrypted key
#create csr for kine and sign it with intermediate to get cert file
#  - to be supplied in kine together with unencrypted key
#create csr for mysql and sign it with intermediate to get cert file
#  - to be supplied in mysql together with unencrypted key

