import sys
import torch
from torchvision import transforms, models

model = models.resnet50(pretrained=True)
model.eval()

args = sys.argv

if len(args) < 2:
    sys.exit(0)

from PIL import Image
from torchvision import transforms

input_image = Image.open(args[1])
preprocess = transforms.Compose([
    transforms.Resize(256),
    transforms.CenterCrop(224),
    transforms.ToTensor(),
    transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
])
input_tensor = preprocess(input_image)
input_batch = input_tensor.unsqueeze(0)  # create a mini-batch as expected by the model

# move the input and model to GPU for speed if available

with torch.no_grad():
    output = model(input_batch)
# Tensor of shape 1000, with confidence scores over Imagenet's 1000 classes
a = torch.nn.functional.softmax(output[0], dim=0).tolist()
print(a.index(max(a)), end = '')