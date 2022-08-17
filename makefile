mockgen:
	cd ./internal && for file in `egrep -rnl "type.*?interface" ./domain | grep -v "_test" `; do \
		echo $$file ; \
		cd .. && mockgen -destination="./internal/infra/mock/$$file" -source="./internal/$$file" && cd ./internal ; \
	done