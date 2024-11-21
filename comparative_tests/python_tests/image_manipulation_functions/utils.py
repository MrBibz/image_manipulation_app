import os
from PIL import Image

def ensure_output_dir(outputs_path):
    """
    Ensure the output directory exists.
    """
    output_dir = os.path.dirname(outputs_path)
    if not os.path.exists(output_dir):
        try:
            os.makedirs(output_dir)
        except OSError as e:
            raise RuntimeError(f"Error creating output directory: {e}") from e

def create_output_file(outputs_path):
    """
    Create the output file and return the file path.
    """
    try:
        # Open the file in write mode; 'wb' ensures binary writing
        return open(outputs_path, 'wb')
    except OSError as e:
        raise RuntimeError(f"Error creating output file: {e}") from e

def save_image_as_jpeg(output_file, img):
    """
    Encode the image as JPEG and save it to the output file.
    """
    try:
        img.save(output_file, format="JPEG")
    except IOError as e:
        raise RuntimeError(f"Error encoding output file: {e}") from e
    finally:
        # Close the output file
        output_file.close()