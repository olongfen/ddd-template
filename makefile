mockgen:
	cd ./internal && for file in `egrep -rnl "type.*?interface" ./ | grep -v "_test" `; do \
		echo $$file ; \
		cd .. && mockgen -destination="./internal/adapters/mock/$$file" -source="./internal/$$file" && cd ./internal ; \
	done

gqlgen:
	 gqlgen generate  --config gqlgen.yml