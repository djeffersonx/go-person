build-container-image:
	docker build -t localhost:5001/go-person:1.0 .
	#docker image ls | grep localhost:5001/go-person:1.0
	docker tag 9c99de864904 localhost:5001/go-person:1.0
	docker image push localhost:5001/go-person:1.0

kube-deploy:
	kubectl apply -f kubernets/deployment.yml
	kubectl apply -f kubernets/service.yml
	kubectl apply -f kubernets/configmap.yml

kube-destroy:
	kubectl delete deploy person-service-deployment
	kubectl delete service person-service
	kubectl delete configmap person-service-config

re-run: kube-destroy build-container-image kube-deploy