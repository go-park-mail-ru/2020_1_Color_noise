import grpc
from concurrent import futures
from PIL import Image
import torch
from torchvision import transforms, models
import csv
import logging
import image_pb2
import image_pb2_grpc

table = list()

with open('data.csv', newline='') as csvfile:
    spamreader = csv.reader(csvfile, delimiter=';')
    for row in spamreader:
        table.append(row)

model = models.resnet50(pretrained=True)
model.eval()

class ImageServiceServicer(image_pb2_grpc.ImageServiceServicer):
    """Provides methods that implement functionality of route guide server."""

    def Analyze(self, request, context):
        input_image = Image.open(request.Image)
        preprocess = transforms.Compose([
            transforms.Resize(256),
            transforms.CenterCrop(224),
            transforms.ToTensor(),
            transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
        ])
        input_tensor = preprocess(input_image)
        input_batch = input_tensor.unsqueeze(0)  # create a mini-batch as expected by the model

        with torch.no_grad():
            output = model(input_batch)

        a = torch.nn.functional.softmax(output[0], dim=0).tolist()

        return image_pb2.Tags(Tags=table[a.index(max(a))][:2])

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=8))
    image_pb2_grpc.add_ImageServiceServicer_to_server(
        ImageServiceServicer(), server)
    server.add_insecure_port('[::]:8000')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()