import os
import time
from PIL import Image, ImageOps


# Function to ensure the output directory exists
def ensure_output_dir(output_path):
    directory = os.path.dirname(output_path)
    if not os.path.exists(directory):
        os.makedirs(directory)


# Function to invert colors of an image
def invert_colors(img_path, output_path):
    start = time.time()

    # Ensure the output directory exists
    ensure_output_dir(output_path)

    # Open the image
    img = Image.open(img_path)

    # Invert the colors
    inverted_img = ImageOps.invert(img.convert('RGB'))

    # Save the inverted image as a JPEG file
    inverted_img.save(output_path, 'JPEG')

    print(f"Inverted color image saved to: {output_path}")
    print(f"Execution time: {time.time() - start:.4f} seconds")


# Main function
def main():
    image_path = "../../images/bread.jpg"
    output_path = "../outputs/test_write_image_output.jpg"

    # Invert colors of the image and save it
    invert_colors(image_path, output_path)


if __name__ == "__main__":
    main()
