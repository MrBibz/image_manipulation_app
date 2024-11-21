from basic_functions.read_image import read_image
from basic_functions.get_image_dimensions import get_image_dimensions
from basic_functions.extract_pixels import extract_pixels

def main():
    IMAGE_PATH = "../../images/bread.jpg"

    # Test ReadImage
    print("\nTest ReadImage:")
    img = read_image(IMAGE_PATH)
    print("Image:", img)

    # Test GetImageDimensions
    print("\nTest GetImageDimensions:")
    if img:
        width, height = get_image_dimensions(img)
        print(f"Image dimensions: {width} x {height}")

    # Test ExtractPixels
    print("\nTest ExtractPixels:")
    if img:
        extract_pixels(IMAGE_PATH)

if __name__ == "__main__":
    main()
