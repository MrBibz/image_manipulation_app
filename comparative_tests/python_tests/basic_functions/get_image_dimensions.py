from PIL import Image

def get_image_dimensions(img):
    # Get the dimensions of the image and return
    width, height = img.size
    return width, height