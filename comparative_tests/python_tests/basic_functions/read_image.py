from PIL import Image


def read_image(file_path):
    try:
        # Open the image file and check for errors
        img = Image.open(file_path)

        if img.format != 'JPEG':
            raise ValueError("The provided image is not a JPEG file")

        # Return the image
        return img
    except Exception as e:
        # Error handling
        print("Error during the opening or decoding of the file : ", e)
        return None
