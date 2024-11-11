from PIL import Image


def get_image_dimensions(image_path):
    img = Image.open(image_path)

    # // Get the dimensions of the image and return
    width, height = img.size

    return width, height
