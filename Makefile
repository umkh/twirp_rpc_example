generate:
	cd api; buf generate; cd ..

clean:
	rm -rf ./pkg/genproto; rm -rf ./pkg/swaggers