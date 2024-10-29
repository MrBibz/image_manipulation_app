import imageio.v3 as iio
import matplotlib.pyplot as plt
import skimage as ski
import os
import numpy as np

plt.ion()

image_path = "../../images/mike_back.bmp"
image = iio.imread(image_path)

fig, ax = plt.subplots()
ax.imshow(image)
plt.show()

sigma = 10.0
blurred = ski.filters.gaussian(image, sigma=(sigma, sigma), truncate=3.5, channel_axis=-1)

blurred_uint8 = (blurred * 255).astype(np.uint8)

blurred_image_path = os.path.join(os.path.dirname(image_path), "mike_back_blurred.bmp")
iio.imwrite(blurred_image_path, blurred_uint8)