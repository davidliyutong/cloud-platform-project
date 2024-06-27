GIT_VERSION := $(shell git describe --tags --abbrev=0 --always)
AUTHOR="davidliyutong"

task.helm.debug:
	helm install --namespace=clpl cloud-platform ./helm --dry-run --debug
task.helm.install:
	helm install --namespace=clpl cloud-platform ./helm 

task.helm.uninstall:
	helm uninstall --namespace=clpl cloud-platform 