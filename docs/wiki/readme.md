
#
- docker pull jgraph/drawio
- docker run -itd  --name="drawio" -p 9090:8080 -p 8443:8443 jgraph/drawio
- http://localhost:9090/?offline=1&https=0