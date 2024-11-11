import os
from basic_functions.read_image import read_image
from image_manipulation_functions.utils import ensure_output_dir
from image_manipulation_functions.blur_filter import blur_filter
from image_manipulation_functions.grayscale_filter import grayscale_filter
from image_manipulation_functions.resize_image import resize_image
from image_manipulation_functions.rotate_image import rotate_image
from image_manipulation_functions.contrast_filter import contrast_filter


def main():
    IMAGE_PATH = "../../images/bread.jpg"
    OUTPUTS_PATH = "../python_outputs/"

    # Assurer que le répertoire de sortie existe
    ensure_output_dir(OUTPUTS_PATH)

    # Test BlurFilter
    print("\nTest BlurFilter:")
    img = read_image(IMAGE_PATH)
    blur_filter(img, intensity=3, outputs_path=os.path.join(OUTPUTS_PATH, "blurred_bread.jpg"))

    # Test GrayscaleFilter
    print("\nTest GrayscaleFilter:")
    img = read_image(IMAGE_PATH)
    grayscale_filter(img, intensity=100, outputs_path=os.path.join(OUTPUTS_PATH, "grayscale_bread.jpg"))

    # Test ResizeImage
    print("\nTest ResizeImage:")
    img = read_image(IMAGE_PATH)
    resize_image(img, new_width=800, new_height=400, outputs_path=os.path.join(OUTPUTS_PATH, "resized_bread.jpg"))

    # Test RotateImage
    print("\nTest RotateImage:")
    img = read_image(IMAGE_PATH)
    rotate_image(img, angle=90, outputs_path=os.path.join(OUTPUTS_PATH, "rotated_bread.jpg"))

    # Test ContrastFilter
    print("\nTest ContrastFilter:")
    img = read_image(IMAGE_PATH)
    contrast_filter(img, contrast_factor=100, outputs_path=os.path.join(OUTPUTS_PATH, "contrast_bread.jpg"))

if __name__ == "__main__":
    main()