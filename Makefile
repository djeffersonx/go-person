build-image:
	docker build -t idws/go-person:1.0 .

kube-deploy:
	minikube image load idws/go-person:1.0
	kubectl apply -f kubernets/deployment.yml
	kubectl apply -f kubernets/service.yml
	kubectl apply -f kubernets/ingress.yml
	kubectl apply -f kubernets/configmap.yml

kube-destroy:
	kubectl delete deploy person-service
	kubectl delete service person-service
	kubectl delete configmap person-service
	kubectl delete ingress person-service-ingress

re-run: kube-destroy build-container-image kube-deploy