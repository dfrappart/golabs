wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
#export PATH=$PATH:/usr/local/go/bin
#echo $PATH >> /etc/environment
#echo $PATH >> /etc/profile