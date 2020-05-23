FROM python:3.8-slim
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY ./cmd/image/analyze.py .
COPY ./cmd/image/image_pb2.py .
COPY ./cmd/image/image_pb2_grpc.py .
COPY ./cmd/image/data.csv .

RUN python -m pip install --upgrade pip
RUN python -m pip install grpcio
RUN pip install torch==1.5.0+cpu torchvision==0.6.0+cpu -f https://download.pytorch.org/whl/torch_stable.html

CMD python analyze.py