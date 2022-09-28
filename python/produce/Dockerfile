# syntax=docker/dockerfile:1

FROM python:3.10-slim-buster

WORKDIR /app

COPY requirements.txt ./
RUN pip3 install -r requirements.txt

COPY ./ ./

ENTRYPOINT [ "python3", "produce/main.py" ]
