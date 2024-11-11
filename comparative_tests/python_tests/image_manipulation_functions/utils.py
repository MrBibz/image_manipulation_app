import os
from PIL import Image


def ensure_output_dir(outputs_path):
    output_dir = os.path.dirname(outputs_path)
    if not os.path.exists(output_dir):
        try:
            os.makedirs(output_dir, exist_ok=True)
        except OSError as e:
            raise Exception(f"Error creating output directory: {e}")


def create_output_file(outputs_path):
    try:
        output_file = open(outputs_path, 'wb')
        return output_file
    except IOError as e:
        raise Exception(f"Error creating output file: {e}")


def save_image_as_jpeg(output_file, img):
    try:
        img.save(output_file, format='JPEG')
    except Exception as e:
        raise Exception(f"Error encoding output file: {e}")
    finally:
        output_file.close()