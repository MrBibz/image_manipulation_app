from PIL import Image
import numpy as np
from . import utils
import time


def blur_filter(img, intensity, outputs_path):
    start_time = time.time()

    # Ensure the output directory exists
    utils.ensure_output_dir(outputs_path)

    # Get the image dimensions
    width, height = img.size

    # Convert image to numpy array
    img_array = np.array(img.convert("RGBA"))

    # Create the blur kernel
    kernel_size = 2 * intensity + 1
    kernel = np.ones((kernel_size, kernel_size)) / (kernel_size * kernel_size)

    # Create the output file
    blurred_array = np.zeros_like(img_array)

    # Apply the blur filter
    for y in range(intensity, height - intensity):
        for x in range(intensity, width - intensity):
            r_sum, g_sum, b_sum = 0, 0, 0

            # Apply the kernel to the image
            for ky in range(-intensity, intensity + 1):
                for kx in range(-intensity, intensity + 1):
                    px = img_array[y + ky, x + kx]
                    weight = kernel[ky + intensity, kx + intensity]

                    r_sum += px[0] * weight
                    g_sum += px[1] * weight
                    b_sum += px[2] * weight

            # Limit the RGB values to 255
            blurred_array[y, x] = [int(r_sum), int(g_sum), int(b_sum), 255]

    # Conversion of the numpy array to an image
    blurred_img = Image.fromarray(blurred_array, "RGBA")

    # Save the blurred image
    output_file = utils.create_output_file(outputs_path)
    utils.save_image_as_jpeg(output_file, blurred_img.convert("RGB"))

    print("Blurred image saved to:", output_file)
    print("Execution time:", time.time() - start_time, "seconds")
