from PIL import Image
import numpy as np
from . import utils
import time


def resize_image(img, new_width, new_height, outputs_path):
    start_time = time.time()

    # Ensure the output directory exists
    utils.ensure_output_dir(outputs_path)

    # Convert image to numpy array
    img_array = np.array(img.convert("RGBA"))
    original_height, original_width = img_array.shape[:2]

    # Create the output array
    resized_array = np.zeros((new_height, new_width, 4), dtype=np.uint8)

    # Resize the image with bilinear interpolation
    for y in range(new_height):
        for x in range(new_width):
            # Calculate the original pixel coordinates
            original_x = x * (original_width / new_width)
            original_y = y * (original_height / new_height)

            x1, y1 = int(original_x), int(original_y)
            x2, y2 = min(x1 + 1, original_width - 1), min(y1 + 1, original_height - 1)

            # Weight of the pixels
            wx, wy = original_x - x1, original_y - y1

            # Colors of the neighboring pixels
            c11 = img_array[y1, x1]
            c12 = img_array[y2, x1]
            c21 = img_array[y1, x2]
            c22 = img_array[y2, x2]

            # Bilinear interpolation
            r = int((1 - wx) * (1 - wy) * c11[0] + wx * (1 - wy) * c21[0] + (1 - wx) * wy * c12[0] + wx * wy * c22[0])
            g = int((1 - wx) * (1 - wy) * c11[1] + wx * (1 - wy) * c21[1] + (1 - wx) * wy * c12[1] + wx * wy * c22[1])
            b = int((1 - wx) * (1 - wy) * c11[2] + wx * (1 - wy) * c21[2] + (1 - wx) * wy * c12[2] + wx * wy * c22[2])
            a = int((1 - wx) * (1 - wy) * c11[3] + wx * (1 - wy) * c21[3] + (1 - wx) * wy * c12[3] + wx * wy * c22[3])

            # Apply the new pixel value
            resized_array[y, x] = [r, g, b, a]

    # Convert the numpy array to an image
    resized_img = Image.fromarray(resized_array, "RGBA")

    # Save the resized image
    output_file = utils.create_output_file(outputs_path)
    utils.save_image_as_jpeg(output_file, resized_img.convert("RGB"))

    print("Resized image saved to:", output_file)
    print("Execution time:", time.time() - start_time, "seconds")
