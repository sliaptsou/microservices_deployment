.PHONY: vm_mysql vm_nginx
vm_nginx:
	ansible-playbook playbook-ws.yaml --limit web_servers --flush-cache

vm_mysql:
	ansible-playbook playbook-db.yaml --limit databases --flush-cache