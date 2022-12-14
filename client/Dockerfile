FROM python:3-alpine
COPY index.html /data/
COPY dist/ /data/dist/
EXPOSE 8082
WORKDIR /data
CMD ["python", "-m" , "http.server", "8082"]