import tensorflow as tf
from tensorflow.keras.applications.resnet50 import ResNet50
from tensorflow.keras.preprocessing.image import load_img, img_to_array
from tensorflow.keras.applications.resnet50 import preprocess_input
import urllib.request
from PIL import Image
import numpy as np
import sys
import os


class Model:


    def __init__(self):


        # Load the pretrained ResNet50 model
        self.model = ResNet50(weights='imagenet')


    def predict(self,img_url: str) -> int:
        
        extension = img_url[-3:]
        img_path = "img." + extension
        urllib.request.urlretrieve(img_url,img_path)
        img = tf.keras.utils.load_img(img_path, target_size=(224,224))
        os.remove(img_path)
        #image preprocessing
        img = img_to_array(img)
        img = preprocess_input(img)

        # Make a prediction with the model
        prediction = self.model.predict(img.reshape(1, 224, 224, 3))

        # Convert the prediction to a binary classification
        if prediction[0][0] > prediction[0][1]:
            classification = '0'  # First class
        else:
            classification = '1'  # Second class

        print(f'The image is classified as: {classification}')
        print(classification)

        return classification
