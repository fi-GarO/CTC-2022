# Verze b pro úlohu 05

./ctcgrpc client -e localhost:8080 del -> Správně vypisuje error, že nedostal dostatečný počet argumentů
./ctcgrpc client -e localhost:8080 del a -> Vrací rpc error: code = Unimplemented desc = unknown method Delete for service core.Api. 

Myslím si, že problém nastává pravděpodobně tedy proto, že importuji vaše package, které funkci Delete neobsahují. 
Verze, ve které jsem se snažil dosadit vlastní importy je v adresáři /05/
