from PIL import Image
import os
import time

def ensure_output_dir(output_path):
    # Ensure the output directory exists
    directory = os.path.dirname(output_path)
    if not os.path.exists(directory):
        os.makedirs(directory)

def read_image(image_path):
    # Open the image from the given path
    return Image.open(image_path)

def save_image_as_jpeg(output_path, image):
    # Save the image as JPEG to the specified output path
    image.save(output_path, format="JPEG")

def invert_colors(img, outputs_path):
    start = time.time()

    # Ensure the output directory exists
    ensure_output_dir(outputs_path)

    # Get the image dimensions
    width, height = img.size
    inverted_image = Image.new("RGBA", img.size)

    # Invert the colors
    for y in range(height):
        for x in range(width):
            # Get the pixel color
            px_color = img.getpixel((x, y))

            # Invert the color
            inverted_pixel = (
                255 - px_color[0],  # R
                255 - px_color[1],  # G
                255 - px_color[2],  # B
                px_color[3] if len(px_color) > 3 else 255  # A (if exists, else default to 255)
            )
            inverted_image.putpixel((x, y), inverted_pixel)

    # Save the inverted image as JPEG
    save_image_as_jpeg(outputs_path, inverted_image)

    print("Inverted color image saved to:", outputs_path)
    print("Execution time:", time.time() - start)

def main():
    image_path = "../../images/bread.jpg"
    outputs_path = "./outputs/test_write_image_output.jpg"

    # Read the image
    img = read_image(image_path)

    # Convert to RGBA if not already in that mode
    img = img.convert("RGBA")

    # Invert the colors and save the output
    invert_colors(img, outputs_path)

if __name__ == "__main__":
    main()
