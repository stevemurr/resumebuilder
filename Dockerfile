FROM ivanpondal/alpine-latex
RUN apk update --no-cache && apk add python-dev py-pip
RUN pip install tex
ADD convert.py convert.py
ENTRYPOINT ["python", "convert.py"]
