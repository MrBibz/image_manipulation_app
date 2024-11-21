from PIL import Image

def read_image(file_path):
    # Open the image file and check for errors
    try:
        img = Image.open(file_path)
    except IOError as e:
        print(f"Error opening file: {e}")
        return None

    # Decode the image file and check for errors
    try:
        img.load()  # Ensure the image data is read
    except IOError as e:
        print(f"Error decoding file: {e}")
        return None

    return img