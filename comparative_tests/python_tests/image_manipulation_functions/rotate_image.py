from . import utils
import time


def rotate_image(img, angle, outputs_path):
    start_time = time.time()

    # Ensure the output directory exists
    utils.ensure_output_dir(outputs_path)

    # Do the rotation
    if angle in [90, -90, 180, -180]:
        rotated_img = img.rotate(angle, expand=True)
    else:
        print("Unsupported rotation angle. Please use 90, -90, 180, or -180 degrees.")
        return

    # Save the rotated image
    output_file = utils.create_output_file(outputs_path)
    utils.save_image_as_jpeg(output_file, rotated_img)

    print("Rotated image saved to:", output_file)
    print("Execution time:", time.time() - start_time, "seconds")
