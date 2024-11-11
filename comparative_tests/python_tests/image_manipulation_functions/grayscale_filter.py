from PIL import Image
import numpy as np
from . import utils
import time


def grayscale_filter(img, intensity, outputs_path):
    start_time = time.time()

    # Ensure the output directory exists
    utils.ensure_output_dir(outputs_path)

    # Conversion of the image to a numpy array
    img_array = np.array(img.convert("RGBA"))

    # Apply the grayscale filter
    grayscale_array = img_array.copy()
    blend_factor = intensity / 100.0

    for y in range(img_array.shape[0]):
        for x in range(img_array.shape[1]):
            original_color = img_array[y, x]
            gray_value = int(0.3 * original_color[0] + 0.59 * original_color[1] + 0.11 * original_color[2])

            # Calculate the new color
            r = int((1 - blend_factor) * original_color[0] + blend_factor * gray_value)
            g = int((1 - blend_factor) * original_color[1] + blend_factor * gray_value)
            b = int((1 - blend_factor) * original_color[2] + blend_factor * gray_value)

            # Update the pixel value
            grayscale_array[y, x] = [r, g, b, original_color[3]]

    # Conversion of the numpy array to an image
    grayscale_img = Image.fromarray(grayscale_array, "RGBA")

    # Save the grayscale image
    output_file = utils.create_output_file(outputs_path)
    utils.save_image_as_jpeg(output_file, grayscale_img.convert("RGB"))

    print("Grayed image saved to:", output_file)
    print("Execution time:", time.time() - start_time, "seconds")
