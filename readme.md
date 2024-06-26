To run the docker container with the image, use the following command:

```bash
docker build -t backend-image -f Dockerfile .
docker run -p 8080:8080 backend-image
```
