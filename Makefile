plan:
	go run main.go && cd .infra && terraform init && terraform plan

apply:
	cd .infra && terraform apply