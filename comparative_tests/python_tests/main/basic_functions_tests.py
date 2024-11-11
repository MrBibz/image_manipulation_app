from basic_functions.extract_pixels import extract_pixels
from basic_functions.get_image_dimensions import get_image_dimensions
from basic_functions.read_image import read_image


def main():
    IMAGE_PATH = "../../images/bread.jpg"

    # Test read_image
    print("\nTest read_image :")
    img = read_image(IMAGE_PATH)
    print("Image :", img)

    # Test get_image_dimensions
    if img is not None:
        print("\nTest get_image_dimensions :")
        width, height = get_image_dimensions(IMAGE_PATH)
        print("Dimensions de l'image :", width, "x", height)

    # Test extract_pixels
    print("\nTest extract_pixels :")
    extract_pixels(IMAGE_PATH)

# Run the main function
if __name__ == "__main__":
    main()
