from PIL import Image
import numpy as np
from . import utils
import time


def truncate(value):
    return max(0, min(int(value), 255))

def contrast_filter(img, contrast_factor, outputs_path):
    start_time = time.time()

    # Ensure the output directory exists
    utils.ensure_output_dir(outputs_path)

    # Convert image to numpy array
    img_array = np.array(img.convert("RGBA"))

    # Calculating the contrast factor
    contrast = (259 * (contrast_factor + 255)) / (255 * (259 - contrast_factor))

    # Apply the contrast adjustment
    contrast_adjusted_array = img_array.copy()
    for y in range(img_array.shape[0]):
        for x in range(img_array.shape[1]):
            # Adjusting the contrast for each channel
            r = truncate(contrast * (img_array[y, x, 0] - 128) + 128)
            g = truncate(contrast * (img_array[y, x, 1] - 128) + 128)
            b = truncate(contrast * (img_array[y, x, 2] - 128) + 128)

            # Replacing the pixel values
            contrast_adjusted_array[y, x] = [r, g, b, img_array[y, x, 3]]

    # Conversion of the numpy array to an image
    contrast_adjusted_img = Image.fromarray(contrast_adjusted_array, "RGBA")

    # Save the contrast adjusted image
    output_file = utils.create_output_file(outputs_path)
    utils.save_image_as_jpeg(output_file, contrast_adjusted_img.convert("RGB"))

    print("Contrast adjusted image saved to:", output_file)
    print("Execution time:", time.time() - start_time, "seconds")
