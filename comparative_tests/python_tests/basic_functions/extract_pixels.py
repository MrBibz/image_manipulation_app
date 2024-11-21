from PIL import Image

def extract_pixels(image_path):
    # Get the dimensions of the image
    img = Image.open(image_path)
    width, height = img.size

    # Create a matrix to store the pixels
    pixels = []
    for y in range(height):
        row = []
        for x in range(width):
            # Get the color of the pixel at (x, y)
            rgba = img.getpixel((x, y))
            # Ensure it is in RGBA format (if the image has an alpha channel)
            if len(rgba) == 3:
                rgba = (*rgba, 255)  # Add full opacity if alpha is missing
            row.append(rgba)
        pixels.append(row)

    # Print the first five pixels of the first five rows
    for y in range(min(5, height)):
        for x in range(min(5, width)):
            r, g, b, a = pixels[y][x]
            print(f"Pixel [{y}, {x}]: R={r} G={g} B={b} A={a}")