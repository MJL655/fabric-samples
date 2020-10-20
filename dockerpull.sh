# Clean all docker images
docker rmi -f `docker images -q`
# pull all GM fabric images
docker pull studyzy/fabric-ccenv
docker tag studyzy/fabric-ccenv hyperledger/fabric-ccenv
docker tag hyperledger/fabric-ccenv hyperledger/fabric-ccenv:1.4
docker tag hyperledger/fabric-ccenv hyperledger/fabric-ccenv:1.4.8

docker pull studyzy/fabric-ca
docker tag studyzy/fabric-ca hyperledger/fabric-ca

docker pull studyzy/fabric-tools
docker tag studyzy/fabric-tools hyperledger/fabric-tools

docker pull studyzy/fabric-peer
docker tag studyzy/fabric-peer hyperledger/fabric-peer

docker pull studyzy/fabric-orderer
docker tag studyzy/fabric-orderer hyperledger/fabric-orderer
